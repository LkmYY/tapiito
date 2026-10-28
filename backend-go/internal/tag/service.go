package tag

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/lkmyy/tapiito-go/internal/bloom"
	"github.com/lkmyy/tapiito-go/internal/cache"
	"github.com/lkmyy/tapiito-go/internal/dbgate"
)

var (
	// ErrUnavailable 表示依赖不可用（Bloom 未就绪 / Redis 故障 / 闸门满），应返回 503。
	ErrUnavailable = errors.New("service temporarily unavailable")
)

// StatsRecorder 记录访问统计（异步，不阻塞主链路）。
type StatsRecorder interface {
	Record(ctx context.Context, tagID int64, ip, ua, referer string)
}

type Service struct {
	repo         *Repo
	cache        *cache.PageCache
	bloom        *bloom.Filter
	readGate     *dbgate.Gate
	writeGate    *dbgate.Gate
	queryTimeout time.Duration
	stats        StatsRecorder
	log          *slog.Logger
}

type Deps struct {
	Repo         *Repo
	Cache        *cache.PageCache
	Bloom        *bloom.Filter
	ReadGate     *dbgate.Gate
	WriteGate    *dbgate.Gate
	QueryTimeout time.Duration
	Stats        StatsRecorder
	Logger       *slog.Logger
}

func NewService(d Deps) *Service {
	if d.QueryTimeout <= 0 {
		d.QueryTimeout = 3 * time.Second
	}
	return &Service{
		repo:         d.Repo,
		cache:        d.Cache,
		bloom:        d.Bloom,
		readGate:     d.ReadGate,
		writeGate:    d.WriteGate,
		queryTimeout: d.QueryTimeout,
		stats:        d.Stats,
		log:          d.Logger,
	}
}

// PublicPage 是公开读链路：Bloom → 缓存 → singleflight → DB 闸门 → MySQL。
//
// 停用标签同样返回 200（status=disabled），由前端展示品牌页，不能当 404。
func (s *Service) PublicPage(ctx context.Context, nfcID, ip, ua, referer string) (*PublicPage, error) {
	exists, ready, err := s.bloom.Exists(ctx, nfcID)
	if err != nil {
		s.log.Warn("bloom check failed", "nfc_id", nfcID, "err", err)
		return nil, ErrUnavailable
	}
	if !ready {
		// 过滤器缺失：绝不能把所有请求判为不存在
		s.log.Error("bloom filter not ready", "nfc_id", nfcID)
		return nil, ErrUnavailable
	}
	if !exists {
		return nil, ErrNotFound
	}

	raw, fromCache, err := s.cache.GetOrLoad(ctx, nfcID, func(ctx context.Context) ([]byte, error) {
		return s.loadPage(ctx, nfcID, ip, ua, referer)
	})
	if err != nil {
		switch {
		case errors.Is(err, cache.ErrNotFound), errors.Is(err, ErrNotFound):
			return nil, ErrNotFound
		case errors.Is(err, dbgate.ErrBusy):
			return nil, ErrUnavailable
		default:
			return nil, err
		}
	}
	if !fromCache {
		s.log.Debug("cache miss, served from db", "nfc_id", nfcID)
	}

	var page PublicPage
	if err := json.Unmarshal(raw, &page); err != nil {
		return nil, fmt.Errorf("decode cached page: %w", err)
	}
	return &page, nil
}

