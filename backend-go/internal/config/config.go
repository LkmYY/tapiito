package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config 是服务全量配置，全部来自环境变量，缺省值只用于本地开发。
type Config struct {
	Env        string // dev / prod
	PublicAddr string // 公开 API 监听地址，例如 :8080
	AdminAddr  string // 管理 API 监听地址，例如 :8081

	DB      DBConfig
	Cache   RedisConfig // 允许 allkeys-lru，丢失可重建
	Control RedisConfig // noeviction + AOF，保存 Bloom / 限流 / 黑名单

	JWT   JWTConfig
	Limit LimitConfig
	CDN   CacheConfig
	Bloom BloomConfig

	UploadDir      string
	// 访问明细保留天数：超过后由 worker 分批删除（聚合表不受影响）
	StatsRetentionDays int
	Upload         UploadConfig
	OSS            OSSConfig
	Twilio         TwilioConfig
	PublicBaseURL  string
	CORSOrigins    []string
	TrustedProxy   bool
	DefaultAdmin   string
	DefaultAdminPw string
}

// UploadConfig 限制上传大小与类型，避免被大文件打穿。
type UploadConfig struct {
	MaxSize      int64
	AllowedTypes map[string]struct{}
}

// OSSConfig 阿里云对象存储（OSS_ENABLE=0 时回退本地存储）。
type OSSConfig struct {
	Enable    bool
	Endpoint  string
	KeyID     string
	KeySecret string
	Bucket    string
}

// TwilioConfig 短信通知（未配置时提交不发送，返回 sent=false）。
type TwilioConfig struct {
	AccountSID string
	AuthToken  string
	From       string
}

type DBConfig struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ReadConcurrency int
	WriteConcurrency int
	QueueLimit      int
	QueueTimeout    time.Duration
	QueryTimeout    time.Duration
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

type JWTConfig struct {
	Secret   string
	Expire   time.Duration
	Issuer   string
}

type LimitConfig struct {
	GlobalQPS        int
	PublicReadRPS    int
	PublicReadBurst  int
	PublicWriteRPS   int
	PublicWriteBurst int
	AdminRPS         int
	AdminBurst       int
	LoginPerMinute   int
	BlacklistWindow  int // 统计窗口（秒）
	BlacklistHits    int // 窗口内 429 次数阈值
	BlacklistTTL     time.Duration
}

type CacheConfig struct {
	PageTTL    time.Duration
	PageJitter time.Duration
	MissTTL    time.Duration
}

type BloomConfig struct {
	KeyV1      string
	KeyV2      string
	CurrentKey string
	BuildingKey string
	DeltaKey   string
	ErrorRate  float64
	Capacity   int
	DeltaTTL   time.Duration
}

