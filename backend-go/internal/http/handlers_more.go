package httpapi

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/lkmyy/tapiito-go/internal/feedback"
	"github.com/lkmyy/tapiito-go/internal/order"
	"github.com/lkmyy/tapiito-go/internal/stats"
	"github.com/lkmyy/tapiito-go/internal/tag"
)

// ---------- 公开写：留言 / 照片 / 短信提交 ----------

type noteReq struct {
	Content string `json:"content" binding:"required,max=2000"`
}

// Note 对应 POST /api/public/nfc/{nfc_id}/note。
func (h *PublicHandler) Note(c *gin.Context) {
	var req noteReq
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}
	out, err := h.svc.AddNote(c.Request.Context(), c.Param("nfc_id"), req.Content, ClientIP(c))
	if err != nil {
		h.failPublic(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

// Photo 对应 POST /api/public/nfc/{nfc_id}/photo（multipart）。
func (h *PublicHandler) Photo(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "file required"})
		return
	}
	if h.upload == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"detail": "upload disabled"})
		return
	}
	if err := h.upload.Validate(file.Size, file.Header.Get("Content-Type")); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "cannot read file"})
		return
	}
	defer src.Close()
	data, err := io.ReadAll(io.LimitReader(src, file.Size+1))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "cannot read file"})
		return
	}

	key, err := h.upload.Save(c.Request.Context(), file.Filename, file.Header.Get("Content-Type"), data)
	if err != nil {
		h.log.Error("upload failed", "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "upload failed"})
		return
	}
	out, err := h.svc.AddPhoto(c.Request.Context(), c.Param("nfc_id"), key,
		file.Header.Get("Content-Type"), file.Size, ClientIP(c))
	if err != nil {
		h.failPublic(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

// Photos 对应 GET /api/public/nfc/{nfc_id}/photos。
func (h *PublicHandler) Photos(c *gin.Context) {
	out, err := h.svc.ListPhotos(c.Request.Context(), c.Param("nfc_id"))
	if err != nil {
		h.failPublic(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

type submitReq struct {
	Content  string  `json:"content" binding:"required,max=2000"`
	PhotoIDs []int64 `json:"photo_ids"`
}

// Submit 对应 POST /api/public/nfc/{nfc_id}/submit。
func (h *PublicHandler) Submit(c *gin.Context) {
	var req submitReq
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}
	base := h.baseURL
	if base == "" {
		base = "http://" + c.Request.Host
	}
	out, err := h.svc.Submit(c.Request.Context(), c.Param("nfc_id"), req.Content, req.PhotoIDs,
		h.sendSMS, base)
	if err != nil {
		h.failPublic(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *PublicHandler) failPublic(c *gin.Context, err error) {
	switch {
	case errors.Is(err, tag.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"detail": "Tag not found"})
	case errors.Is(err, tag.ErrUnavailable):
		c.Header("Retry-After", "1")
		c.JSON(http.StatusServiceUnavailable, gin.H{"detail": "service busy"})
	case errors.Is(err, tag.ErrPhoneNotConfigured):
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Phone not configured"})
	default:
		h.log.Error("public write failed", "path", c.FullPath(), "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
	}
}

// ---------- 订单 ----------

// OrderHandler 同时承载公开下单与管理端订单管理。
type OrderHandler struct {
	svc *order.Service
	log *slog.Logger
}

func NewOrderHandler(svc *order.Service, log *slog.Logger) *OrderHandler {
	return &OrderHandler{svc: svc, log: log}
}

type orderCreateReq struct {
	TemplateType    string  `json:"template_type" binding:"required"`
	CustomTemplate  *string `json:"custom_template"`
	TemplatePayload any     `json:"template_payload"`
	TagName         *string `json:"tag_name" binding:"omitempty,max=128"`
	ExternalOrderNo *string `json:"external_order_no" binding:"omitempty,max=64"`
	Quantity        int     `json:"quantity"`
	Amount          float64 `json:"amount"`
	ContactName     string  `json:"contact_name" binding:"required,min=1,max=64"`
	ContactPhone    string  `json:"contact_phone" binding:"required,min=1,max=32"`
	ContactAddress  string  `json:"contact_address" binding:"required,min=1,max=256"`
	ContactZipcode  *string `json:"contact_zipcode" binding:"omitempty,max=16"`
	ContactNote     *string `json:"contact_note" binding:"omitempty,max=256"`
}

// Create 对应 POST /api/orders（公开，无需登录）。
func (h *OrderHandler) Create(c *gin.Context) {
	var req orderCreateReq
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}
	payload, err := marshalPayload(req.TemplatePayload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "invalid template_payload"})
		return
	}
	qty := req.Quantity
	if qty < 1 {
		qty = 1
	}
	if qty > 9999 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "quantity too large"})
		return
	}
	out, err := h.svc.Create(c.Request.Context(), order.CreateInput{
		TemplateType: req.TemplateType, CustomTemplate: req.CustomTemplate,
		TemplatePayload: payload, TagName: req.TagName, ExternalOrderNo: req.ExternalOrderNo,
		Quantity: qty, Amount: req.Amount,
		ContactName: strings.TrimSpace(req.ContactName),
		ContactPhone: strings.TrimSpace(req.ContactPhone),
		ContactAddress: strings.TrimSpace(req.ContactAddress),
		ContactZipcode: req.ContactZipcode, ContactNote: req.ContactNote,
	})
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

// List 对应 GET /api/orders。
func (h *OrderHandler) List(c *gin.Context) {
	q := order.ListQuery{
		Page:     parseInt(c.DefaultQuery("page", "1"), 1),
		PageSize: parseInt(c.DefaultQuery("page_size", "10"), 10),
		Search:   c.Query("search"),
		Status:   c.Query("status"),
	}
	items, total, err := h.svc.List(c.Request.Context(), q)
	if err != nil {
		h.fail(c, err)
		return
	}
	totalPages := 0
	if total > 0 && q.PageSize > 0 {
		totalPages = (total + q.PageSize - 1) / q.PageSize
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items, "total": total, "page": q.Page, "page_size": q.PageSize, "total_pages": totalPages,
	})
}

// Get 对应 GET /api/orders/{order_id}。
func (h *OrderHandler) Get(c *gin.Context) {
	id, ok := parsePathID(c, "order_id")
	if !ok {
		return
	}
	out, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

type convertReq struct {
	NfcID *string `json:"nfc_id" binding:"omitempty,max=128"`
	Name  *string `json:"name" binding:"omitempty,max=128"`
	Phone *string `json:"phone" binding:"omitempty,max=32"`
}

// Convert 对应 POST /api/orders/{order_id}/convert。
func (h *OrderHandler) Convert(c *gin.Context) {
	id, ok := parsePathID(c, "order_id")
	if !ok {
		return
	}
	var req convertReq
	// body 允许为空
	if c.Request.ContentLength > 0 {
		if err := bindJSON(c, &req); err != nil {
			bindError(c, err)
			return
		}
	}
	out, err := h.svc.Convert(c.Request.Context(), id,
		order.ConvertInput{NfcID: req.NfcID, Name: req.Name, Phone: req.Phone})
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

// Delete 对应 DELETE /api/orders/{order_id}。
func (h *OrderHandler) Delete(c *gin.Context) {
	id, ok := parsePathID(c, "order_id")
	if !ok {
		return
	}
	if _, err := h.svc.Delete(c.Request.Context(), id); err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *OrderHandler) fail(c *gin.Context, err error) {
	switch {
	case errors.Is(err, order.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"detail": "Order not found"})
	case errors.Is(err, order.ErrConverted):
		c.JSON(http.StatusBadRequest, gin.H{"detail": "该订单已转换成标签"})
	case errors.Is(err, order.ErrCancelled):
		c.JSON(http.StatusBadRequest, gin.H{"detail": "已取消的订单不能转换"})
	case errors.Is(err, order.ErrNfcIDConflict):
		c.JSON(http.StatusConflict, gin.H{"detail": "NFC ID 已存在"})
	default:
		h.log.Error("order request failed", "path", c.FullPath(), "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
	}
}

// ---------- 反馈 ----------

type FeedbackHandler struct {
	svc *feedback.Service
	log *slog.Logger
}

func NewFeedbackHandler(svc *feedback.Service, log *slog.Logger) *FeedbackHandler {
	return &FeedbackHandler{svc: svc, log: log}
}

type feedbackCreateReq struct {
	OrderNo *string `json:"order_no" binding:"omitempty,max=64"`
	Content string  `json:"content" binding:"required,max=2000"`
}

// Create 对应 POST /api/feedback（公开）。
func (h *FeedbackHandler) Create(c *gin.Context) {
	var req feedbackCreateReq
	if err := bindJSON(c, &req); err != nil {
		bindError(c, err)
		return
	}
	orderNo := ""
	if req.OrderNo != nil {
		orderNo = *req.OrderNo
	}
	out, err := h.svc.Create(c.Request.Context(), orderNo, req.Content, ClientIP(c))
	if errors.Is(err, feedback.ErrContentTooShort) {
		// 与 Pydantic 的 min_length 拦截保持同构，并回填实际输入
		validationError(c, "content", "string_too_short",
			"String should have at least 5 characters", req.Content, map[string]any{"min_length": 5})
		return
	}
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

// Adopted 对应 GET /api/feedback/adopted（公开）。
func (h *FeedbackHandler) Adopted(c *gin.Context) {
	limit := parseInt(c.DefaultQuery("limit", "12"), 12)
	out, err := h.svc.ListAdopted(c.Request.Context(), limit)
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

// AdminList 对应 GET /api/feedback/admin。
func (h *FeedbackHandler) AdminList(c *gin.Context) {
	q := feedback.ListQuery{
		Page:     parseInt(c.DefaultQuery("page", "1"), 1),
		PageSize: parseInt(c.DefaultQuery("page_size", "10"), 10),
		Search:   c.Query("search"),
		Status:   c.Query("status"),
	}
	page, err := h.svc.List(c.Request.Context(), q)
	if err != nil {
		h.fail(c, err)
		return
	}
	totalPages := 1
	if page.Total > 0 && q.PageSize > 0 {
		totalPages = (page.Total + q.PageSize - 1) / q.PageSize
	}
	c.JSON(http.StatusOK, gin.H{
		"items": page.Items, "total": page.Total, "page": q.Page, "page_size": q.PageSize,
		"total_pages": totalPages, "pending_total": page.PendingTotal, "adopted_total": page.AdoptedTotal,
	})
}

type adoptReq struct {
	Adopted bool `json:"adopted"`
}

// Adopt 对应 PATCH /api/feedback/admin/{feedback_id}。
func (h *FeedbackHandler) Adopt(c *gin.Context) {
	id, ok := parsePathID(c, "feedback_id")
	if !ok {
		return
	}
	var req adoptReq
	if c.Request.ContentLength > 0 {
		if err := bindJSON(c, &req); err != nil {
			bindError(c, err)
			return
		}
	}
	out, err := h.svc.Adopt(c.Request.Context(), id, req.Adopted)
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *FeedbackHandler) fail(c *gin.Context, err error) {
	switch {
	case errors.Is(err, feedback.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"detail": "意见不存在"})
	case errors.Is(err, feedback.ErrContentTooShort):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "content too short"})
	default:
		h.log.Error("feedback request failed", "path", c.FullPath(), "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
	}
}

// ---------- 统计 ----------

type StatsHandler struct {
	svc *stats.Service
	log *slog.Logger
}

func NewStatsHandler(svc *stats.Service, log *slog.Logger) *StatsHandler {
	return &StatsHandler{svc: svc, log: log}
}

// Overview 对应 GET /api/stats/overview。
func (h *StatsHandler) Overview(c *gin.Context) {
	out, err := h.svc.Overview(c.Request.Context(), c.Query("start_date"), c.Query("end_date"))
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

// ByNfc 对应 GET /api/stats/nfc/{nfc_id}。
func (h *StatsHandler) ByNfc(c *gin.Context) {
	out, err := h.svc.ByNfcID(c.Request.Context(), c.Param("nfc_id"),
		c.Query("start_date"), c.Query("end_date"))
	if err != nil {
		h.fail(c, err)
		return
	}
	c.JSON(http.StatusOK, out)
}

// Export 对应 GET /api/stats/export（json / csv）。
func (h *StatsHandler) Export(c *gin.Context) {
	rows, err := h.svc.Export(c.Request.Context(), c.Query("start_date"), c.Query("end_date"))
	if err != nil {
		h.fail(c, err)
		return
	}
	if c.DefaultQuery("format", "json") == "csv" {
		filename := fmt.Sprintf("stats_%s.csv", time.Now().Format("20060102_150405"))
		c.Header("Content-Type", "text/csv; charset=utf-8")
		c.Header("Content-Disposition", "attachment; filename="+filename)
		w := csv.NewWriter(c.Writer)
		_ = w.Write([]string{"nfc_id", "ip_address", "user_agent", "referer", "country", "region", "city", "created_at"})
		for _, r := range rows {
			_ = w.Write([]string{r.NfcID, derefStr(r.IPAddress), derefStr(r.UserAgent),
				derefStr(r.Referer), derefStr(r.Country), derefStr(r.Region),
				derefStr(r.City), derefStr(r.CreatedAt)})
		}
		w.Flush()
		return
	}
	c.JSON(http.StatusOK, gin.H{"stats": rows})
}

func (h *StatsHandler) fail(c *gin.Context, err error) {
	switch {
	case errors.Is(err, stats.ErrTagNotFound):
		c.JSON(http.StatusNotFound, gin.H{"detail": "Tag not found"})
	default:
		h.log.Error("stats request failed", "path", c.FullPath(), "err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"detail": "internal error"})
	}
}

// ---------- helpers ----------

func parsePathID(c *gin.Context, name string) (int64, bool) {
	id, err := strconv.ParseInt(c.Param(name), 10, 64)
	if err != nil || id <= 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "invalid id"})
		return 0, false
	}
	return id, true
}

func derefStr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