// loadPage 在闸门内查库并把 DTO 序列化返回（由 singleflight 收敛并发）。
func (s *Service) loadPage(ctx context.Context, nfcID, ip, ua, referer string) ([]byte, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	tag, err := s.repo.GetByNfcID(qctx, nfcID)
	if errors.Is(err, ErrNotFound) {
		return nil, cache.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	page := tag.ToPublic()
	raw, err := json.Marshal(page)
	if err != nil {
		return nil, err
	}

	if s.stats != nil {
		s.stats.Record(ctx, tag.ID, ip, ua, referer)
	}
	return raw, nil
}

// Gallery 返回展厅卡片：直接查库，受读闸门保护。
func (s *Service) Gallery(ctx context.Context, limit int) ([]GalleryItem, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()
	return s.repo.Gallery(qctx, limit)
}

// PrewarmBloom 供启动时/定时任务构建 Bloom。
func (s *Service) AllNfcIDs(ctx context.Context) ([]string, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()
	return s.repo.AllNfcIDs(ctx)
}

// ---------- 管理写入（DB 先提交，再失效缓存；绝不先写 Redis） ----------

func (s *Service) AdminCreate(ctx context.Context, in CreateInput) (*APITag, error) {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	t, err := s.repo.Create(qctx, in)
	if err != nil {
		return nil, err
	}

	// 事务已提交：补 Bloom（失败仅告警，delta + 轮换兜底），清负缓存
	bg := context.WithoutCancel(ctx)
	if err := s.bloom.Add(bg, t.NfcID); err != nil {
		s.log.Error("bloom add failed", "nfc_id", t.NfcID, "err", err)
	}
	s.invalidate(bg, t.NfcID)

	out := t.ToAPI()
	return &out, nil
}

func (s *Service) AdminUpdate(ctx context.Context, id int64, in UpdateInput) (*APITag, error) {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	t, err := s.repo.Update(qctx, id, in)
	if err != nil {
		return nil, err
	}
	s.invalidate(context.WithoutCancel(ctx), t.NfcID)

	out := t.ToAPI()
	return &out, nil
}

func (s *Service) AdminDelete(ctx context.Context, id int64) error {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return ErrUnavailable
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	// 先取 nfc_id，删除后就查不到了
	t, err := s.repo.GetByID(qctx, id)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(qctx, id); err != nil {
		return err
	}
	// Bloom 不支持高效删除：保留，由下次轮换自然清理
	s.invalidate(context.WithoutCancel(ctx), t.NfcID)
	return nil
}

// BatchAction 执行 enable / disable / delete。
func (s *Service) BatchAction(ctx context.Context, ids []int64, action string) (int, error) {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return 0, ErrUnavailable
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var nfcIDs []string
	switch action {
	case "enable", "disable":
		status := "enabled"
		if action == "disable" {
			status = "disabled"
		}
		nfcIDs, err = s.repo.UpdateStatusBatch(qctx, ids, status)
	case "delete":
		nfcIDs, err = s.repo.DeleteBatch(qctx, ids)
	default:
		return 0, fmt.Errorf("unsupported action: %s", action)
	}
	if err != nil {
		return 0, err
	}

	if len(nfcIDs) > 0 {
		s.invalidate(context.WithoutCancel(ctx), nfcIDs...)
	}
	return len(nfcIDs), nil
}

func (s *Service) Creators(ctx context.Context) ([]CreatorItem, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()
	return s.repo.Creators(qctx)
}

func (s *Service) GetByID(ctx context.Context, id int64) (*APITag, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	t, err := s.repo.GetByID(qctx, id)
	if err != nil {
		return nil, err
	}
	out := t.ToAPI()
	return &out, nil
}

// List 返回管理端分页列表（total_pages 由调用方计算，保持与 FastAPI 一致）。
func (s *Service) List(ctx context.Context, q ListQuery) ([]Tag, int, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, 0, ErrUnavailable
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()
	return s.repo.List(qctx, q)
}

// invalidate 失效缓存：失败重试并告警，最终由 TTL 兜底。
func (s *Service) invalidate(ctx context.Context, nfcIDs ...string) {
	if len(nfcIDs) == 0 {
		return
	}
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		if lastErr = s.cache.Del(ctx, nfcIDs...); lastErr == nil {
			return
		}
		time.Sleep(time.Duration(50*(attempt+1)) * time.Millisecond)
	}
	s.log.Error("cache invalidate failed", "nfc_ids", nfcIDs, "err", lastErr)
}
