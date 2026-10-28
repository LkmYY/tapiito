package tag

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

// NoteResp 与 FastAPI 的 NoteResp 对齐。
type NoteResp struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

// PhotoResp 与 FastAPI 的 PhotoResp 对齐。
type PhotoResp struct {
	ID       int64   `json:"id"`
	OSSKey   string  `json:"oss_key"`
	MimeType *string `json:"mime_type"`
	Size     *int64  `json:"size"`
}

// enabledTagID 取启用中标签的主键；不存在或停用一律按 404 处理（公开写入不得落到停用标签）。
func (r *Repo) enabledTagID(ctx context.Context, nfcID string) (int64, error) {
	var id int64
	err := r.db.QueryRowContext(ctx,
		"SELECT id FROM nfc_tag WHERE nfc_id = ? AND `status` = 'enabled' LIMIT 1", nfcID).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, ErrNotFound
	}
	return id, err
}

// AddNote 给标签添加备注。
func (s *Service) AddNote(ctx context.Context, nfcID, content, ip string) (*NoteResp, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, fmt.Errorf("content required")
	}
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	tagID, err := s.repo.enabledTagID(qctx, nfcID)
	if err != nil {
		return nil, err
	}
	res, err := s.repo.db.ExecContext(qctx,
		"INSERT INTO note (nfc_tag_id, content, created_ip, created_at) VALUES (?, ?, ?, NOW())",
		tagID, content, ip)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &NoteResp{ID: id, Content: content}, nil
}

// AddPhoto 记录已保存的照片。
func (s *Service) AddPhoto(ctx context.Context, nfcID, ossKey, mimeType string, size int64, ip string) (*PhotoResp, error) {
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	tagID, err := s.repo.enabledTagID(qctx, nfcID)
	if err != nil {
		return nil, err
	}
	res, err := s.repo.db.ExecContext(qctx,
		"INSERT INTO photo (nfc_tag_id, oss_key, mime_type, `size`, created_ip, created_at) VALUES (?, ?, ?, ?, ?, NOW())",
		tagID, ossKey, mimeType, size, ip)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	mt, sz := mimeType, size
	return &PhotoResp{ID: id, OSSKey: ossKey, MimeType: &mt, Size: &sz}, nil
}

// ListPhotos 列出标签的照片。
func (s *Service) ListPhotos(ctx context.Context, nfcID string) ([]PhotoResp, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	tagID, err := s.repo.enabledTagID(qctx, nfcID)
	if err != nil {
		return nil, err
	}
	rows, err := s.repo.db.QueryContext(qctx,
		"SELECT id, oss_key, mime_type, `size` FROM photo WHERE nfc_tag_id = ? ORDER BY id DESC", tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]PhotoResp, 0, 8)
	for rows.Next() {
		var p PhotoResp
		if err := rows.Scan(&p.ID, &p.OSSKey, &p.MimeType, &p.Size); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// SubmitResult 与 FastAPI 的 SubmitResp 对齐。
type SubmitResult struct {
	Sent bool   `json:"sent"`
	Mode string `json:"mode"`
}

// Submit 写入备注并返回短信发送结果。
//
// 短信由外部注入的 sender 执行（失败不影响备注落库），与 FastAPI 行为一致：
// 未配置手机号 → 400；未配置 Twilio → sent=false 但仍记录备注。
func (s *Service) Submit(
	ctx context.Context,
	nfcID, content string,
	photoIDs []int64,
	send func(ctx context.Context, to, body string, mediaURLs []string) bool,
	baseURL string,
) (*SubmitResult, error) {
	content = strings.TrimSpace(content)
	if content == "" {
		return nil, fmt.Errorf("content required")
	}
	release, err := s.writeGate.Acquire(ctx)
	if err != nil {
		return nil, ErrUnavailable
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var tagID int64
	var phone *string
	err = s.repo.db.QueryRowContext(qctx,
		"SELECT id, phone FROM nfc_tag WHERE nfc_id = ? AND `status` = 'enabled' LIMIT 1", nfcID).
		Scan(&tagID, &phone)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	if phone == nil || strings.TrimSpace(*phone) == "" {
		return nil, ErrPhoneNotConfigured
	}

	if _, err := s.repo.db.ExecContext(qctx,
		"INSERT INTO note (nfc_tag_id, content, created_ip, created_at) VALUES (?, ?, '', NOW())",
		tagID, content); err != nil {
		return nil, err
	}

	// 取照片的访问地址（用于彩信）
	mediaURLs := make([]string, 0, len(photoIDs))
	if len(photoIDs) > 0 {
		placeholders := strings.TrimSuffix(strings.Repeat("?,", len(photoIDs)), ",")
		args := append([]any{tagID}, int64SliceToAny(photoIDs)...)
		rows, err := s.repo.db.QueryContext(qctx,
			"SELECT oss_key FROM photo WHERE nfc_tag_id = ? AND id IN ("+placeholders+") ORDER BY id ASC", args...)
		if err == nil {
			for rows.Next() {
				var key string
				if err := rows.Scan(&key); err == nil && key != "" {
					if strings.HasPrefix(key, "http") {
						mediaURLs = append(mediaURLs, key)
					} else {
						mediaURLs = append(mediaURLs, strings.TrimRight(baseURL, "/")+"/uploads/"+key)
					}
				}
			}
			rows.Close()
		}
	}

	sent := false
	if send != nil {
		sent = send(context.WithoutCancel(ctx), *phone, "备注："+content, mediaURLs)
	}
	mode := "sms"
	if len(mediaURLs) > 0 {
		mode = "mms"
	}
	return &SubmitResult{Sent: sent, Mode: mode}, nil
}

// ErrPhoneNotConfigured 表示标签未绑定手机号，无法发送通知。
var ErrPhoneNotConfigured = errors.New("phone not configured")
