package httpapi

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/lkmyy/tapiito-go/internal/auth"
	"github.com/lkmyy/tapiito-go/internal/tag"
)

// AdminHandler 承载管理入口的全部接口（仅监听管理端口）。
type AdminHandler struct {
	db     *sql.DB
	svc    *tag.Service
	issuer *auth.Issuer
	log    *slog.Logger
}

func NewAdminHandler(db *sql.DB, svc *tag.Service, issuer *auth.Issuer, log *slog.Logger) *AdminHandler {
	return &AdminHandler{db: db, svc: svc, issuer: issuer, log: log}
}

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login 对应 POST /api/auth/loginAdmin，返回结构与 FastAPI 一致。
func (h *AdminHandler) Login(c *gin.Context) {
	var req loginReq
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}

	var hash string
	err := h.db.QueryRowContext(c.Request.Context(),
		"SELECT password_hash FROM admin_user WHERE username = ? LIMIT 1", req.Username).Scan(&hash)
	if errors.Is(err, sql.ErrNoRows) || (err == nil && !auth.CheckPassword(hash, req.Password)) {
		// 不区分「用户不存在」与「密码错误」，避免枚举
		c.JSON(http.StatusUnauthorized, gin.H{"detail": "用户名或密码错误"})
		return
	}
	if err != nil {
		h.log.Error("admin login query failed", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
		return
	}

	token, _, err := h.issuer.Issue(req.Username)
	if err != nil {
		h.log.Error("issue token failed", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"token_type":   "bearer",
		"user_type":    "admin",
	})
}

// ListNfc 对应 GET /api/nfc。
func (h *AdminHandler) ListNfc(c *gin.Context) {
	page := parseInt(c.DefaultQuery("page", "1"), 1)
	pageSize := parseInt(c.DefaultQuery("page_size", "10"), 10)

	q := tag.ListQuery{
		Search:   c.Query("search"),
		Status:   c.Query("status"),
		Page:     page,
		PageSize: pageSize,
	}
	if v := c.Query("creator_id"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			q.CreatorID = &n
		}
	}

	items, total, err := h.svc.List(c.Request.Context(), q)
	if err != nil {
		h.fail(c, err)
		return
	}

	out := make([]tag.APITag, 0, len(items))
	for _, t := range items {
		out = append(out, t.ToAPI())
	}
	totalPages := 0
	if q.PageSize > 0 {
		totalPages = (total + q.PageSize - 1) / q.PageSize
	}
	c.JSON(http.StatusOK, gin.H{
		"items":       out,
		"total":       total,
		"page":        q.Page,
		"page_size":   q.PageSize,
		"total_pages": totalPages,
	})
}

// GetNfc 对应 GET /api/nfc/{id}。
func (h *AdminHandler) GetNfc(c *gin.Context) {
	id, ok := h.parseID(c)
	if !ok {
		return
	}
	t, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, t)
}

type nfcCreateReq struct {
	NfcID           string  `json:"nfc_id" binding:"required,max=128"`
	Name            *string `json:"name" binding:"omitempty,max=128"`
	Phone           *string `json:"phone" binding:"omitempty,max=32"`
	UserID          *int64  `json:"user_id"`
	Status          *string `json:"status"`
	TemplateType    string  `json:"template_type" binding:"required"`
	CustomTemplate  *string `json:"custom_template"`
	TemplatePayload any     `json:"template_payload"`

	// 访问控制
	AccessMode       string  `json:"access_mode" binding:"omitempty,oneof=none passcode ticket both"`
	Passcode         *string `json:"passcode" binding:"omitempty,max=64"`
	PasscodeHint     *string `json:"passcode_hint" binding:"omitempty,max=64"`
	TicketTTLSeconds int     `json:"ticket_ttl_seconds"`
	AutoCloseSeconds int     `json:"auto_close_seconds"`
	WatermarkEnabled bool    `json:"watermark_enabled"`
}

// CreateNfc 对应 POST /api/nfc。
func (h *AdminHandler) CreateNfc(c *gin.Context) {
	var req nfcCreateReq
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}
	status := "enabled"
	if req.Status != nil {
		status = *req.Status
	}
	in := tag.CreateInput{
		NfcID:          strings.TrimSpace(req.NfcID),
		Name:           req.Name,
		Phone:          req.Phone,
		UserID:         req.UserID,
		Status:         status,
		TemplateType:   req.TemplateType,
		CustomTemplate: req.CustomTemplate,
	}
	payload, err := marshalPayload(req.TemplatePayload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "invalid template_payload"})
		return
	}
	in.TemplatePayload = payload
	in.AccessMode = req.AccessMode
	in.PasscodeHint = req.PasscodeHint
	in.TicketTTLSeconds = req.TicketTTLSeconds
	in.AutoCloseSeconds = req.AutoCloseSeconds
	in.WatermarkEnabled = req.WatermarkEnabled
	if req.Passcode != nil && strings.TrimSpace(*req.Passcode) != "" {
		hash, err := tag.HashPasscode(*req.Passcode)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "暗号太短（至少 2 位）"})
			return
		}
		in.PasscodeHash = &hash
	}

	t, svcErr := h.svc.AdminCreate(c.Request.Context(), in)
	if svcErr != nil {
		h.fail(c, svcErr)
		return
	}
	c.JSON(http.StatusOK, t)
}

