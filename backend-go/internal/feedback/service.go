// Package feedback 实现用户意见：公开提交与展示、管理列表与采纳操作。
package feedback

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/lkmyy/tapiito-go/internal/dbgate"
)

var (
	ErrNotFound      = errors.New("feedback not found")
	ErrContentTooShort = errors.New("content too short")
)

// Item 对应 feedback 表。
type Item struct {
	ID        int64
	OrderNo   *string
	Content   string
	Status    string
	CreatedIP *string
	CreatedAt *time.Time
	AdoptedAt *time.Time
}

// APIItem 是管理端响应，字段与 FeedbackResp 对齐。
type APIItem struct {
	ID        int64   `json:"id"`
	OrderNo   *string `json:"order_no"`
	Content   string  `json:"content"`
	Status    string  `json:"status"`
	CreatedAt *string `json:"created_at"`
	AdoptedAt *string `json:"adopted_at"`
}

// PublicItem 是公开响应，不暴露订单号，字段与 AdoptedFeedbackResp 对齐。
type PublicItem struct {
	ID        int64   `json:"id"`
	Content   string  `json:"content"`
	CreatedAt *string `json:"created_at"`
	AdoptedAt *string `json:"adopted_at"`
}

func formatTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format("2006-01-02T15:04:05.999999")
	return &s
}

func (i *Item) ToAPI() APIItem {
	return APIItem{
		ID: i.ID, OrderNo: i.OrderNo, Content: i.Content,
		Status: i.Status, CreatedAt: formatTime(i.CreatedAt), AdoptedAt: formatTime(i.AdoptedAt),
	}
}

func (i *Item) ToPublic() PublicItem {
	return PublicItem{
		ID: i.ID, Content: i.Content,
		CreatedAt: formatTime(i.CreatedAt), AdoptedAt: formatTime(i.AdoptedAt),
	}
}

type Service struct {
	db           *sql.DB
	readGate     *dbgate.Gate
	writeGate    *dbgate.Gate
	queryTimeout time.Duration
	log          *slog.Logger
}

type Deps struct {
	DB           *sql.DB
	ReadGate     *dbgate.Gate
	WriteGate    *dbgate.Gate
	QueryTimeout time.Duration
	Logger       *slog.Logger
}

func NewService(d Deps) *Service {
	if d.QueryTimeout <= 0 {
		d.QueryTimeout = 3 * time.Second
	}
	return &Service{db: d.DB, readGate: d.ReadGate, writeGate: d.WriteGate,
		queryTimeout: d.QueryTimeout, log: d.Logger}
}

const selectCols = "id, order_no, content, `status`, created_ip, created_at, adopted_at"

type scanner interface{ Scan(dest ...any) error }

func scanItem(row scanner) (*Item, error) {
	it := &Item{}
	err := row.Scan(&it.ID, &it.OrderNo, &it.Content, &it.Status, &it.CreatedIP, &it.CreatedAt, &it.AdoptedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan feedback: %w", err)
	}
	return it, nil
}

// Create 提交意见（公开）：内容至少 5 个字，订单号可选。
func (s *Service) Create(ctx context.Context, orderNo, content, ip string) (*APIItem, error) {
	content = strings.TrimSpace(content)
	if len([]rune(content)) < 5 {
		return nil, ErrContentTooShort
	}
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()

	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var orderNoVal any
	if v := strings.TrimSpace(orderNo); v != "" {
		orderNoVal = v
	}
	res, err := s.db.ExecContext(qctx,
		"INSERT INTO feedback (order_no, content, `status`, created_ip, created_at) VALUES (?, ?, 'pending', ?, NOW())",
		orderNoVal, content, ip)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	row := s.db.QueryRowContext(qctx, "SELECT "+selectCols+" FROM feedback WHERE id = ?", id)
	it, err := scanItem(row)
	if err != nil {
		return nil, err
	}
	out := it.ToAPI()
	return &out, nil
}

// ListAdopted 公开的已采纳意见。
func (s *Service) ListAdopted(ctx context.Context, limit int) ([]PublicItem, error) {
	if limit < 1 {
		limit = 12
	}
	if limit > 50 {
		limit = 50
	}
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	rows, err := s.db.QueryContext(qctx, "SELECT "+selectCols+
		" FROM feedback WHERE `status` = 'adopted' ORDER BY adopted_at DESC, id DESC LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]PublicItem, 0, limit)
	for rows.Next() {
		it, err := scanItem(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, it.ToPublic())
	}
	return out, rows.Err()
}

// ListQuery 是管理端筛选条件。
type ListQuery struct {
	Page     int
	PageSize int
	Search   string
	Status   string
}

// Page 是管理端分页结果（含两个状态计数）。
type Page struct {
	Items        []APIItem
	Total        int
	PendingTotal int
	AdoptedTotal int
}

// List 管理端分页列表。
func (s *Service) List(ctx context.Context, q ListQuery) (*Page, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	where := "1=1"
	args := make([]any, 0, 3)
	if kw := strings.TrimSpace(q.Search); kw != "" {
		like := "%" + kw + "%"
		where += " AND (content LIKE ? OR order_no LIKE ?)"
		args = append(args, like, like)
	}
	if q.Status == "pending" || q.Status == "adopted" {
		where += " AND `status` = ?"
		args = append(args, q.Status)
	}

	var total int
	if err := s.db.QueryRowContext(qctx, "SELECT COUNT(*) FROM feedback WHERE "+where, args...).Scan(&total); err != nil {
		return nil, err
	}
	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 1 || q.PageSize > 100 {
		q.PageSize = 10
	}

	items := make([]APIItem, 0, q.PageSize)
	if total > 0 {
		rows, err := s.db.QueryContext(qctx, "SELECT "+selectCols+" FROM feedback WHERE "+where+
			" ORDER BY created_at DESC, id DESC LIMIT ? OFFSET ?",
			append(args, q.PageSize, (q.Page-1)*q.PageSize)...)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			it, err := scanItem(rows)
			if err != nil {
				return nil, err
			}
			items = append(items, it.ToAPI())
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}

	var pending, adopted int
	// COALESCE 兜底：空表时 SUM 返回 NULL，直接 Scan 到 int 会报错
	if err := s.db.QueryRowContext(qctx,
		"SELECT COALESCE(SUM(`status`='pending'),0), COALESCE(SUM(`status`='adopted'),0) FROM feedback").
		Scan(&pending, &adopted); err != nil {
		return nil, err
	}
	return &Page{Items: items, Total: total, PendingTotal: pending, AdoptedTotal: adopted}, nil
}

// Adopt 采纳/取消采纳。
func (s *Service) Adopt(ctx context.Context, id int64, adopted bool) (*APIItem, error) {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	status := "pending"
	var adoptedAt any
	if adopted {
		status = "adopted"
		adoptedAt = time.Now().UTC()
	}
	res, err := s.db.ExecContext(qctx,
		"UPDATE feedback SET `status` = ?, adopted_at = ? WHERE id = ?", status, adoptedAt, id)
	if err != nil {
		return nil, err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		// 可能是状态未变化，仍需确认记录存在
		var cnt int
		if err := s.db.QueryRowContext(qctx, "SELECT COUNT(*) FROM feedback WHERE id = ?", id).Scan(&cnt); err != nil {
			return nil, err
		}
		if cnt == 0 {
			return nil, ErrNotFound
		}
	}
	row := s.db.QueryRowContext(qctx, "SELECT "+selectCols+" FROM feedback WHERE id = ?", id)
	it, err := scanItem(row)
	if err != nil {
		return nil, err
	}
	out := it.ToAPI()
	return &out, nil
}
