// Command worker 常驻运行访问统计消费者：Redis Stream → MySQL 批量落库 + 日聚合。
package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/lkmyy/tapiito-go/internal/config"
	"github.com/lkmyy/tapiito-go/internal/db"
	"github.com/lkmyy/tapiito-go/internal/stats"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("config invalid", "err", err)
		os.Exit(1)
	}
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	conn, err := db.Open(cfg.DB)
	if err != nil {
		log.Error("mysql unavailable", "err", err)
		os.Exit(1)
	}
	defer conn.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Control.Addr,
		Password: cfg.Control.Password,
		DB:       cfg.Control.DB,
		PoolSize: cfg.Control.PoolSize,
	})
	defer rdb.Close()

	host, _ := os.Hostname()
	name := host + "-" + timeNow()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	w := stats.NewWorker(rdb, conn, name, cfg.StatsRetentionDays, log)
	log.Info("stats worker started", "consumer", name, "retention_days", cfg.StatsRetentionDays)
	if err := w.Run(ctx); err != nil && ctx.Err() == nil {
		log.Error("worker exited", "err", err)
		os.Exit(1)
	}
	log.Info("stats worker stopped")
}

func timeNow() string {
	return time.Now().Format("150405")
}