type nfcUpdateReq struct {
	Name            *string `json:"name"`
	Phone           *string `json:"phone"`
	Status          *string `json:"status"`
	TemplateType    *string `json:"template_type"`
	CustomTemplate  *string `json:"custom_template"`
	TemplatePayload any     `json:"template_payload"`

	// 访问控制：passcode 传空串表示清除暗号，不传表示不修改
	AccessMode       *string `json:"access_mode"`
	Passcode         *string `json:"passcode"`
	PasscodeHint     *string `json:"passcode_hint"`
	TicketTTLSeconds *int    `json:"ticket_ttl_seconds"`
	AutoCloseSeconds *int    `json:"auto_close_seconds"`
	WatermarkEnabled *bool   `json:"watermark_enabled"`
}

// UpdateNfc 对应 PUT /api/nfc/{id}。
func (h *AdminHandler) UpdateNfc(c *gin.Context) {
	id, ok := h.parseID(c)
	if !ok {
		return
	}
	var req nfcUpdateReq
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}
	in := tag.UpdateInput{
		Name:             req.Name,
		Phone:            req.Phone,
		Status:           req.Status,
		TemplateType:     req.TemplateType,
		CustomTemplate:   req.CustomTemplate,
		AccessMode:       req.AccessMode,
		PasscodeHint:     req.PasscodeHint,
		TicketTTLSeconds: req.TicketTTLSeconds,
		AutoCloseSeconds: req.AutoCloseSeconds,
		WatermarkEnabled: req.WatermarkEnabled,
	}
	// 暗号：不传 = 不改；传空串 = 清除；传值 = 重新哈希
	if req.Passcode != nil {
		plain := strings.TrimSpace(*req.Passcode)
		if plain == "" {
			empty := ""
			in.PasscodeHash = &empty
		} else {
			hash, err := tag.HashPasscode(plain)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "暗号太短（至少 2 位）"})
				return
			}
			in.PasscodeHash = &hash
		}
	}
	if req.TemplatePayload != nil {
		payload, err := marshalPayload(req.TemplatePayload)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "invalid template_payload"})
			return
		}
		in.TemplatePayload = payload
	}

	t, err := h.svc.AdminUpdate(c.Request.Context(), id, in)
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, t)
}

// DeleteNfc 对应 DELETE /api/nfc/{id}。
func (h *AdminHandler) DeleteNfc(c *gin.Context) {
	id, ok := h.parseID(c)
	if !ok {
		return
	}
	if err := h.svc.AdminDelete(c.Request.Context(), id); err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

type batchReq struct {
	IDs    []int64 `json:"ids" binding:"required,min=1"`
	Action string  `json:"action" binding:"required,oneof=enable disable delete"`
}

// BatchNfc 对应 POST /api/nfc/batch。
func (h *AdminHandler) BatchNfc(c *gin.Context) {
	var req batchReq
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}
	n, err := h.svc.BatchAction(c.Request.Context(), req.IDs, req.Action)
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "affected": n})
}

// Creators 对应 GET /api/nfc/creators。
func (h *AdminHandler) Creators(c *gin.Context) {
	items, err := h.svc.Creators(c.Request.Context())
	if err != nil {
		h.fail(c, err)
		return
	}
	// 与 FastAPI 保持一致：返回裸数组，前端直接赋给筛选项
	c.JSON(http.StatusOK, items)
}

// ---------- helpers ----------

func (h *AdminHandler) parseID(c *gin.Context) (int64, bool) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "invalid id"})
		return 0, false
	}
	return id, true
}

func (h *AdminHandler) fail(c *gin.Context, err error) {
	switch {
	case errors.Is(err, tag.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"detail": "Tag not found"})
	case errors.Is(err, tag.ErrUnavailable):
		c.Header("Retry-After", "1")
		c.JSON(http.StatusServiceUnavailable, gin.H{"detail": "service busy"})
	default:
		h.log.Error("admin request failed", "path", c.FullPath(), "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
	}
}

func parseInt(s string, def int) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return def
}
