package tag

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrNotFound = errors.New("tag not found")

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo { return &Repo{db: db} }

const selectCols = `
  t.id, t.nfc_id, t.name, t.phone, t.user_id, u.username AS creator_name,
  t.status, t.template_type, t.custom_template, t.template_payload,
  t.order_no, t.created_at,
  t.access_mode, t.passcode_hash, t.passcode_hint,
  t.ticket_ttl_seconds, t.auto_close_seconds, t.watermark_enabled
FROM nfc_tag t
LEFT JOIN ` + "`user`" + ` u ON u.id = t.user_id`

// GetByNfcID 按芯片 ID 读取标签（含创建人用户名）。
func (r *Repo) GetByNfcID(ctx context.Context, nfcID string) (*Tag, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT "+selectCols+" WHERE t.nfc_id = ? LIMIT 1", nfcID)
	return scanTag(row)
}

// GetByID 按主键读取。
func (r *Repo) GetByID(ctx context.Context, id int64) (*Tag, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT "+selectCols+" WHERE t.id = ? LIMIT 1", id)
	return scanTag(row)
}

// ListQuery 是管理端列表的筛选条件。
type ListQuery struct {
	Search    string
	Status    string
	CreatorID *int64 // 0 表示只看管理员创建（user_id IS NULL）
	Page      int
	PageSize  int
}

// List 返回分页结果。
func (r *Repo) List(ctx context.Context, q ListQuery) ([]Tag, int, error) {
	where, args := buildWhere(q)

	var total int
	if err := r.db.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM nfc_tag t WHERE "+where, args...).Scan(&total); err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return []Tag{}, 0, nil
	}

	if q.Page < 1 {
		q.Page = 1
	}
	if q.PageSize < 1 || q.PageSize > 100 {
		q.PageSize = 10
	}
	offset := (q.Page - 1) * q.PageSize

	rows, err := r.db.QueryContext(ctx,
		"SELECT "+selectCols+" WHERE "+where+" ORDER BY t.id DESC LIMIT ? OFFSET ?",
		append(args, q.PageSize, offset)...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	out := make([]Tag, 0, q.PageSize)
	for rows.Next() {
		t, err := scanTagRows(rows)
		if err != nil {
			return nil, 0, err
		}
		out = append(out, *t)
	}
	return out, total, rows.Err()
}

func buildWhere(q ListQuery) (string, []any) {
	conds := []string{"1=1"}
	args := make([]any, 0, 4)
	if s := strings.TrimSpace(q.Search); s != "" {
		like := "%" + s + "%"
		conds = append(conds,
			"(t.nfc_id LIKE ? OR t.phone LIKE ? OR t.name LIKE ? OR t.order_no LIKE ?)")
		args = append(args, like, like, like, like)
	}
	if st := strings.TrimSpace(q.Status); st != "" {
		conds = append(conds, "t.status = ?")
		args = append(args, st)
	}
	if q.CreatorID != nil {
		if *q.CreatorID == 0 {
			conds = append(conds, "t.user_id IS NULL")
		} else {
			conds = append(conds, "t.user_id = ?")
			args = append(args, *q.CreatorID)
		}
	}
	return strings.Join(conds, " AND "), args
}

// AllNfcIDs 供 Bloom 构建使用：只取 nfc_id，避免全表行宽浪费。
func (r *Repo) AllNfcIDs(ctx context.Context) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT nfc_id FROM nfc_tag")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := make([]string, 0, 1024)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, rows.Err()
}

