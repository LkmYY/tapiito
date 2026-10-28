package httpapi

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/lkmyy/tapiito-go/internal/tag"
	"github.com/lkmyy/tapiito-go/internal/upload"
	"github.com/redis/go-redis/v9"
)

var nfcIDPattern = regexp.MustCompile(`^[A-Za-z0-9_-]{1,128}$`)

// PublicHandler 承载公开接口（读 + 公开写）。
type PublicHandler struct {
	svc      *tag.Service
	upload   *upload.Store
	sendSMS  func(ctx context.Context, to, body string, media []string) bool
	baseURL  string
	log      *slog.Logger
	attempts *redis.Client // 暗号失败计数（IP + 标签双维度）
}

// 暗号防爆破：连续失败 3 次锁 30 分钟
const (
	maxUnlockAttempts = 3
)

var unlockLockTTL = 30 * time.Minute

func NewPublicHandler(svc *tag.Service, up *upload.Store,
	sendSMS func(ctx context.Context, to, body string, media []string) bool,
	baseURL string, attempts *redis.Client, log *slog.Logger) *PublicHandler {
	return &PublicHandler{
		svc: svc, upload: up, sendSMS: sendSMS,
		baseURL: baseURL, attempts: attempts, log: log,
	}
}

// Page 对应 GET /api/public/nfc/{nfc_id}/page。
func (h *PublicHandler) Page(c *gin.Context) {
	nfcID := strings.TrimSpace(c.Param("nfc_id"))
	if !nfcIDPattern.MatchString(nfcID) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "invalid nfc_id"})
		return
	}

	ctx := c.Request.Context()
	// 访问控制：none 模式零额外开销，其余模式校验 ticket / 要求暗号
	sessionID := c.GetHeader("X-Access-Session")
	ticket := strings.TrimSpace(c.Query("ticket"))

	cfg, accErr := h.svc.VerifyAccess(ctx, nfcID, ticket, sessionID)
	switch {
	case accErr == nil:
		// 放行
	case errors.Is(accErr, tag.ErrPasscodeRequired):
		c.JSON(http.StatusForbidden, gin.H{
			"detail":   "passcode required",
			"required": "passcode",
			"hint":     cfg.PasscodeHint,
		})
		return
	case errors.Is(accErr, tag.ErrTicketRequired):
		// 求助类（ticket 模式）：扫码即发，不阻塞善意路人
		if cfg.Mode == tag.AccessTicket {
			t, _, _, err := h.svc.IssueOnScan(ctx, nfcID, ClientIP(c))
			if err == nil {
				ticket = t
				break
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"detail": "link expired", "required": "ticket"})
		return
	case errors.Is(accErr, tag.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"detail": "Tag not found"})
		return
	default:
		h.log.Error("verify access failed", "nfc_id", nfcID, "err", accErr)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
		return
	}

	page, err := h.svc.PublicPage(ctx, nfcID, ClientIP(c),
		c.GetHeader("User-Agent"), c.GetHeader("Referer"))
	if err != nil {
		switch {
		case errors.Is(err, tag.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"detail": "Tag not found"})
		case errors.Is(err, tag.ErrUnavailable):
			c.Header("Retry-After", "1")
			c.JSON(http.StatusServiceUnavailable, gin.H{"detail": "service busy"})
		default:
			h.log.Error("public page failed", "nfc_id", nfcID, "err", err)
			c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"nfc_id":             page.NfcID,
		"status":             page.Status,
		"name":               page.Name,
		"template_type":      page.TemplateType,
		"custom_template":    page.CustomTemplate,
		"template_payload":   page.TemplatePayload,
		"watermark":          cfg.Watermark,
		"auto_close_seconds": cfg.AutoCloseSecond,
		"ticket":             ticket,
	})
}

