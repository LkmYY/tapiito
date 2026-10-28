// Package dbgate 提供 MySQL 并发闸门：HTTP 并发不等于数据库并发。
//
// 读、写使用各自独立的信号量，写侧预留容量给管理面；拿不到槽位时进入有界队列，
// 队列满或等待超时返回 ErrBusy（由上层翻译成 503 + Retry-After）。
package dbgate

import (
	"context"
	"errors"
	"time"
)

// ErrBusy 表示闸门已满/排队超时，调用方应返回 503。
var ErrBusy = errors.New("db gate busy")

type Gate struct {
	sem     chan struct{}
	queue   chan struct{}
	timeout time.Duration
}

// New 创建闸门：concurrency 为同时可进入数据库的请求数，queueLimit 为允许排队数。
func New(concurrency, queueLimit int, queueTimeout time.Duration) *Gate {
	if concurrency <= 0 {
		concurrency = 1
	}
	if queueLimit < 0 {
		queueLimit = 0
	}
	return &Gate{
		sem:     make(chan struct{}, concurrency),
		queue:   make(chan struct{}, queueLimit),
		timeout: queueTimeout,
	}
}

// Acquire 获取一个数据库槽位，返回的 release 必须调用（通常 defer）。
func (g *Gate) Acquire(ctx context.Context) (func(), error) {
	// 快路径：有空闲槽位直接进入
	select {
	case g.sem <- struct{}{}:
		return g.release, nil
	default:
	}

	// 慢路径：排队，队列满则立即失败
	select {
	case g.queue <- struct{}{}:
		defer func() { <-g.queue }()
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		return nil, ErrBusy
	}

	timer := time.NewTimer(g.timeout)
	defer timer.Stop()

	select {
	case g.sem <- struct{}{}:
		return g.release, nil
	case <-timer.C:
		return nil, ErrBusy
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (g *Gate) release() { <-g.sem }

// InUse 返回当前占用槽位数（用于指标）。
func (g *Gate) InUse() int { return len(g.sem) }

// QueueLen 返回当前排队数（用于指标）。
func (g *Gate) QueueLen() int { return len(g.queue) }