// Gallery 返回展厅卡片（只含已启用标签的展示字段）。
func (r *Repo) Gallery(ctx context.Context, limit int) ([]GalleryItem, error) {
	if limit < 1 {
		limit = 24
	}
	if limit > 60 {
		limit = 60
	}
	rows, err := r.db.QueryContext(ctx, `
		SELECT nfc_id, name, template_type, custom_template, created_at, template_payload
		FROM nfc_tag
		WHERE status = 'enabled'
		ORDER BY id DESC
		LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make([]GalleryItem, 0, limit)
	for rows.Next() {
		var (
			nfcID, ttype string
			name, custom *string
			created      *time.Time
			payload      []byte
		)
		if err := rows.Scan(&nfcID, &name, &ttype, &custom, &created, &payload); err != nil {
			return nil, err
		}
		items = append(items, GalleryItem{
			NfcID:          nfcID,
			Name:           displayName(name, payloadName(payload)),
			TemplateType:   ttype,
			CustomTemplate: custom,
			CreatedAt:      FormatTime(created),
		})
	}
	return items, rows.Err()
}

// payloadName 复刻 FastAPI 展厅取名顺序，保证卡片名称一致。
func payloadName(payload []byte) string {
	if len(payload) == 0 {
		return ""
	}
	var m map[string]any
	if err := json.Unmarshal(payload, &m); err != nil {
		return ""
	}
	for _, k := range []string{"name", "shopName", "companyName", "salonName", "gymName", "schoolName", "title"} {
		if v, ok := m[k].(string); ok && strings.TrimSpace(v) != "" {
			return v
		}
	}
	return ""
}

// displayName 复刻 FastAPI 的取名兜底顺序，保证展厅卡片名称一致。
func displayName(name *string, payloadName string) string {
	if name != nil && strings.TrimSpace(*name) != "" {
		return *name
	}
	if payloadName != "" {
		return payloadName
	}
	return "未命名标签"
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanTag(row rowScanner) (*Tag, error) {
	t := &Tag{}
	err := row.Scan(
		&t.ID, &t.NfcID, &t.Name, &t.Phone, &t.UserID, &t.CreatorName,
		&t.Status, &t.TemplateType, &t.CustomTemplate, &t.TemplatePayload,
		&t.OrderNo, &t.CreatedAt,
		&t.AccessMode, &t.PasscodeHash, &t.PasscodeHint,
		&t.TicketTTLSeconds, &t.AutoCloseSeconds, &t.WatermarkEnabled,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("scan tag: %w", err)
	}
	return t, nil
}

func scanTagRows(rows *sql.Rows) (*Tag, error) { return scanTag(rows) }

// CreateInput 是创建标签的入参（字段与 FastAPI 的 NfcTagCreate 对齐）。
type CreateInput struct {
	NfcID           string
	Name            *string
	Phone           *string
	UserID          *int64
	Status          string
	TemplateType    string
	CustomTemplate  *string
	TemplatePayload json.RawMessage

	// 访问控制
	AccessMode         string // none / passcode / ticket / both
	PasscodeHash       *string
	PasscodeHint       *string
	TicketTTLSeconds   int
	AutoCloseSeconds   int
	WatermarkEnabled   bool
}

// Create 插入新标签并返回落库结果。
func (r *Repo) Create(ctx context.Context, in CreateInput) (*Tag, error) {
	status := in.Status
	if status == "" {
		status = "enabled"
	}
	mode := in.AccessMode
	if mode == "" {
		mode = AccessNone
	}
	ttl := in.TicketTTLSeconds
	if ttl <= 0 {
		ttl = 3600
	}
	mark := 0
	if in.WatermarkEnabled {
		mark = 1
	}
	res, err := r.db.ExecContext(ctx, `
		INSERT INTO nfc_tag (nfc_id, name, phone, user_id, status, template_type, custom_template, template_payload,
			access_mode, passcode_hash, passcode_hint, ticket_ttl_seconds, auto_close_seconds, watermark_enabled,
			created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`,
		in.NfcID, in.Name, in.Phone, in.UserID, status,
		in.TemplateType, in.CustomTemplate, in.TemplatePayload,
		mode, in.PasscodeHash, in.PasscodeHint, ttl, in.AutoCloseSeconds, mark)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.GetByID(ctx, id)
}

// UpdateInput 是更新入参；nil 表示该字段不修改。
type UpdateInput struct {
	Name            *string
	Phone           *string
	Status          *string
	TemplateType    *string
	CustomTemplate  *string
	TemplatePayload json.RawMessage

	// 访问控制：nil = 不修改；PasscodeHash 传空指针值 "" 表示清除暗号
	AccessMode       *string
	PasscodeHash     *string
	PasscodeHint     *string
	TicketTTLSeconds *int
	AutoCloseSeconds *int
	WatermarkEnabled *bool
}

// Update 按主键更新并返回最新记录。
func (r *Repo) Update(ctx context.Context, id int64, in UpdateInput) (*Tag, error) {
	sets := make([]string, 0, 6)
	args := make([]any, 0, 8)
	if in.Name != nil {
		sets = append(sets, "name = ?")
		args = append(args, *in.Name)
	}
	if in.Phone != nil {
		sets = append(sets, "phone = ?")
		args = append(args, *in.Phone)
	}
	if in.Status != nil {
		sets = append(sets, "status = ?")
		args = append(args, *in.Status)
	}
	if in.TemplateType != nil {
		sets = append(sets, "template_type = ?")
		args = append(args, *in.TemplateType)
	}
	if in.CustomTemplate != nil {
		sets = append(sets, "custom_template = ?")
		args = append(args, *in.CustomTemplate)
	}
	if in.AccessMode != nil {
		sets = append(sets, "access_mode = ?")
		args = append(args, *in.AccessMode)
	}
	if in.PasscodeHash != nil {
		sets = append(sets, "passcode_hash = NULLIF(?, '')")
		args = append(args, *in.PasscodeHash)
	}
	if in.PasscodeHint != nil {
		sets = append(sets, "passcode_hint = NULLIF(?, '')")
		args = append(args, *in.PasscodeHint)
	}
	if in.TicketTTLSeconds != nil {
		sets = append(sets, "ticket_ttl_seconds = ?")
		args = append(args, *in.TicketTTLSeconds)
	}
	if in.AutoCloseSeconds != nil {
		sets = append(sets, "auto_close_seconds = ?")
		args = append(args, *in.AutoCloseSeconds)
	}
	if in.WatermarkEnabled != nil {
		v := 0
		if *in.WatermarkEnabled {
			v = 1
		}
		sets = append(sets, "watermark_enabled = ?")
		args = append(args, v)
	}
	if in.TemplatePayload != nil {
		sets = append(sets, "template_payload = ?")
		args = append(args, in.TemplatePayload)
	}
	if len(sets) == 0 {
		return r.GetByID(ctx, id)
	}
	sets = append(sets, "updated_at = NOW()")
	args = append(args, id)

	if _, err := r.db.ExecContext(ctx,
		"UPDATE nfc_tag SET "+strings.Join(sets, ", ")+" WHERE id = ?", args...); err != nil {
		return nil, err
	}
	return r.GetByID(ctx, id)
}

// Delete 删除标签（关联的 note/photo 依赖表级外键或应用层清理）。
func (r *Repo) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM nfc_tag WHERE id = ?", id)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

// UpdateStatusBatch 批量启停，返回受影响的 nfc_id 列表（用于失效缓存）。
func (r *Repo) UpdateStatusBatch(ctx context.Context, ids []int64, status string) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")
	args := make([]any, 0, len(ids)+1)
	args = append(args, status)
	for _, id := range ids {
		args = append(args, id)
	}

	rows, err := r.db.QueryContext(ctx,
		"SELECT nfc_id FROM nfc_tag WHERE id IN ("+placeholders+")", int64SliceToAny(ids)...)
	if err != nil {
		return nil, err
	}
	var nfcIDs []string
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			rows.Close()
			return nil, err
		}
		nfcIDs = append(nfcIDs, s)
	}
	rows.Close()

	if _, err := r.db.ExecContext(ctx,
		"UPDATE nfc_tag SET status = ?, updated_at = NOW() WHERE id IN ("+placeholders+")", args...); err != nil {
		return nil, err
	}
	return nfcIDs, nil
}

// DeleteBatch 批量删除，返回被删除的 nfc_id 列表。
func (r *Repo) DeleteBatch(ctx context.Context, ids []int64) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")
	args := int64SliceToAny(ids)

	rows, err := r.db.QueryContext(ctx,
		"SELECT nfc_id FROM nfc_tag WHERE id IN ("+placeholders+")", args...)
	if err != nil {
		return nil, err
	}
	var nfcIDs []string
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			rows.Close()
			return nil, err
		}
		nfcIDs = append(nfcIDs, s)
	}
	rows.Close()

	if _, err := r.db.ExecContext(ctx,
		"DELETE FROM nfc_tag WHERE id IN ("+placeholders+")", args...); err != nil {
		return nil, err
	}
	return nfcIDs, nil
}

// Creators 返回创建人筛选项，结构与 FastAPI 完全一致：
// 仅有标签的普通用户，外加一条「管理员创建」（id=0，带 count）。
func (r *Repo) Creators(ctx context.Context) ([]CreatorItem, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT u.id, u.username, COUNT(t.id) AS cnt
		FROM `+"`user`"+` u
		JOIN nfc_tag t ON t.user_id = u.id
		GROUP BY u.id, u.username
		ORDER BY u.username`)
	if err != nil {
		return nil, err
	}

	items := make([]CreatorItem, 0, 8)
	for rows.Next() {
		var it CreatorItem
		if err := rows.Scan(&it.ID, &it.Name, &it.Count); err != nil {
			rows.Close()
			return nil, err
		}
		it.Type = "user"
		items = append(items, it)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}

	var adminCount int
	if err := r.db.QueryRowContext(ctx,
		"SELECT COUNT(*) FROM nfc_tag WHERE user_id IS NULL").Scan(&adminCount); err != nil {
		return nil, err
	}
	if adminCount > 0 {
		items = append([]CreatorItem{{
			ID: 0, Name: "管理员创建", Type: "admin", Count: adminCount,
		}}, items...)
	}
	return items, nil
}

// CreatorItem 是创建人筛选项；Count 仅管理员那条有值（对齐 FastAPI 的返回）。
type CreatorItem struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count,omitempty"`
}

func int64SliceToAny(ids []int64) []any {
	out := make([]any, len(ids))
	for i, v := range ids {
		out[i] = v
	}
	return out
}
