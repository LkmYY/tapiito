package order

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/lkmyy/tapiito-go/internal/bloom"
	"github.com/lkmyy/tapiito-go/internal/cache"
	"github.com/lkmyy/tapiito-go/internal/dbgate"
)

const (
	orderPrefix    = "TP"
	dailySeqWidth  = 4
	maxRetryNo     = 20
)

var cst = time.FixedZone("CST", 8*3600)

type Service struct {
	db           *sql.DB
	readGate     *dbgate.Gate
	writeGate    *dbgate.Gate
	queryTimeout time.Duration
	bloom        *bloom.Filter
	cache        *cache.PageCache
	log          *slog.Logger
}

type Deps struct {
	DB           *sql.DB
	ReadGate     *dbgate.Gate
	WriteGate    *dbgate.Gate
	QueryTimeout time.Duration
	Bloom        *bloom.Filter
	Cache        *cache.PageCache
	Logger       *slog.Logger
}

func NewService(d Deps) *Service {
	if d.QueryTimeout <= 0 {
		d.QueryTimeout = 3 * time.Second
	}
	return &Service{
		db: d.DB, readGate: d.ReadGate, writeGate: d.WriteGate,
		queryTimeout: d.QueryTimeout, bloom: d.Bloom, cache: d.Cache, log: d.Logger,
	}
}

const selectCols = "id, order_no, template_type, custom_template, template_payload, tag_name, " +
	"external_order_no, quantity, amount, contact_name, contact_phone, contact_address, " +
	"contact_zipcode, contact_note, `status`, nfc_tag_id, nfc_id, created_at, updated_at"

type scanner interface{ Scan(dest ...any) error }

func scanOrder(row scanner) (*Order, error) {
	o := &Order{}
	err := row.Scan(&o.ID, &o.OrderNo, &o.TemplateType, &o.CustomTemplate, &o.TemplatePayload,
		&o.TagName, &o.ExternalOrderNo, &o.Quantity, &o.Amount, &o.ContactName, &o.ContactPhone,
		&o.ContactAddress, &o.ContactZipcode, &o.ContactNote, &o.Status,
		&o.NfcTagID, &o.NfcID, &o.CreatedAt, &o.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan order: %w", err)
	}
	return o, nil
}

// Create 下单：生成订单号并落库（并发重号靠唯一索引 + 重试兜底）。
func (s *Service) Create(ctx context.Context, in CreateInput) (*APIOrder, error) {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var lastErr error
	for attempt := 0; attempt < maxRetryNo; attempt++ {
		no, err := s.nextOrderNo(qctx)
		if err != nil {
			return nil, err
		}
		o, err := s.insertOrder(qctx, no, in)
		if err == nil {
			out := o.ToAPI()
			return &out, nil
		}
		lastErr = err
		if !isDuplicate(err) {
			return nil, err
		}
	}
	return nil, fmt.Errorf("create order after retries: %w", lastErr)
}

func (s *Service) nextOrderNo(ctx context.Context) (string, error) {
	prefix := orderPrefix + time.Now().In(cst).Format("20060102")
	var last string
	err := s.db.QueryRowContext(ctx,
		"SELECT order_no FROM `order` WHERE order_no LIKE ? ORDER BY order_no DESC LIMIT 1",
		prefix+"%").Scan(&last)
	seq := 0
	if err == nil && len(last) > len(prefix) {
		if n, convErr := strconv.Atoi(strings.TrimSpace(last[len(prefix):])); convErr == nil {
			seq = n
		}
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}
	return fmt.Sprintf("%s%0*d", prefix, dailySeqWidth, seq+1), nil
}

func (s *Service) insertOrder(ctx context.Context, no string, in CreateInput) (*Order, error) {
	status := "pending"
	res, err := s.db.ExecContext(ctx, "INSERT INTO `order` "+
		"(order_no, template_type, custom_template, template_payload, tag_name, external_order_no, "+
		"quantity, amount, contact_name, contact_phone, contact_address, contact_zipcode, contact_note, "+
		"`status`, created_at, updated_at) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())",
		no, in.TemplateType, in.CustomTemplate, in.TemplatePayload, in.TagName, in.ExternalOrderNo,
		in.Quantity, in.Amount, in.ContactName, in.ContactPhone, in.ContactAddress,
		in.ContactZipcode, in.ContactNote, status)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return s.GetByID(ctx, id)
}

func isDuplicate(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "Duplicate entry") ||
		strings.Contains(err.Error(), "1062"))
}

// GetByID 按主键取订单。
func (s *Service) GetByID(ctx context.Context, id int64) (*Order, error) {
	row := s.db.QueryRowContext(ctx, "SELECT "+selectCols+" FROM `order` WHERE id = ?", id)
	return scanOrder(row)
}

// Get 是带闸门与超时的对外查询。
func (s *Service) Get(ctx context.Context, id int64) (*APIOrder, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()
	o, err := s.GetByID(qctx, id)
	if err != nil {
		return nil, err
	}
	out := o.ToAPI()
	return &out, nil
}