func Load() (*Config, error) {
	cfg := &Config{
		Env:        getEnv("APP_ENV", "dev"),
		PublicAddr: getEnv("PUBLIC_ADDR", ":8080"),
		AdminAddr:  getEnv("ADMIN_ADDR", ":8081"),

		DB: DBConfig{
			DSN:              getEnv("DB_DSN", "nfc_user:nfc_pass@tcp(127.0.0.1:3306)/nfc?charset=utf8mb4&parseTime=true&loc=Local"),
			MaxOpenConns:     getEnvInt("DB_MAX_OPEN_CONNS", 100),
			MaxIdleConns:     getEnvInt("DB_MAX_IDLE_CONNS", 20),
			ConnMaxLifetime:  getEnvDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
			ReadConcurrency:  getEnvInt("DB_READ_CONCURRENCY", 80),
			WriteConcurrency: getEnvInt("DB_WRITE_CONCURRENCY", 20),
			QueueLimit:       getEnvInt("DB_QUEUE_LIMIT", 200),
			QueueTimeout:     getEnvDuration("DB_QUEUE_TIMEOUT", 200*time.Millisecond),
			QueryTimeout:     getEnvDuration("DB_QUERY_TIMEOUT", 3*time.Second),
		},
		Cache: RedisConfig{
			Addr:     getEnv("REDIS_CACHE_ADDR", "127.0.0.1:6379"),
			Password: getEnv("REDIS_CACHE_PASSWORD", ""),
			DB:       getEnvInt("REDIS_CACHE_DB", 0),
			PoolSize: getEnvInt("REDIS_CACHE_POOL", 200),
		},
		Control: RedisConfig{
			Addr:     getEnv("REDIS_CONTROL_ADDR", "127.0.0.1:6380"),
			Password: getEnv("REDIS_CONTROL_PASSWORD", ""),
			DB:       getEnvInt("REDIS_CONTROL_DB", 0),
			PoolSize: getEnvInt("REDIS_CONTROL_POOL", 100),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", ""),
			Expire: getEnvDuration("JWT_ACCESS_TOKEN_EXPIRE_MINUTES", 60*time.Minute),
			Issuer: getEnv("JWT_ISSUER", "tapiito"),
		},
		Limit: LimitConfig{
			GlobalQPS:        getEnvInt("RL_GLOBAL_QPS", 400),
			PublicReadRPS:    getEnvInt("RL_PUBLIC_READ_RPS", 10),
			PublicReadBurst:  getEnvInt("RL_PUBLIC_READ_BURST", 20),
			PublicWriteRPS:   getEnvInt("RL_PUBLIC_WRITE_RPS", 2),
			PublicWriteBurst: getEnvInt("RL_PUBLIC_WRITE_BURST", 5),
			AdminRPS:         getEnvInt("RL_ADMIN_RPS", 20),
			AdminBurst:       getEnvInt("RL_ADMIN_BURST", 40),
			LoginPerMinute:   getEnvInt("RL_LOGIN_PER_MINUTE", 5),
			BlacklistWindow:  getEnvInt("BL_WINDOW_SECONDS", 60),
			BlacklistHits:    getEnvInt("BL_HITS", 20),
			BlacklistTTL:     getEnvDuration("BL_TTL_MINUTES", 10*time.Minute),
		},
		CDN: CacheConfig{
			PageTTL:    getEnvDuration("CACHE_PAGE_TTL_MINUTES", 10*time.Minute),
			PageJitter: getEnvDuration("CACHE_PAGE_JITTER_SECONDS", 60*time.Second),
			MissTTL:    getEnvDuration("CACHE_MISS_TTL_SECONDS", 30*time.Second),
		},
		Bloom: BloomConfig{
			KeyV1:       getEnv("BLOOM_KEY_V1", "bloom:v1"),
			KeyV2:       getEnv("BLOOM_KEY_V2", "bloom:v2"),
			CurrentKey:  getEnv("BLOOM_CURRENT_KEY", "bloom:current"),
			BuildingKey: getEnv("BLOOM_BUILDING_KEY", "bloom:building"),
			DeltaKey:    getEnv("BLOOM_DELTA_KEY", "bloom:delta"),
			ErrorRate:   getEnvFloat("BLOOM_ERROR_RATE", 0.001),
			Capacity:    getEnvInt("BLOOM_CAPACITY", 1_000_000),
			DeltaTTL:    getEnvDuration("BLOOM_DELTA_TTL_HOURS", 24*time.Hour),
		},
		UploadDir:          getEnv("UPLOAD_DIR", "./data/uploads"),
		StatsRetentionDays: getEnvInt("STATS_RETENTION_DAYS", 30),
		Upload: UploadConfig{
			MaxSize:      int64(getEnvInt("MAX_UPLOAD_SIZE_MB", 8)) << 20,
			AllowedTypes: parseTypes(getEnv("ALLOWED_UPLOAD_TYPES", "image/jpeg,image/png,image/webp,image/gif")),
		},
		OSS: OSSConfig{
			Enable:    getEnvBool("OSS_ENABLE", false),
			Endpoint:  getEnv("OSS_ENDPOINT", ""),
			KeyID:     getEnv("OSS_ACCESS_KEY_ID", ""),
			KeySecret: getEnv("OSS_ACCESS_KEY_SECRET", ""),
			Bucket:    getEnv("OSS_BUCKET_NAME", ""),
		},
		Twilio: TwilioConfig{
			AccountSID: getEnv("TWILIO_ACCOUNT_SID", ""),
			AuthToken:  getEnv("TWILIO_AUTH_TOKEN", ""),
			From:       getEnv("TWILIO_PHONE_NUMBER", ""),
		},
		PublicBaseURL: getEnv("PUBLIC_BASE_URL", ""),
		CORSOrigins:    splitCSV(getEnv("CORS_ORIGINS", "*")),
		TrustedProxy:   getEnvBool("TRUST_PROXY", true),
		DefaultAdmin:   getEnv("DEFAULT_ADMIN_USERNAME", ""),
		DefaultAdminPw: getEnv("DEFAULT_ADMIN_PASSWORD", ""),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) IsProd() bool { return strings.EqualFold(c.Env, "prod") }

// validate：生产环境下禁止使用弱配置，避免默认密码/密钥上线。
func (c *Config) validate() error {
	if c.DB.ReadConcurrency <= 0 || c.DB.WriteConcurrency <= 0 {
		return fmt.Errorf("DB_READ_CONCURRENCY / DB_WRITE_CONCURRENCY 必须大于 0")
	}
	if c.Limit.GlobalQPS <= 0 {
		return fmt.Errorf("RL_GLOBAL_QPS 必须大于 0")
	}
	if c.JWT.Secret == "" {
		if c.IsProd() {
			return fmt.Errorf("生产环境必须设置 JWT_SECRET")
		}
		c.JWT.Secret = "dev-only-secret-do-not-use-in-prod"
	}
	if c.IsProd() {
		if len(c.JWT.Secret) < 32 {
			return fmt.Errorf("生产环境 JWT_SECRET 长度不得小于 32")
		}
		if c.JWT.Secret == "dev-only-secret-do-not-use-in-prod" {
			return fmt.Errorf("生产环境不得使用开发默认 JWT_SECRET")
		}
		if c.DefaultAdminPw == "" {
			return fmt.Errorf("生产环境必须显式设置 DEFAULT_ADMIN_PASSWORD（或预先建好管理员）")
		}
	}
	return nil
}

func getEnv(key, def string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return def
}

func getEnvInt(key string, def int) int {
	if v, err := strconv.Atoi(getEnv(key, "")); err == nil {
		return v
	}
	return def
}

func getEnvFloat(key string, def float64) float64 {
	if v, err := strconv.ParseFloat(getEnv(key, ""), 64); err == nil {
		return v
	}
	return def
}

func getEnvBool(key string, def bool) bool {
	if v, err := strconv.ParseBool(getEnv(key, "")); err == nil {
		return v
	}
	return def
}

// getEnvDuration 支持纯数字（按给定默认单位）或 Go duration 字符串。
func getEnvDuration(key string, def time.Duration) time.Duration {
	raw := getEnv(key, "")
	if raw == "" {
		return def
	}
	if n, err := strconv.Atoi(raw); err == nil {
		// 纯数字沿用默认值的单位
		switch def {
		case 60 * time.Minute, 10 * time.Minute, 5 * time.Minute:
			return time.Duration(n) * time.Minute
		case 24 * time.Hour:
			return time.Duration(n) * time.Hour
		default:
			return time.Duration(n) * time.Second
		}
	}
	if d, err := time.ParseDuration(raw); err == nil {
		return d
	}
	return def
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if v := strings.TrimSpace(p); v != "" {
			out = append(out, v)
		}
	}
	return out
}

func parseTypes(csv string) map[string]struct{} {
	out := make(map[string]struct{}, 8)
	for _, v := range strings.Split(csv, ",") {
		if t := strings.TrimSpace(v); t != "" {
			out[t] = struct{}{}
		}
	}
	return out
}
