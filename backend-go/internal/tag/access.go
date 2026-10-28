// Package tag 的访问控制：暗号校验、一次性 ticket 签发与验证、防爆破。
package tag

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 访问控制模式
const (
	AccessNone     = "none"     // 直出，零开销
	AccessPasscode = "passcode" // 仅暗号
	AccessTicket   = "ticket"   // 仅一次性 ticket（扫码即发）
	AccessBoth     = "both"     // 暗号换 ticket
)

var (
	ErrPasscodeRequired = errors.New("passcode required")
	ErrTicketRequired   = errors.New("ticket required")
	ErrPasscodeWrong    = errors.New("passcode wrong")
	ErrTooManyAttempts  = errors.New("too many attempts")
	ErrTicketInvalid    = errors.New("ticket invalid")
	ErrPasscodeMissing  = errors.New("passcode not configured")
)

// AccessConfig 是标签的访问控制配置。
type AccessConfig struct {
	Mode            string
	HasPasscode     bool
	PasscodeHint    string
	TicketTTL       time.Duration
	AutoCloseSecond int
	Watermark       bool
}

// 防爆破：同一维度连续失败 3 次锁 30 分钟
const (
	maxAttempts   = 3
	lockDuration  = 30 * time.Minute
	attemptWindow = 30 * time.Minute
)

func (c AccessConfig) IsZero() bool { return c.Mode == "" || c.Mode == AccessNone }

func (c AccessConfig) NeedPasscode() bool {
	return c.Mode == AccessPasscode || c.Mode == AccessBoth
}

func (c AccessConfig) NeedTicket() bool {
	return c.Mode == AccessTicket || c.Mode == AccessBoth
}

// accessConfig 从数据库读取访问控制配置。
func (r *Repo) accessConfig(ctx context.Context, tagID int64) (AccessConfig, error) {
	var (
		mode     string
		hash     *string
		hint     *string
		ttl      int
		autoClos int
		mark     int
	)
	err := r.db.QueryRowContext(ctx,
		"SELECT access_mode, passcode_hash, passcode_hint, ticket_ttl_seconds, auto_close_seconds, watermark_enabled "+
			"FROM nfc_tag WHERE id = ?", tagID).
		Scan(&mode, &hash, &hint, &ttl, &autoClos, &mark)
	if err != nil {
		return AccessConfig{}, err
	}
	cfg := AccessConfig{
		Mode:            mode,
		HasPasscode:     hash != nil && *hash != "",
		TicketTTL:       time.Duration(ttl) * time.Second,
		AutoCloseSecond: autoClos,
		Watermark:       mark == 1,
	}
	if hint != nil {
		cfg.PasscodeHint = *hint
	}
	if cfg.TicketTTL <= 0 {
		cfg.TicketTTL = time.Hour
	}
	return cfg, nil
}

// IssueTicket 签发一次性 ticket（记录来源 IP，暂不绑定会话）。
func (r *Repo) IssueTicket(ctx context.Context, tagID int64, ttl time.Duration, ip string) (string, time.Time, error) {
	raw := make([]byte, 32)
	if _, err := rand.Read(raw); err != nil {
		return "", time.Time{}, fmt.Errorf("rand: %w", err)
	}
	ticket := hex.EncodeToString(raw)
	expires := time.Now().Add(ttl)
	_, err := r.db.ExecContext(ctx,
		"INSERT INTO access_ticket (nfc_tag_id, ticket, issued_at, expires_at, issued_ip) VALUES (?, ?, NOW(), ?, ?)",
		tagID, ticket, expires, ip)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("insert ticket: %w", err)
	}
	return ticket, expires, nil
}

// VerifyTicket 校验 ticket：存在、未过期、且未被其它会话使用。
//
// 绑定会话策略：首次使用时写入 sessionID；后续同一 sessionID 可重复访问（刷新不白屏），
// 其它 sessionID 一律拒绝 —— 这样转发出去的链接在对方浏览器必然失效。
func (r *Repo) VerifyTicket(ctx context.Context, tagID int64, ticket, sessionID string) error {
	if strings.TrimSpace(ticket) == "" {
		return ErrTicketInvalid
	}
	var (
		dbSession *string
		expires   time.Time
	)
	err := r.db.QueryRowContext(ctx,
		"SELECT session_id, expires_at FROM access_ticket WHERE ticket = ? AND nfc_tag_id = ?",
		ticket, tagID).Scan(&dbSession, &expires)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrTicketInvalid
	}
	if err != nil {
		return err
	}
	if time.Now().After(expires) {
		return ErrTicketInvalid
	}
	if dbSession == nil || *dbSession == "" {
		// 首次使用：绑定当前会话
		if sessionID == "" {
			return ErrTicketInvalid
		}
		if _, err := r.db.ExecContext(ctx,
			"UPDATE access_ticket SET session_id = ?, used_at = NOW() WHERE ticket = ? AND (session_id IS NULL OR session_id = '')",
			sessionID, ticket); err != nil {
			return err
		}
		return nil
	}
	if *dbSession != sessionID {
		// 已被别的会话用过 —— 转发场景
		return ErrTicketInvalid
	}
	return nil
}

