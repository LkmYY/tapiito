package tag

import (
	"encoding/json"
	"time"
)

// Tag 对应 nfc_tag 表。
type Tag struct {
	ID              int64
	NfcID           string
	Name            *string
	Phone           *string
	UserID          *int64
	CreatorName     *string
	Status          string
	TemplateType    string
	CustomTemplate  *string
	TemplatePayload json.RawMessage
	OrderNo         *string
	CreatedAt       *time.Time

	// 访问控制（标签级）
	AccessMode         string
	PasscodeHint       *string
	TicketTTLSeconds   int
	AutoCloseSeconds   int
	WatermarkEnabled   bool
	PasscodeHash       *string
}

// PublicPage 是公开展示 DTO：只含展示字段，不含手机号、user_id、审计信息。
type PublicPage struct {
	NfcID           string          `json:"nfc_id"`
	Status          string          `json:"status"`
	Name            *string         `json:"name"`
	TemplateType    *string         `json:"template_type"`
	CustomTemplate  *string         `json:"custom_template"`
	TemplatePayload json.RawMessage `json:"template_payload"`
}

// GalleryItem 是展厅卡片概览。
type GalleryItem struct {
	NfcID          string  `json:"nfc_id"`
	Name           string  `json:"name"`
	TemplateType   string  `json:"template_type"`
	CustomTemplate *string `json:"custom_template"`
	CreatedAt      *string `json:"created_at"`
}

// APITag 是管理端标签响应，字段与 FastAPI 的 NfcTagResp 严格对齐。
type APITag struct {
	ID              int64           `json:"id"`
	NfcID           string          `json:"nfc_id"`
	Name            *string         `json:"name"`
	Phone           *string         `json:"phone"`
	UserID          *int64          `json:"user_id"`
	CreatorName     *string         `json:"creator_name"`
	Status          string          `json:"status"`
	TemplateType    string          `json:"template_type"`
	CustomTemplate  *string         `json:"custom_template"`
	TemplatePayload json.RawMessage `json:"template_payload"`
	OrderNo         *string         `json:"order_no"`
	CreatedAt       *string         `json:"created_at"`

	// 访问控制：只回传配置，绝不回传 passcode_hash 明文
	AccessMode         string  `json:"access_mode"`
	HasPasscode        bool    `json:"has_passcode"`
	PasscodeHint       *string `json:"passcode_hint"`
	TicketTTLSeconds   int     `json:"ticket_ttl_seconds"`
	AutoCloseSeconds   int     `json:"auto_close_seconds"`
	WatermarkEnabled   bool    `json:"watermark_enabled"`
}

func (t *Tag) ToAPI() APITag {
	return APITag{
		ID:              t.ID,
		NfcID:           t.NfcID,
		Name:            t.Name,
		Phone:           t.Phone,
		UserID:          t.UserID,
		CreatorName:     t.CreatorName,
		Status:          t.Status,
		TemplateType:    t.TemplateType,
		CustomTemplate:  t.CustomTemplate,
		TemplatePayload: t.TemplatePayload,
		OrderNo:         t.OrderNo,
		CreatedAt:       FormatTime(t.CreatedAt),

		AccessMode:       t.AccessMode,
		HasPasscode:      t.PasscodeHash != nil && *t.PasscodeHash != "",
		PasscodeHint:     t.PasscodeHint,
		TicketTTLSeconds: t.TicketTTLSeconds,
		AutoCloseSeconds: t.AutoCloseSeconds,
		WatermarkEnabled: t.WatermarkEnabled,
	}
}

// ToPublic 生成公开展示 DTO；停用标签不下发内容。
func (t *Tag) ToPublic() PublicPage {
	if t.Status != "enabled" {
		return PublicPage{
			NfcID:  t.NfcID,
			Status: t.Status,
			Name:   t.Name,
		}
	}
	tt := t.TemplateType
	return PublicPage{
		NfcID:           t.NfcID,
		Status:          t.Status,
		Name:            t.Name,
		TemplateType:    &tt,
		CustomTemplate:  t.CustomTemplate,
		TemplatePayload: t.TemplatePayload,
	}
}

// FormatTime 输出与 FastAPI 一致的 ISO8601（本地时间、不带时区后缀）。
func FormatTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format("2006-01-02T15:04:05.999999")
	return &s
}
