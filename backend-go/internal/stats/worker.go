package stats

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	consumerGroup = "stats-workers"
	batchSize     = 500
	blockMs       = 2000
)

// Worker 消费 stats:visits，批量写明细表并累加日聚合表。
type Worker struct {
	rdb           *redis.Client
	db            *sql.DB
	log           *slog.Logger
	name          string
	retentionDays int
}

func NewWorker(rdb *redis.Client, db *sql.DB, name string, retentionDays int, log *slog.Logger) *Worker {
	if retentionDays <= 0 {
		retentionDays = 30
	}
	return &Worker{rdb: rdb, db: db, name: name, retentionDays: retentionDays, log: log}
}

// Run 阻塞运行，直到 ctx 取消。
func (w *Worker) Run(ctx context.Context) error {
	if err := w.ensureGroup(ctx); err != nil {
		return err
	}
	// 明细保留期清理：不清理的话 400 QPS 一天就是 3450 万行，磁盘会爆
	go w.cleanupLoop(ctx)
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if err := w.consumeOnce(ctx); err != nil && !errors.Is(err, redis.Nil) && ctx.Err() == nil {
			w.log.Error("stats consume failed", "err", err)
			time.Sleep(time.Second)
		}
	}
}

func (w *Worker) ensureGroup(ctx context.Context) error {
	err := w.rdb.XGroupCreateMkStream(ctx, StreamKey, consumerGroup, "0").Err()
	if err != nil && !isBusyGroup(err) {
		return err
	}
	return nil
}

func isBusyGroup(err error) bool {
	return err != nil && strings.Contains(err.Error(), "BUSYGROUP")
}

func (w *Worker) consumeOnce(ctx context.Context) error {
	res, err := w.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    consumerGroup,
		Consumer: w.name,
		Streams:  []string{StreamKey, ">"},
		Count:    batchSize,
		Block:    blockMs * time.Millisecond,
	}).Result()
	if err != nil {
		return err
	}
	if len(res) == 0 || len(res[0].Messages) == 0 {
		return nil
	}
	msgs := res[0].Messages

	if err := w.persist(ctx, msgs); err != nil {
		// 写库失败：不 ACK，交给重试（配合 pending 清理避免无限堆积）
		return err
	}
	ids := make([]string, 0, len(msgs))
	for _, m := range msgs {
		ids = append(ids, m.ID)
	}
	return w.rdb.XAck(ctx, StreamKey, consumerGroup, ids...).Err()
}

func (w *Worker) persist(ctx context.Context, msgs []redis.XMessage) error {
	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	insDetail, err := tx.PrepareContext(ctx,
		"INSERT INTO visit_stats (nfc_tag_id, ip_address, user_agent, referer, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer insDetail.Close()

	// 日聚合按 (tag, date) 累加
	upDaily, err := tx.PrepareContext(ctx, `
		INSERT INTO visit_stats_daily (nfc_tag_id, stat_date, count)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE count = count + VALUES(count)`)
	if err != nil {
		return err
	}
	defer upDaily.Close()

	type dailyKey struct {
		tagID int64
		date  string
	}
	daily := make(map[dailyKey]int64, 64)

	for _, m := range msgs {
		tagID := ParseInt(m.Values["tag_id"])
		if tagID <= 0 {
			continue
		}
		ts := time.Unix(ParseInt(m.Values["ts"]), 0)
		if ts.IsZero() || ts.Unix() <= 0 {
			ts = time.Now()
		}
		if _, err := insDetail.ExecContext(ctx, tagID,
			ParseStr(m.Values["ip"]), ParseStr(m.Values["ua"]), ParseStr(m.Values["referer"]),
			ts.Format("2006-01-02 15:04:05")); err != nil {
			return err
		}
		k := dailyKey{tagID: tagID, date: ts.Format("2006-01-02")}
		daily[k]++
	}

	for k, n := range daily {
		if _, err := upDaily.ExecContext(ctx, k.tagID, k.date, n); err != nil {
			return err
		}
	}

	return tx.Commit()
}

// cleanupLoop 每天清理一次超期明细。启动后延迟执行，避免与应用启动争抢 IO。
func (w *Worker) cleanupLoop(ctx context.Context) {
	timer := time.NewTimer(2 * time.Minute)
	select {
	case <-ctx.Done():
		timer.Stop()
		return
	case <-timer.C:
	}
	w.cleanupOnce(ctx)

	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.cleanupOnce(ctx)
		}
	}
}

// cleanupOnce 分批删除超期明细：一次删太多会长时间持锁、也会打满 IO。
func (w *Worker) cleanupOnce(ctx context.Context) {
	const batch = 5000
	cutoff := time.Now().AddDate(0, 0, -w.retentionDays)
	total := 0

	for ctx.Err() == nil {
		res, err := w.db.ExecContext(ctx,
			"DELETE FROM visit_stats WHERE created_at < ? LIMIT ?", cutoff, batch)
		if err != nil {
			w.log.Error("stats cleanup failed", "err", err, "cutoff", cutoff)
			return
		}
		n, _ := res.RowsAffected()
		total += int(n)
		if n < batch {
			break
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(100 * time.Millisecond):
		}
	}
	if total > 0 {
		w.log.Info("stats cleanup done", "deleted", total, "retention_days", w.retentionDays)
	}
}
