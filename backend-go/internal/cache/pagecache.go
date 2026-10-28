// Package cache 提供 Cache Aside 缓存与热点保护（singleflight）。
//
// 只缓存公开展示 DTO，绝不缓存手机号、user_id、审计字段等内部数据。
// key 带版本前缀，发布新版本时靠改前缀失效，不需要 FLUSHALL。
package cache

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

const (
	pageKeyPrefix = "nfc:page:v1:"
	missKeyPrefix = "nfc:miss:v1:"
)

// ErrNotFound 由 loader 返回，表示数据确实不存在（会被写成负缓存）。
var ErrNotFound = errors.New("not found")

type PageCache struct {
	rdb     *redis.Client
	ttl     time.Duration
	jitter  time.Duration
	missTTL time.Duration
	sf      singleflight.Group
}

func NewPageCache(rdb *redis.Client, ttl, jitter, missTTL time.Duration) *PageCache {
	return &PageCache{rdb: rdb, ttl: ttl, jitter: jitter, missTTL: missTTL}
}

func pageKey(nfcID string) string { return pageKeyPrefix + nfcID }
func missKey(nfcID string) string { return missKeyPrefix + nfcID }

// TTL 返回带随机抖动的过期时间，避免同一时刻大面积失效。
func (c *PageCache) TTL() time.Duration {
	if c.jitter <= 0 {
		return c.ttl
	}
	return c.ttl + time.Duration(rand.Int63n(int64(c.jitter)))
}

// Get 读取正缓存。
func (c *PageCache) Get(ctx context.Context, nfcID string) ([]byte, bool, error) {
	raw, err := c.rdb.Get(ctx, pageKey(nfcID)).Bytes()
	if err == nil {
		return raw, true, nil
	}
	if errors.Is(err, redis.Nil) {
		return nil, false, nil
	}
	// Redis 故障时降级为直接查库，不阻断请求
	return nil, false, err
}

// IsMiss 判断是否命中负缓存。
func (c *PageCache) IsMiss(ctx context.Context, nfcID string) bool {
	n, err := c.rdb.Exists(ctx, missKey(nfcID)).Result()
	return err == nil && n > 0
}

// Set 回填正缓存并清除负缓存。
func (c *PageCache) Set(ctx context.Context, nfcID string, payload []byte) error {
	pipe := c.rdb.Pipeline()
	pipe.Set(ctx, pageKey(nfcID), payload, c.TTL())
	pipe.Del(ctx, missKey(nfcID))
	_, err := pipe.Exec(ctx)
	return err
}

// SetMiss 写入短负缓存，降低不存在 ID 对 MySQL 的重复查询。
func (c *PageCache) SetMiss(ctx context.Context, nfcID string) error {
	return c.rdb.Set(ctx, missKey(nfcID), "1", c.missTTL).Err()
}

// Del 删除缓存（写路径使用）。失败时由 TTL 兜底，上层负责重试与告警。
func (c *PageCache) Del(ctx context.Context, nfcIDs ...string) error {
	if len(nfcIDs) == 0 {
		return nil
	}
	keys := make([]string, 0, len(nfcIDs)*2)
	for _, id := range nfcIDs {
		keys = append(keys, pageKey(id), missKey(id))
	}
	return c.rdb.Del(ctx, keys...).Err()
}

// GetOrLoad 是 Cache Aside 的读路径：缓存优先，未命中时用 singleflight 收敛回源。
//
// loader 返回 ErrNotFound 时写负缓存并透传错误；返回其他错误不写缓存。
func (c *PageCache) GetOrLoad(
	ctx context.Context,
	nfcID string,
	loader func(context.Context) ([]byte, error),
) (data []byte, fromCache bool, err error) {
	if raw, ok, cacheErr := c.Get(ctx, nfcID); ok {
		return raw, true, nil
	} else if cacheErr != nil {
		// 缓存不可用：跳过负缓存判断，直接回源
		return c.load(ctx, nfcID, loader)
	}

	if c.IsMiss(ctx, nfcID) {
		return nil, false, ErrNotFound
	}

	v, sfErr, _ := c.sf.Do(pageKey(nfcID), func() (interface{}, error) {
		// 双重检查：并发等待者里可能已有人回填
		if raw, ok, _ := c.Get(ctx, nfcID); ok {
			return raw, nil
		}
		raw, loadErr := loader(ctx)
		if loadErr != nil {
			if errors.Is(loadErr, ErrNotFound) {
				_ = c.SetMiss(ctx, nfcID)
			}
			return nil, loadErr
		}
		_ = c.Set(ctx, nfcID, raw)
		return raw, nil
	})
	if sfErr != nil {
		return nil, false, sfErr
	}
	out, ok := v.([]byte)
	if !ok {
		return nil, false, fmt.Errorf("cache: unexpected singleflight payload type %T", v)
	}
	return out, false, nil
}

func (c *PageCache) load(ctx context.Context, nfcID string, loader func(context.Context) ([]byte, error)) ([]byte, bool, error) {
	raw, err := loader(ctx)
	if err != nil {
		return nil, false, err
	}
	return raw, false, nil
}