// CheckPasscode 校验暗号，并按 IP + 标签双维度累计失败次数。
//
// 计数放 Redis 由上层（限流器）负责，这里只做 bcrypt 比对。
func (r *Repo) CheckPasscode(ctx context.Context, tagID int64, passcode string) error {
	var hash *string
	err := r.db.QueryRowContext(ctx, "SELECT passcode_hash FROM nfc_tag WHERE id = ?", tagID).Scan(&hash)
	if err != nil {
		return err
	}
	if hash == nil || *hash == "" {
		return ErrPasscodeMissing
	}
	if bcrypt.CompareHashAndPassword([]byte(*hash), []byte(passcode)) != nil {
		return ErrPasscodeWrong
	}
	return nil
}

// HashPasscode 生成 bcrypt 哈希（cost 10，单次比对约 50-80ms，天然抗爆破）。
func HashPasscode(plain string) (string, error) {
	if len(strings.TrimSpace(plain)) < 2 {
		return "", errors.New("passcode too short")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), 10)
	if err != nil {
		return "", err
	}
	return string(h), nil
}

// CleanupExpiredTickets 删除过期 ticket（由 worker 定期调用）。
func (r *Repo) CleanupExpiredTickets(ctx context.Context) (int64, error) {
	res, err := r.db.ExecContext(ctx,
		"DELETE FROM access_ticket WHERE expires_at < NOW() - INTERVAL 1 DAY LIMIT 5000")
	if err != nil {
		return 0, err
	}
	n, _ := res.RowsAffected()
	return n, nil
}

// AccessInfo 按 nfc_id 读取访问控制配置（走 nfc_id 唯一索引，单行查询）。
//
// 只有非 none 模式才会在请示链路里被调用，none 标签保持零额外开销。
func (s *Service) AccessInfo(ctx context.Context, nfcID string) (AccessConfig, int64, error) {
	var (
		tagID    int64
		mode     string
		hash     *string
		hint     *string
		ttl      int
		autoClos int
		mark     int
	)
	err := s.repo.db.QueryRowContext(ctx,
		"SELECT id, access_mode, passcode_hash, passcode_hint, ticket_ttl_seconds, auto_close_seconds, watermark_enabled "+
			"FROM nfc_tag WHERE nfc_id = ?", nfcID).
		Scan(&tagID, &mode, &hash, &hint, &ttl, &autoClos, &mark)
	if errors.Is(err, sql.ErrNoRows) {
		return AccessConfig{}, 0, ErrNotFound
	}
	if err != nil {
		return AccessConfig{}, 0, err
	}
	cfg := AccessConfig{
		Mode:            mode,
		HasPasscode:     hash != nil && *hash != "",
		TicketTTL:       time.Duration(ttl) * time.Second,
		AutoCloseSecond: autoClos,
		Watermark:       mark == 1,
	}
	if hint != nil {
		cfg.PasscodeHint = *hint
	}
	if cfg.TicketTTL <= 0 {
		cfg.TicketTTL = time.Hour
	}
	return cfg, tagID, nil
}

// Unlock 校验暗号并签发 ticket（both / passcode 模式）。
func (s *Service) Unlock(ctx context.Context, nfcID, passcode, ip string) (string, time.Time, AccessConfig, error) {
	cfg, tagID, err := s.AccessInfo(ctx, nfcID)
	if err != nil {
		return "", time.Time{}, cfg, err
	}
	if !cfg.HasPasscode {
		return "", time.Time{}, cfg, ErrPasscodeMissing
	}
	if err := s.repo.CheckPasscode(ctx, tagID, passcode); err != nil {
		return "", time.Time{}, cfg, err
	}
	ticket, expires, err := s.repo.IssueTicket(ctx, tagID, cfg.TicketTTL, ip)
	return ticket, expires, cfg, err
}

// IssueOnScan 求助类：扫码即发 ticket，不阻塞访问。
func (s *Service) IssueOnScan(ctx context.Context, nfcID, ip string) (string, time.Time, AccessConfig, error) {
	cfg, tagID, err := s.AccessInfo(ctx, nfcID)
	if err != nil {
		return "", time.Time{}, cfg, err
	}
	ticket, expires, err := s.repo.IssueTicket(ctx, tagID, cfg.TicketTTL, ip)
	return ticket, expires, cfg, err
}

// VerifyAccess 校验访问凭证，返回是否需要拦截。
func (s *Service) VerifyAccess(ctx context.Context, nfcID, ticket, sessionID string) (AccessConfig, error) {
	cfg, tagID, err := s.AccessInfo(ctx, nfcID)
	if err != nil {
		return cfg, err
	}
	if cfg.Mode == AccessNone || cfg.Mode == "" {
		return cfg, nil
	}
	if cfg.NeedTicket() {
		if err := s.repo.VerifyTicket(ctx, tagID, ticket, sessionID); err != nil {
			if cfg.NeedPasscode() {
				return cfg, ErrPasscodeRequired
			}
			return cfg, ErrTicketRequired
		}
		return cfg, nil
	}
	// 仅 passcode 模式：没有 ticket 就必须先输暗号
	if cfg.NeedPasscode() {
		return cfg, ErrPasscodeRequired
	}
	return cfg, nil
}
