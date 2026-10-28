// Package stats 负责访问统计的异步写入：请求只写 Redis Stream，由 worker 批量落库。
package stats

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	StreamKey = "stats:visits"
	// MaxLen 控制 Stream 长度，避免无限增长
	MaxLen = 100000
)

// Recorder 把一次访问写进 Stream；失败只丢日志，绝不阻塞公开请求。
type Recorder struct {
	rdb *redis.Client
}

func NewRecorder(rdb *redis.Client) *Recorder { return &Recorder{rdb: rdb} }

// Record 实现 tag.StatsRecorder。
func (r *Recorder) Record(ctx context.Context, tagID int64, ip, ua, referer string) {
	// 用独立 context：请求可能已取消，但统计仍应尽量记录
	c, cancel := context.WithTimeout(context.WithoutCancel(ctx), 200*time.Millisecond)
	defer cancel()

	_ = r.rdb.XAdd(c, &redis.XAddArgs{
		Stream: StreamKey,
		MaxLen: MaxLen,
		Approx: true,
		Values: map[string]any{
			"tag_id":  tagID,
			"ip":      ip,
			"ua":      truncate(ua, 500),
			"referer": truncate(referer, 500),
			"ts":      time.Now().Unix(),
		},
	}).Err()
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n]
}

// ParseInt 解析 Stream 里的字符串字段。
func ParseInt(v any) int64 {
	switch x := v.(type) {
	case string:
		n, _ := strconv.ParseInt(x, 10, 64)
		return n
	case int64:
		return x
	case int:
		return int64(x)
	}
	return 0
}

func ParseStr(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