// Unlock 提交暗号换取 ticket。同一 IP / 同一标签连续失败 3 次锁 30 分钟。
func (h *PublicHandler) Unlock(c *gin.Context) {
	nfcID := strings.TrimSpace(c.Param("nfc_id"))
	if !nfcIDPattern.MatchString(nfcID) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "invalid nfc_id"})
		return
	}
	var req struct {
		Passcode string `json:"passcode" binding:"required,max=64"`
	}
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}

	ip := ClientIP(c)
	// 双维度防爆破：IP 与标签各计一次失败
	ipKey := "unlock:ip:" + ip
	tagKey := "unlock:tag:" + nfcID

	if locked, retry := h.attemptLocked(c, ipKey, tagKey); locked {
		c.JSON(http.StatusTooManyRequests, gin.H{"detail": "尝试次数过多，请稍后再试", "retry_after": retry})
		return
	}

	ticket, expires, _, err := h.svc.Unlock(c.Request.Context(), nfcID, req.Passcode, ip)
	if err != nil {
		switch {
		case errors.Is(err, tag.ErrPasscodeWrong):
			remaining := h.recordFailure(c, ipKey, tagKey)
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "暗号不正确", "remaining": remaining})
		case errors.Is(err, tag.ErrPasscodeMissing):
			c.JSON(http.StatusBadRequest, gin.H{"detail": "该标签未设置暗号"})
		case errors.Is(err, tag.ErrNotFound):
			c.JSON(http.StatusNotFound, gin.H{"detail": "Tag not found"})
		default:
			h.log.Error("unlock failed", "nfc_id", nfcID, "err", err)
			c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
		}
		return
	}

	h.clearFailures(c, ipKey, tagKey)
	c.JSON(http.StatusOK, gin.H{
		"ticket":     ticket,
		"expires_at": expires.Format("2006-01-02T15:04:05"),
	})
}

// 失败计数在 Redis 里做（带 TTL），避免写放大到 MySQL。
func (h *PublicHandler) attemptLocked(c *gin.Context, keys ...string) (bool, int) {
	if h.attempts == nil {
		return false, 0
	}
	ctx := c.Request.Context()
	for _, k := range keys {
		if n, err := h.attempts.Get(ctx, k).Int(); err == nil && n >= maxUnlockAttempts {
			ttl := h.attempts.TTL(ctx, k).Val()
			return true, int(ttl.Seconds())
		}
	}
	return false, 0
}

func (h *PublicHandler) recordFailure(c *gin.Context, keys ...string) int {
	if h.attempts == nil {
		return 0
	}
	ctx := c.Request.Context()
	cfg := unlockLockTTL
	remaining := 0
	for _, k := range keys {
		n, _ := h.attempts.Incr(ctx, k).Result()
		h.attempts.Expire(ctx, k, cfg)
		if int(n) > remaining {
			remaining = int(n)
		}
	}
	left := maxUnlockAttempts - remaining
	if left < 0 {
		left = 0
	}
	return left
}

func (h *PublicHandler) clearFailures(c *gin.Context, keys ...string) {
	if h.attempts == nil {
		return
	}
	ctx := c.Request.Context()
	for _, k := range keys {
		h.attempts.Del(ctx, k)
	}
}

// Gallery 对应 GET /api/public/gallery。
func (h *PublicHandler) Gallery(c *gin.Context) {
	limit := 24
	if v := c.Query("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = n
		}
	}
	items, err := h.svc.Gallery(c.Request.Context(), limit)
	if err != nil {
		if errors.Is(err, tag.ErrUnavailable) {
			c.Header("Retry-After", "1")
			c.JSON(http.StatusServiceUnavailable, gin.H{"detail": "service busy"})
			return
		}
		h.log.Error("gallery failed", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
		return
	}
	// 与 FastAPI 一致：{"items": [...], "total": n}
	c.JSON(http.StatusOK, gin.H{"items": items, "total": len(items)})
}

// Health 是健康检查（同时报告依赖状态，供 K8s/Nginx 探活）。
func Health(checks ...func() error) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := http.StatusOK
		detail := gin.H{"status": "ok"}
		for i, check := range checks {
			if err := check(); err != nil {
				status = http.StatusServiceUnavailable
				detail["status"] = "degraded"
				detail["dependency_"+strconv.Itoa(i)] = err.Error()
			}
		}
		c.JSON(status, detail)
	}
}

// CORSMiddleware 处理跨域（仅公开入口需要，管理入口建议同域）。
func CORSMiddleware(origins []string) gin.HandlerFunc {
	allowAll := len(origins) == 0
	for _, o := range origins {
		if o == "*" {
			allowAll = true
		}
	}
	allowed := make(map[string]struct{}, len(origins))
	for _, o := range origins {
		allowed[o] = struct{}{}
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			if allowAll {
				c.Header("Access-Control-Allow-Origin", "*")
			} else if _, ok := allowed[origin]; ok {
				c.Header("Access-Control-Allow-Origin", origin)
				c.Header("Vary", "Origin")
			}
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Request-ID")
			c.Header("Access-Control-Max-Age", "600")
		}
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