// ListQuery 是订单列表筛选条件。
type ListQuery struct {
	Page     int
	PageSize int
	Search   string
	Status   string
}

// List 分页查询订单。
func (s *Service) List(ctx context.Context, q ListQuery) ([]APIOrder, int, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	where := "1=1"
	args := make([]any, 0, 6)
	if s := strings.TrimSpace(q.Search); s != "" {
		like := "%" + s + "%"
		where += " AND (order_no LIKE ? OR contact_name LIKE ? OR contact_phone LIKE ? OR tag_name LIKE ? OR external_order_no LIKE ?)"
		args = append(args, like, like, like, like, like)
	}
	if st := strings.TrimSpace(q.Status); st != "" {
		where += " AND `status` = ?"
		args = append(args, st)
	}

	var total int
	if err := s.db.QueryRowContext(qctx, "SELECT COUNT(*) FROM `order` WHERE "+where, args...).Scan(&total); err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return []APIOrder{}, 0, nil
	}

	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 1 || q.PageSize > 100 {
		q.PageSize = 10
	}
	rows, err := s.db.QueryContext(qctx, "SELECT "+selectCols+" FROM `order` WHERE "+where+
		" ORDER BY id DESC LIMIT ? OFFSET ?", append(args, q.PageSize, (q.Page-1)*q.PageSize)...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	out := make([]APIOrder, 0, q.PageSize)
	for rows.Next() {
		o, err := scanOrder(rows)
		if err != nil {
			return nil, 0, err
		}
		out = append(out, o.ToAPI())
	}
	return out, total, rows.Err()
}

// ConvertResult 是「转换成标签」的返回。
type ConvertResult struct {
	OK       bool    `json:"ok"`
	OrderNo  string  `json:"order_no"`
	NfcTagID int64   `json:"nfc_tag_id"`
	NfcID    string  `json:"nfc_id"`
	Name     *string `json:"name"`
}

// Convert 把订单转成 NFC 标签：事务内建标签并回填订单，提交后再更新 Bloom 与缓存。
func (s *Service) Convert(ctx context.Context, id int64, in ConvertInput) (*ConvertResult, error) {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(qctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()

	row := tx.QueryRowContext(qctx, "SELECT "+selectCols+" FROM `order` WHERE id = ? FOR UPDATE", id)
	o, err := scanOrder(row)
	if err != nil {
		return nil, err
	}
	if o.Status == "converted" && o.NfcTagID != nil {
		return nil, ErrConverted
	}
	if o.Status == "cancelled" {
		return nil, ErrCancelled
	}

	nfcID := strings.TrimSpace(deref(in.NfcID))
	if nfcID == "" {
		nfcID = o.OrderNo
	}
	var exists int
	if err := tx.QueryRowContext(qctx, "SELECT COUNT(*) FROM nfc_tag WHERE nfc_id = ?", nfcID).Scan(&exists); err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, ErrNfcIDConflict
	}

	name := strings.TrimSpace(deref(in.Name))
	if name == "" {
		name = strings.TrimSpace(deref(o.TagName))
	}
	if name == "" {
		name = "标签 " + o.OrderNo
	}
	phone := strings.TrimSpace(deref(in.Phone))
	if phone == "" {
		phone = o.ContactPhone
	}

	res, err := tx.ExecContext(qctx, "INSERT INTO nfc_tag "+
		"(nfc_id, name, phone, `status`, template_type, custom_template, template_payload, order_no, created_at, updated_at) "+
		"VALUES (?, ?, ?, 'enabled', ?, ?, ?, ?, NOW(), NOW())",
		nfcID, name, phone, o.TemplateType, o.CustomTemplate, o.TemplatePayload, o.OrderNo)
	if err != nil {
		return nil, err
	}
	tagID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	if _, err := tx.ExecContext(qctx,
		"UPDATE `order` SET `status` = 'converted', nfc_tag_id = ?, nfc_id = ?, updated_at = NOW() WHERE id = ?",
		tagID, nfcID, id); err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	// 事务已提交：新标签必须立刻可访问（补 Bloom + 清缓存）
	bg := context.WithoutCancel(ctx)
	if err := s.bloom.Add(bg, nfcID); err != nil {
		s.log.Error("bloom add failed after convert", "nfc_id", nfcID, "err", err)
	}
	if err := s.cache.Del(bg, nfcID); err != nil {
		s.log.Error("cache invalidate failed after convert", "nfc_id", nfcID, "err", err)
	}

	return &ConvertResult{OK: true, OrderNo: o.OrderNo, NfcTagID: tagID, NfcID: nfcID, Name: &name}, nil
}

// Delete 删除订单（已转换的标签保留）。
func (s *Service) Delete(ctx context.Context, id int64) (*string, error) {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var orderNo string
	err = s.db.QueryRowContext(qctx, "SELECT order_no FROM `order` WHERE id = ?", id).Scan(&orderNo)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if _, err := s.db.ExecContext(qctx, "DELETE FROM `order` WHERE id = ?", id); err != nil {
		return nil, err
	}
	return &orderNo, nil
}

func deref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
