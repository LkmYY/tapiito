// Command api 启动 Tapiito 双入口服务：
//
//	:8080 公开 API（只读 + 少量公开写）
//	:8081 管理 API（JWT 保护，绝不暴露到公网入口）
package main

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/lkmyy/tapiito-go/internal/auth"
	"github.com/lkmyy/tapiito-go/internal/bloom"
	"github.com/lkmyy/tapiito-go/internal/cache"
	"github.com/lkmyy/tapiito-go/internal/config"
	"github.com/lkmyy/tapiito-go/internal/db"
	"github.com/lkmyy/tapiito-go/internal/dbgate"
	"github.com/lkmyy/tapiito-go/internal/feedback"
	httpapi "github.com/lkmyy/tapiito-go/internal/http"
	"github.com/lkmyy/tapiito-go/internal/order"
	"github.com/lkmyy/tapiito-go/internal/ratelimit"
	"github.com/lkmyy/tapiito-go/internal/sms"
	"github.com/lkmyy/tapiito-go/internal/stats"
	"github.com/lkmyy/tapiito-go/internal/upload"
	"github.com/lkmyy/tapiito-go/internal/tag"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("config invalid", "err", err)
		os.Exit(1)
	}

	log := newLogger(cfg)
	slog.SetDefault(log)

	if cfg.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	// ---------- 依赖 ----------
	mysqlDB, err := db.Open(cfg.DB)
	if err != nil {
		log.Error("mysql unavailable", "err", err)
		os.Exit(1)
	}
	defer mysqlDB.Close()

	cacheRDB := newRedis(cfg.Cache, "cache")
	defer cacheRDB.Close()
	controlRDB := newRedis(cfg.Control, "control")
	defer controlRDB.Close()

	// ---------- 组件 ----------
	readGate := dbgate.New(cfg.DB.ReadConcurrency, cfg.DB.QueueLimit, cfg.DB.QueueTimeout)
	writeGate := dbgate.New(cfg.DB.WriteConcurrency, cfg.DB.QueueLimit, cfg.DB.QueueTimeout)
	pageCache := cache.NewPageCache(cacheRDB, cfg.CDN.PageTTL, cfg.CDN.PageJitter, cfg.CDN.MissTTL)
	filter := bloom.New(controlRDB, bloom.Options{
		KeyV1:       cfg.Bloom.KeyV1,
		KeyV2:       cfg.Bloom.KeyV2,
		CurrentKey:  cfg.Bloom.CurrentKey,
		BuildingKey: cfg.Bloom.BuildingKey,
		DeltaKey:    cfg.Bloom.DeltaKey,
		ErrorRate:   cfg.Bloom.ErrorRate,
		Capacity:    cfg.Bloom.Capacity,
		DeltaTTL:    int(cfg.Bloom.DeltaTTL.Seconds()),
	})
	limiter := ratelimit.New(controlRDB, cfg.Limit.GlobalQPS)
	blacklist := ratelimit.NewBlacklist(controlRDB,
		time.Duration(cfg.Limit.BlacklistWindow)*time.Second,
		cfg.Limit.BlacklistHits, cfg.Limit.BlacklistTTL)
	issuer := auth.NewIssuer(cfg.JWT.Secret, cfg.JWT.Expire, cfg.JWT.Issuer)

	repo := tag.NewRepo(mysqlDB)
	svc := tag.NewService(tag.Deps{
		Repo: repo, Cache: pageCache, Bloom: filter,
		ReadGate: readGate, WriteGate: writeGate,
		QueryTimeout: cfg.DB.QueryTimeout, Logger: log,
	})

	// ---------- 启动前置检查 ----------
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := ensureAdminPassword(ctx, mysqlDB, cfg, log); err != nil {
		log.Error("ensure admin failed", "err", err)
		os.Exit(1)
	}
	if err := ensureBloom(ctx, filter, svc, log); err != nil {
		log.Error("bloom bootstrap failed", "err", err)
		os.Exit(1)
	}

	// ---------- 路由 ----------
	// 健康检查必须用独立 context：启动期的 ctx 会在超时后失效，导致探活永久失败
	healthChecks := []func() error{
		func() error { return ping(context.Background(), mysqlDB) },
		func() error { return cacheRDB.Ping(context.Background()).Err() },
		func() error { return controlRDB.Ping(context.Background()).Err() },
	}
	uploadStore, err := upload.New(*cfg, log)
	if err != nil {
		log.Error("upload store init failed", "err", err)
		os.Exit(1)
	}
	smsSender := sms.New(cfg.Twilio, log)

	orderSvc := order.NewService(order.Deps{
		DB: mysqlDB, ReadGate: readGate, WriteGate: writeGate,
		QueryTimeout: cfg.DB.QueryTimeout, Bloom: filter, Cache: pageCache, Logger: log,
	})
	feedbackSvc := feedback.NewService(feedback.Deps{
		DB: mysqlDB, ReadGate: readGate, WriteGate: writeGate,
		QueryTimeout: cfg.DB.QueryTimeout, Logger: log,
	})
	statsSvc := stats.NewService(stats.Deps{
		DB: mysqlDB, ReadGate: readGate, QueryTimeout: 5 * time.Second, Logger: log,
	})

	deps := httpapi.Deps{
		Cfg: cfg, Log: log, Limiter: limiter, Blacklist: blacklist, Issuer: issuer,
		Public: httpapi.NewPublicHandler(svc, uploadStore, smsSender.Send, cfg.PublicBaseURL, controlRDB, log),
		Admin:  httpapi.NewAdminHandler(mysqlDB, svc, issuer, log),
		Order:  httpapi.NewOrderHandler(orderSvc, log),
		Feedback: httpapi.NewFeedbackHandler(feedbackSvc, log),
		Stats:  httpapi.NewStatsHandler(statsSvc, log),
		Health: healthChecks,
	}
	publicSrv := &http.Server{
		Addr:              cfg.PublicAddr,
		Handler:           httpapi.NewPublicRouter(deps),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	adminSrv := &http.Server{
		Addr:              cfg.AdminAddr,
		Handler:           httpapi.NewAdminRouter(deps),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errc := make(chan error, 2)
	go func() {
		log.Info("public api listening", "addr", cfg.PublicAddr)
		if err := publicSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errc <- err
		}
	}()
	go func() {
		log.Info("admin api listening", "addr", cfg.AdminAddr)
		if err := adminSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errc <- err
		}
	}()

	// ---------- 优雅关闭 ----------
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errc:
		log.Error("server error", "err", err)
	case sig := <-stop:
		log.Info("shutting down", "signal", sig.String())
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()
	_ = publicSrv.Shutdown(shutdownCtx)
	_ = adminSrv.Shutdown(shutdownCtx)
	log.Info("bye")
}

