// Package order 实现硬件订单：下单、查询、删除与「订单转标签」。
package order

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrNotFound     = errors.New("order not found")
	ErrConverted    = errors.New("order already converted")
	ErrCancelled    = errors.New("order cancelled")
	ErrNfcIDConflict = errors.New("nfc_id conflict")
	ErrGenerateFailed = errors.New("generate order no failed")
)

// Order 对应 order 表。
type Order struct {
	ID               int64
	OrderNo          string
	TemplateType     string
	CustomTemplate   *string
	TemplatePayload  json.RawMessage
	TagName          *string
	ExternalOrderNo  *string
	Quantity         int
	Amount           float64
	ContactName      string
	ContactPhone     string
	ContactAddress   string
	ContactZipcode   *string
	ContactNote      *string
	Status           string
	NfcTagID         *int64
	NfcID            *string
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}

// APIOrder 是响应结构，字段与 FastAPI 的 OrderResp 严格对齐。
type APIOrder struct {
	ID              int64           `json:"id"`
	OrderNo         string          `json:"order_no"`
	TemplateType    string          `json:"template_type"`
	CustomTemplate  *string         `json:"custom_template"`
	TemplatePayload json.RawMessage `json:"template_payload"`
	TagName         *string         `json:"tag_name"`
	ExternalOrderNo *string         `json:"external_order_no"`
	Quantity        int             `json:"quantity"`
	Amount          float64         `json:"amount"`
	ContactName     string          `json:"contact_name"`
	ContactPhone    string          `json:"contact_phone"`
	ContactAddress  string          `json:"contact_address"`
	ContactZipcode  *string         `json:"contact_zipcode"`
	ContactNote     *string         `json:"contact_note"`
	Status          string          `json:"status"`
	NfcTagID        *int64          `json:"nfc_tag_id"`
	NfcID           *string         `json:"nfc_id"`
	CreatedAt       *string         `json:"created_at"`
	UpdatedAt       *string         `json:"updated_at"`
}

func formatTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format("2006-01-02T15:04:05.999999")
	return &s
}

func (o *Order) ToAPI() APIOrder {
	return APIOrder{
		ID: o.ID, OrderNo: o.OrderNo, TemplateType: o.TemplateType,
		CustomTemplate: o.CustomTemplate, TemplatePayload: o.TemplatePayload,
		TagName: o.TagName, ExternalOrderNo: o.ExternalOrderNo,
		Quantity: o.Quantity, Amount: o.Amount,
		ContactName: o.ContactName, ContactPhone: o.ContactPhone,
		ContactAddress: o.ContactAddress, ContactZipcode: o.ContactZipcode,
		ContactNote: o.ContactNote, Status: o.Status,
		NfcTagID: o.NfcTagID, NfcID: o.NfcID,
		CreatedAt: formatTime(o.CreatedAt), UpdatedAt: formatTime(o.UpdatedAt),
	}
}

// CreateInput 与 FastAPI 的 OrderCreate 对齐。
type CreateInput struct {
	TemplateType    string
	CustomTemplate  *string
	TemplatePayload json.RawMessage
	TagName         *string
	ExternalOrderNo *string
	Quantity        int
	Amount          float64
	ContactName     string
	ContactPhone    string
	ContactAddress  string
	ContactZipcode  *string
	ContactNote     *string
}

// ConvertInput 与 FastAPI 的 OrderConvertReq 对齐。
type ConvertInput struct {
	NfcID *string
	Name  *string
	Phone *string
}
