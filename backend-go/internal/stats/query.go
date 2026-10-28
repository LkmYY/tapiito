// Package stats 提供访问统计查询（管理端）。
//
// 优先读 visit_stats_daily 聚合表（10 万级标签下明细表会非常大）；
// 聚合表为空（历史数据未回填）时回退到明细表，保证与 FastAPI 结果一致。
package stats

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/lkmyy/tapiito-go/internal/dbgate"
)

var ErrTagNotFound = fmt.Errorf("tag not found")

type Service struct {
	db           *sql.DB
	readGate     *dbgate.Gate
	queryTimeout time.Duration
	log          *slog.Logger
}

type Deps struct {
	DB           *sql.DB
	ReadGate     *dbgate.Gate
	QueryTimeout time.Duration
	Logger       *slog.Logger
}

func NewService(d Deps) *Service {
	if d.QueryTimeout <= 0 {
		d.QueryTimeout = 5 * time.Second
	}
	return &Service{db: d.DB, readGate: d.ReadGate, queryTimeout: d.QueryTimeout, log: d.Logger}
}

// Overview 是统计概览，字段与 FastAPI 的 /api/stats/overview 对齐。
type Overview struct {
	TotalVisits int64          `json:"total_visits"`
	TopTags     []TopTag       `json:"top_tags"`
	DailyStats  []DailyStat    `json:"daily_stats"`
}

type TopTag struct {
	NfcID      string `json:"nfc_id"`
	VisitCount int64  `json:"visit_count"`
}

type DailyStat struct {
	Date       string `json:"date"`
	VisitCount int64  `json:"visit_count"`
}

// Overview 查询概览：total / 热门标签 Top10 / 按日趋势。
func (s *Service) Overview(ctx context.Context, startDate, endDate string) (*Overview, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	useDaily, err := s.hasDailyData(qctx)
	if err != nil {
		return nil, err
	}

	out := &Overview{TopTags: []TopTag{}, DailyStats: []DailyStat{}}

	if useDaily {
		where, args := dailyRange(startDate, endDate)
		if err := s.db.QueryRowContext(qctx,
			"SELECT COALESCE(SUM(`count`),0) FROM visit_stats_daily"+where, args...).Scan(&out.TotalVisits); err != nil {
			return nil, err
		}
		rows, err := s.db.QueryContext(qctx, `
			SELECT t.nfc_id, SUM(d.`+"`count`"+`) AS c
			FROM visit_stats_daily d JOIN nfc_tag t ON t.id = d.nfc_tag_id`+where+`
			GROUP BY t.id, t.nfc_id ORDER BY c DESC LIMIT 10`, args...)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var tt TopTag
			if err := rows.Scan(&tt.NfcID, &tt.VisitCount); err != nil {
				rows.Close()
				return nil, err
			}
			out.TopTags = append(out.TopTags, tt)
		}
		rows.Close()

		drows, err := s.db.QueryContext(qctx, `
			SELECT DATE_FORMAT(stat_date, '%Y-%m-%d') AS d, SUM(`+"`count`"+`) AS c
			FROM visit_stats_daily`+where+` GROUP BY d ORDER BY d`, args...)
		if err != nil {
			return nil, err
		}
		for drows.Next() {
			var ds DailyStat
			if err := drows.Scan(&ds.Date, &ds.VisitCount); err != nil {
				drows.Close()
				return nil, err
			}
			out.DailyStats = append(out.DailyStats, ds)
		}
		drows.Close()
		return out, nil
	}

	// 回退：明细表
	where, args := detailRange(startDate, endDate)
	if err := s.db.QueryRowContext(qctx,
		"SELECT COUNT(*) FROM visit_stats WHERE "+where, args...).Scan(&out.TotalVisits); err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(qctx, `
		SELECT t.nfc_id, COUNT(v.id) AS c
		FROM visit_stats v JOIN nfc_tag t ON t.id = v.nfc_tag_id
		WHERE `+where+` GROUP BY t.id, t.nfc_id ORDER BY c DESC LIMIT 10`, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var tt TopTag
		if err := rows.Scan(&tt.NfcID, &tt.VisitCount); err != nil {
			rows.Close()
			return nil, err
		}
		out.TopTags = append(out.TopTags, tt)
	}
	rows.Close()

	drows, err := s.db.QueryContext(qctx, `
		SELECT DATE_FORMAT(created_at, '%Y-%m-%d') AS d, COUNT(*) AS c
		FROM visit_stats WHERE `+where+` GROUP BY d ORDER BY d`, args...)
	if err != nil {
		return nil, err
	}
	for drows.Next() {
		var ds DailyStat
		if err := drows.Scan(&ds.Date, &ds.VisitCount); err != nil {
			drows.Close()
			return nil, err
		}
		out.DailyStats = append(out.DailyStats, ds)
	}
	drows.Close()
	return out, nil
}

// Visit 是单条访问明细。
type Visit struct {
	ID        int64   `json:"id"`
	IPAddress *string `json:"ip_address"`
	UserAgent *string `json:"user_agent"`
	Referer   *string `json:"referer"`
	CreatedAt *string `json:"created_at"`
	Country   *string `json:"country"`
	Region    *string `json:"region"`
	City      *string `json:"city"`
}

