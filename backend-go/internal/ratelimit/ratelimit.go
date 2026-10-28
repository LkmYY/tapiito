// Package ratelimit 提供 Redis 原子限流（令牌桶）、临时黑名单与本地降级。
package ratelimit

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// tokenBucketLua 是令牌桶实现，返回 1 表示放行、0 表示超限。
// 注意：Redis 的 Lua 沙箱禁止访问未定义全局变量，nil 判断必须逐行赋值。
const tokenBucketLua = `
local rate = tonumber(ARGV[1])
local burst = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local ttl = tonumber(ARGV[4])

local data = redis.call('HMGET', KEYS[1], 'tokens', 'ts')
local tokens = tonumber(data[1])
local ts = tonumber(data[2])
if tokens == nil or ts == nil then
  tokens = burst
  ts = now
end

local delta = math.max(0, now - ts) / 1000.0
tokens = math.min(burst, tokens + delta * rate)
local allowed = 0
if tokens >= 1 then
  tokens = tokens - 1
  allowed = 1
end
redis.call('HSET', KEYS[1], 'tokens', tokens, 'ts', now)
redis.call('EXPIRE', KEYS[1], ttl)
return allowed
`

type Limiter struct {
	rdb    *redis.Client
	script *redis.Script
	// 本地降级：Redis 不可用时按实例整体限速，避免无限放行
	fallback *localBucket
}

func New(rdb *redis.Client, fallbackQPS int) *Limiter {
	if fallbackQPS <= 0 {
		fallbackQPS = 100
	}
	return &Limiter{
		rdb:      rdb,
		script:   redis.NewScript(tokenBucketLua),
		fallback: newLocalBucket(fallbackQPS, fallbackQPS*2),
	}
}

// Allow 判断 key 是否放行。rate/burst 为该 key 的配置。
func (l *Limiter) Allow(ctx context.Context, key string, rate, burst int) (bool, error) {
	if rate <= 0 {
		return true, nil
	}
	if burst < 1 {
		burst = 1
	}
	ttl := int((float64(burst)/float64(rate))*2) + 1
	raw, err := l.script.Run(ctx, l.rdb, []string{key},
		rate, burst, time.Now().UnixMilli(), ttl).Result()
	if err != nil {
		// Redis 限流不可用：降级到本地限速，宁可再生也不放行洪水
		if l.fallback.allow() {
			return true, fmt.Errorf("ratelimit degraded to local: %w", err)
		}
		return false, fmt.Errorf("ratelimit degraded and local full: %w", err)
	}
	return toInt(raw) == 1, nil
}

// toInt 兼容 RESP2 整数与 RESP3 布尔的返回编码。
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

// ---------- 黑名单 ----------

type Blacklist struct {
	rdb    *redis.Client
	window time.Duration
	hits   int
	ttl    time.Duration
}

func NewBlacklist(rdb *redis.Client, window time.Duration, hits int, ttl time.Duration) *Blacklist {
	if hits <= 0 {
		hits = 20
	}
	return &Blacklist{rdb: rdb, window: window, hits: hits, ttl: ttl}
}

func (b *Blacklist) key(ip string) string   { return "bl:ip:" + ip }
func (b *Blacklist) countKey(ip string) string { return "bl:count:" + ip }

// Blocked 判断 IP 是否在封禁中。
func (b *Blacklist) Blocked(ctx context.Context, ip string) (bool, error) {
	if ip == "" {
		return false, nil
	}
	n, err := b.rdb.Exists(ctx, b.key(ip)).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// RecordViolation 记录一次超限；窗口内达到阈值则封禁。返回是否刚刚被封禁。
func (b *Blacklist) RecordViolation(ctx context.Context, ip string) (bool, error) {
	if ip == "" {
		return false, nil
	}
	ck := b.countKey(ip)
	n, err := b.rdb.Incr(ctx, ck).Result()
	if err != nil {
		return false, err
	}
	if n == 1 {
		_ = b.rdb.Expire(ctx, ck, b.window).Err()
	}
	if int(n) < b.hits {
		return false, nil
	}
	if err := b.rdb.Set(ctx, b.key(ip), "1", b.ttl).Err(); err != nil {
		return false, err
	}
	_ = b.rdb.Del(ctx, ck).Err()
	return true, nil
}

// ---------- 本地降级桶 ----------

type localBucket struct {
	mu     sync.Mutex
	tokens float64
	burst  float64
	rate   float64
	last   time.Time
}

func newLocalBucket(rate, burst int) *localBucket {
	return &localBucket{
		tokens: float64(burst),
		burst:  float64(burst),
		rate:   float64(rate),
		last:   time.Now(),
	}
}

func (b *localBucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	now := time.Now()
	b.tokens += now.Sub(b.last).Seconds() * b.rate
	if b.tokens > b.burst {
		b.tokens = b.burst
	}
	b.last = now
	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}
