package dbgate

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 并发度必须被严格遵守，否则 MySQL 会被打爆。
func TestGateRespectsConcurrency(t *testing.T) {
	g := New(4, 100, time.Second)

	var inFlight, peak int64
	var wg sync.WaitGroup
	for i := 0; i < 40; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			release, err := g.Acquire(context.Background())
			if err != nil {
				t.Errorf("acquire: %v", err)
				return
			}
			cur := atomic.AddInt64(&inFlight, 1)
			for {
				old := atomic.LoadInt64(&peak)
				if cur <= old || atomic.CompareAndSwapInt64(&peak, old, cur) {
					break
				}
			}
			time.Sleep(20 * time.Millisecond)
			atomic.AddInt64(&inFlight, -1)
			release()
		}()
	}
	wg.Wait()

	if peak > 4 {
		t.Fatalf("peak concurrency %d exceeded limit 4", peak)
	}
	if peak < 2 {
		t.Fatalf("peak concurrency %d too low, gate may be over-serializing", peak)
	}
}

// 队列满时必须快速失败，不能无限堆积 goroutine。
func TestGateQueueFullReturnsBusy(t *testing.T) {
	g := New(1, 1, 500*time.Millisecond)

	release, err := g.Acquire(context.Background())
	if err != nil {
		t.Fatalf("first acquire: %v", err)
	}

	// 占满队列中的一个位置
	queued := make(chan error, 1)
	go func() {
		r2, err := g.Acquire(context.Background())
		if err == nil {
			r2()
		}
		queued <- err
	}()
	time.Sleep(30 * time.Millisecond)

	// 第三个请求：槽位被占、队列已满 → 立即 ErrBusy
	if _, err := g.Acquire(context.Background()); !errors.Is(err, ErrBusy) {
		t.Fatalf("expected ErrBusy, got %v", err)
	}

	release()
	if err := <-queued; err != nil {
		t.Fatalf("queued acquire should succeed after release: %v", err)
	}
}

// 排队超时也必须返回 ErrBusy。
func TestGateQueueTimeout(t *testing.T) {
	g := New(1, 10, 50*time.Millisecond)

	release, err := g.Acquire(context.Background())
	if err != nil {
		t.Fatalf("acquire: %v", err)
	}
	defer release()

	start := time.Now()
	if _, err := g.Acquire(context.Background()); !errors.Is(err, ErrBusy) {
		t.Fatalf("expected ErrBusy, got %v", err)
	}
	if elapsed := time.Since(start); elapsed > 500*time.Millisecond {
		t.Fatalf("timeout path took too long: %v", elapsed)
	}
}

// 请求取消应立即返回，不应把槽位泄漏给已取消的请求。
func TestGateContextCancel(t *testing.T) {
	g := New(1, 10, 5*time.Second)
	release, err := g.Acquire(context.Background())
	if err != nil {
		t.Fatalf("acquire: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if _, err := g.Acquire(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	release()
}
