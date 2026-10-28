// Package bloom 实现名片存在性预判与双 Bloom 轮换。
//
// 轮换期间新增 ID 同时写入 bloom:delta（Set），避免构建中的目标过滤器漏掉新数据
// 导致轮换后有效名片被误判为不存在。
package bloom

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	bfReserve = "BF.RESERVE"
	bfAdd     = "BF.ADD"
	bfMAdd    = "BF.MADD"
	bfExists  = "BF.EXISTS"
)

var (
	// ErrNotReady 表示当前过滤器缺失（未构建/丢失），调用方应返回 503 而不是 404。
	ErrNotReady = errors.New("bloom filter not ready")
)

type Filter struct {
	rdb        *redis.Client
	keyV1      string
	keyV2      string
	currentKey string
	buildingKey string
	deltaKey   string
	errorRate  float64
	capacity   int
	deltaTTL   int // 秒
}

type Options struct {
	KeyV1       string
	KeyV2       string
	CurrentKey  string
	BuildingKey string
	DeltaKey    string
	ErrorRate   float64
	Capacity    int
	DeltaTTL    int
}

func New(rdb *redis.Client, o Options) *Filter {
	return &Filter{
		rdb:         rdb,
		keyV1:       o.KeyV1,
		keyV2:       o.KeyV2,
		currentKey:  o.CurrentKey,
		buildingKey: o.BuildingKey,
		deltaKey:    o.DeltaKey,
		errorRate:   o.ErrorRate,
		capacity:    o.Capacity,
		deltaTTL:    o.DeltaTTL,
	}
}

// Exists 判断 nfc_id 是否可能存在。
//
// 返回 (exists, ready, err)：
//   - ready=false 表示过滤器当前不可用，调用方必须返回 503，绝不能当成不存在。
//   - err 仅在 Redis 通信异常时返回。
func (f *Filter) Exists(ctx context.Context, nfcID string) (bool, bool, error) {
	name, err := f.current(ctx)
	if err != nil {
		return false, false, err
	}
	if name == "" {
		return false, false, nil
	}

	raw, err := f.rdb.Do(ctx, bfExists, name, nfcID).Result()
	switch {
	case errors.Is(err, redis.Nil):
		// key 不存在（过滤器被删）：不能判 404
		return false, false, nil
	case err != nil:
		return false, false, err
	}
	if toInt(raw) == 1 {
		return true, true, nil
	}

	// 主过滤器未命中，检查轮换期间的增量集合
	inDelta, err := f.rdb.SIsMember(ctx, f.deltaKey, nfcID).Result()
	if err != nil {
		return false, true, err
	}
	return inDelta, true, nil
}

// Add 记录新增名片：同时写当前过滤器、目标过滤器（若在构建中）与增量集合。
func (f *Filter) Add(ctx context.Context, nfcID string) error {
	name, err := f.current(ctx)
	if err != nil {
		return err
	}

	pipe := f.rdb.Pipeline()
	pipe.SAdd(ctx, f.deltaKey, nfcID)
	if f.deltaTTL > 0 {
		pipe.Expire(ctx, f.deltaKey, seconds(f.deltaTTL))
	}
	if name != "" {
		pipe.Do(ctx, bfAdd, name, nfcID)
	}
	if building, _ := f.rdb.Get(ctx, f.buildingKey).Result(); building != "" {
		pipe.Do(ctx, bfAdd, building, nfcID)
	}
	_, err = pipe.Exec(ctx)
	return err
}

// Current 返回当前过滤器名（空表示未构建）。
func (f *Filter) Current(ctx context.Context) (string, error) { return f.current(ctx) }

func (f *Filter) current(ctx context.Context) (string, error) {
	name, err := f.rdb.Get(ctx, f.currentKey).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return name, err
}

// BuildResult 描述一次轮换构建的结果。
type BuildResult struct {
	Target      string
	Loaded      int
	DeltaMerged int
}

// Build 构建目标过滤器并原子切换。
//
// ids 应来自 MySQL 全表的 nfc_id（写路径已在 delta 中合并，删除的 ID 自然从新过滤器消失）。
// 构建失败时不动当前过滤器，返回错误供告警。
func (f *Filter) Build(ctx context.Context, ids []string) (*BuildResult, error) {
	current, err := f.current(ctx)
	if err != nil {
		return nil, err
	}

	// 选择与当前相反的那个 key 作为目标
	target := f.keyV1
	if current == f.keyV1 {
		target = f.keyV2
	}

	// 只删除目标过滤器，绝不碰当前过滤器
	if err := f.rdb.Del(ctx, target).Err(); err != nil {
		return nil, fmt.Errorf("del target: %w", err)
	}
	if err := f.rdb.Do(ctx, bfReserve, target, strconv.FormatFloat(f.errorRate, 'f', -1, 64), f.capacity).Err(); err != nil {
		return nil, fmt.Errorf("reserve %s: %w", target, err)
	}
	// 标记构建中，写路径会同步写入该过滤器
	if err := f.rdb.Set(ctx, f.buildingKey, target, 0).Err(); err != nil {
		return nil, err
	}
	defer f.rdb.Del(context.WithoutCancel(ctx), f.buildingKey)

	const batch = 1000
	for i := 0; i < len(ids); i += batch {
		end := i + batch
		if end > len(ids) {
			end = len(ids)
		}
		args := make([]interface{}, 0, end-i+2)
		args = append(args, bfMAdd, target)
		for _, id := range ids[i:end] {
			args = append(args, id)
		}
		if err := f.rdb.Do(ctx, args...).Err(); err != nil {
			return nil, fmt.Errorf("madd batch %d: %w", i/batch, err)
		}
	}

	// 合并轮换期间的新增 ID
	deltaIDs, err := f.rdb.SMembers(ctx, f.deltaKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	for i := 0; i < len(deltaIDs); i += batch {
		end := i + batch
		if end > len(deltaIDs) {
			end = len(deltaIDs)
		}
		args := make([]interface{}, 0, end-i+2)
		args = append(args, bfMAdd, target)
		for _, id := range deltaIDs[i:end] {
			args = append(args, id)
		}
		if err := f.rdb.Do(ctx, args...).Err(); err != nil {
			return nil, fmt.Errorf("merge delta batch %d: %w", i/batch, err)
		}
	}

	// 原子切换
	if err := f.rdb.Set(ctx, f.currentKey, target, 0).Err(); err != nil {
		return nil, err
	}

	return &BuildResult{Target: target, Loaded: len(ids), DeltaMerged: len(deltaIDs)}, nil
}

func seconds(s int) time.Duration { return time.Duration(s) * time.Second }

// toInt 兼容 RedisBloom 的两种返回编码：RESP2 的整数与 RESP3 的布尔。
func toInt(v any) int {
	switch x := v.(type) {
	case int64:
		return int(x)
	case int:
		return x
	case bool:
		if x {
			return 1
		}
		return 0
	case string:
		if x == "1" || x == "true" {
			return 1
		}
		return 0
	}
	return 0
}