// TagStats 是单标签统计，字段与 /api/stats/nfc/{id} 对齐。
type TagStats struct {
	NfcID       string         `json:"nfc_id"`
	TotalVisits int            `json:"total_visits"`
	UniqueIPs   int            `json:"unique_ips"`
	DeviceTypes map[string]int `json:"device_types"`
	Visits      []Visit        `json:"visits"`
}

// ByNfcID 查询单个标签的访问统计（明细 + 设备分布）。
func (s *Service) ByNfcID(ctx context.Context, nfcID, startDate, endDate string) (*TagStats, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	var tagID int64
	err = s.db.QueryRowContext(qctx, "SELECT id FROM nfc_tag WHERE nfc_id = ?", nfcID).Scan(&tagID)
	if err == sql.ErrNoRows {
		return nil, ErrTagNotFound
	}
	if err != nil {
		return nil, err
	}

	where, args := detailRange(startDate, endDate)
	q := "SELECT id, ip_address, user_agent, referer, created_at, country, region, city" +
		" FROM visit_stats WHERE nfc_tag_id = ? AND " + where + " ORDER BY created_at DESC LIMIT 5000"
	rows, err := s.db.QueryContext(qctx, q, append([]any{tagID}, args...)...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := &TagStats{NfcID: nfcID, DeviceTypes: map[string]int{}, Visits: []Visit{}}
	ips := map[string]struct{}{}
	for rows.Next() {
		v := Visit{}
		var created *time.Time
		if err := rows.Scan(&v.ID, &v.IPAddress, &v.UserAgent, &v.Referer, &created,
			&v.Country, &v.Region, &v.City); err != nil {
			return nil, err
		}
		if created != nil {
			ts := created.Format("2006-01-02T15:04:05")
			v.CreatedAt = &ts
		}
		out.Visits = append(out.Visits, v)
		if v.IPAddress != nil && *v.IPAddress != "" {
			ips[*v.IPAddress] = struct{}{}
		}
		ua := ""
		if v.UserAgent != nil {
			ua = *v.UserAgent
		}
		if strings.Contains(ua, "Mobile") || strings.Contains(ua, "Android") || strings.Contains(ua, "iPhone") {
			out.DeviceTypes["mobile"]++
		} else {
			out.DeviceTypes["desktop"]++
		}
	}
	out.TotalVisits = len(out.Visits)
	out.UniqueIPs = len(ips)
	return out, rows.Err()
}

// ExportRow 是导出的一行。
type ExportRow struct {
	NfcID     string  `json:"nfc_id"`
	IPAddress *string `json:"ip_address"`
	UserAgent *string `json:"user_agent"`
	Referer   *string `json:"referer"`
	Country   *string `json:"country"`
	Region    *string `json:"region"`
	City      *string `json:"city"`
	CreatedAt *string `json:"created_at"`
}

// Export 导出明细（最多 50000 行，避免一次拉爆内存）。
func (s *Service) Export(ctx context.Context, startDate, endDate string) ([]ExportRow, error) {
	release, err := s.readGate.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer release()
	qctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	where, args := detailRange(startDate, endDate)
	rows, err := s.db.QueryContext(qctx, `
		SELECT t.nfc_id, v.ip_address, v.user_agent, v.referer, v.country, v.region, v.city, v.created_at
		FROM visit_stats v JOIN nfc_tag t ON t.id = v.nfc_tag_id
		WHERE `+where+` ORDER BY v.id DESC LIMIT 50000`, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := make([]ExportRow, 0, 1024)
	for rows.Next() {
		var r ExportRow
		var created *time.Time
		if err := rows.Scan(&r.NfcID, &r.IPAddress, &r.UserAgent, &r.Referer,
			&r.Country, &r.Region, &r.City, &created); err != nil {
			return nil, err
		}
		if created != nil {
			ts := created.Format("2006-01-02T15:04:05")
			r.CreatedAt = &ts
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

func (s *Service) hasDailyData(ctx context.Context) (bool, error) {
	var n int
	if err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM visit_stats_daily LIMIT 1").Scan(&n); err != nil {
		return false, err
	}
	return n > 0, nil
}

// dailyRange 构造聚合表的日期条件。
func dailyRange(startDate, endDate string) (string, []any) {
	where := " WHERE 1=1"
	args := make([]any, 0, 2)
	if v := strings.TrimSpace(startDate); v != "" {
		where += " AND stat_date >= ?"
		args = append(args, v)
	}
	if v := strings.TrimSpace(endDate); v != "" {
		where += " AND stat_date <= ?"
		args = append(args, v)
	}
	return where, args
}

// detailRange 构造明细表的日期条件。
func detailRange(startDate, endDate string) (string, []any) {
	where := "1=1"
	args := make([]any, 0, 2)
	if v := strings.TrimSpace(startDate); v != "" {
		where += " AND created_at >= ?"
		args = append(args, v+" 00:00:00")
	}
	if v := strings.TrimSpace(endDate); v != "" {
		where += " AND created_at <= ?"
		args = append(args, v+" 23:59:59")
	}
	return where, args
}