func newLogger(cfg *config.Config) *slog.Logger {
	level := slog.LevelInfo
	if !cfg.IsProd() {
		level = slog.LevelDebug
	}
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}

func newRedis(cfg config.RedisConfig, name string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
}

func ping(ctx context.Context, db *sql.DB) error {
	c, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return db.PingContext(c)
}

// ensureAdminPassword 在管理员不存在时按环境变量创建；生产环境未提供密码则拒绝启动。
func ensureAdminPassword(ctx context.Context, conn *sql.DB, cfg *config.Config, log *slog.Logger) error {
	if cfg.DefaultAdmin == "" {
		return nil
	}
	var id int64
	err := conn.QueryRowContext(ctx,
		"SELECT id FROM admin_user WHERE username = ? LIMIT 1", cfg.DefaultAdmin).Scan(&id)
	if err == nil {
		return nil // 已存在，绝不覆盖密码
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	if cfg.DefaultAdminPw == "" {
		return errors.New("DEFAULT_ADMIN_PASSWORD 未设置，无法创建初始管理员")
	}
	hash, err := auth.HashPassword(cfg.DefaultAdminPw)
	if err != nil {
		return err
	}
	if _, err := conn.ExecContext(ctx,
		"INSERT INTO admin_user (username, password_hash, role, created_at) VALUES (?, ?, 'admin', NOW())",
		cfg.DefaultAdmin, hash); err != nil {
		return err
	}
	log.Info("default admin created", "username", cfg.DefaultAdmin)
	return nil
}

// ensureBloom 启动时若过滤器缺失则同步构建，避免上线即全站 503。
func ensureBloom(ctx context.Context, f *bloom.Filter, svc *tag.Service, log *slog.Logger) error {
	current, err := f.Current(ctx)
	if err != nil {
		return err
	}
	if current != "" {
		if _, ready, err := f.Exists(ctx, "__warmup_probe__"); err == nil && ready {
			log.Info("bloom ready", "filter", current)
			return nil
		}
		log.Warn("bloom current pointer exists but filter unusable, rebuilding")
	}

	ids, err := svc.AllNfcIDs(ctx)
	if err != nil {
		return err
	}
	res, err := f.Build(ctx, ids)
	if err != nil {
		return err
	}
	log.Info("bloom built", "target", res.Target, "loaded", res.Loaded, "delta", res.DeltaMerged)
	return nil
}
