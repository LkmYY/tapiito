// Command bloom-builder 执行一次 Bloom 轮换构建后退出，供 cron / K8s CronJob 调用。
//
// 建议每天凌晨执行一次：构建目标过滤器 → 合并 delta → 原子切换 bloom:current。
// 构建失败时保持旧过滤器不变，只告警，不影响线上读。
package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/lkmyy/tapiito-go/internal/bloom"
	"github.com/lkmyy/tapiito-go/internal/config"
	"github.com/lkmyy/tapiito-go/internal/db"
	"github.com/lkmyy/tapiito-go/internal/tag"
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

	filter := bloom.New(rdb, bloom.Options{
		KeyV1:       cfg.Bloom.KeyV1,
		KeyV2:       cfg.Bloom.KeyV2,
		CurrentKey:  cfg.Bloom.CurrentKey,
		BuildingKey: cfg.Bloom.BuildingKey,
		DeltaKey:    cfg.Bloom.DeltaKey,
		ErrorRate:   cfg.Bloom.ErrorRate,
		Capacity:    cfg.Bloom.Capacity,
		DeltaTTL:    int(cfg.Bloom.DeltaTTL.Seconds()),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	start := time.Now()
	ids, err := tag.NewRepo(conn).AllNfcIDs(ctx)
	if err != nil {
		log.Error("load nfc ids failed", "err", err)
		os.Exit(1)
	}
	log.Info("bloom build start", "ids", len(ids))

	res, err := filter.Build(ctx, ids)
	if err != nil {
		// 构建失败：旧过滤器仍在服务，只需告警
		log.Error("bloom build failed, keeping old filter", "err", err)
		os.Exit(1)
	}
	log.Info("bloom build done",
		"target", res.Target, "loaded", res.Loaded,
		"delta_merged", res.DeltaMerged, "cost", time.Since(start).String())
}
