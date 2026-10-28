package httpapi

import (
	"errors"
	"log/slog"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/lkmyy/tapiito-go/internal/auth"
	"github.com/lkmyy/tapiito-go/internal/ratelimit"
)

const (
	ctxClientIP = "client_ip"
	ctxRequestID = "request_id"
	ctxUser      = "auth_user"
)

// RequestID 为每个请求生成可追踪的 ID。
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetHeader("X-Request-ID")
		if id == "" {
			id = strconv.FormatInt(time.Now().UnixNano(), 36)
		}
		c.Set(ctxRequestID, id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}

// RealIP 只信任 Nginx 注入的 X-Real-IP，杜绝客户端伪造 X-Forwarded-For。
func RealIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := strings.TrimSpace(c.GetHeader("X-Real-IP"))
		if ip == "" {
			host, _, err := net.SplitHostPort(c.Request.RemoteAddr)
			if err != nil {
				host = c.Request.RemoteAddr
			}
			ip = host
		}
		c.Set(ctxClientIP, ip)
		c.Next()
	}
}

// SecurityHeaders 输出最小安全响应头。
func SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "no-referrer")
		c.Next()
	}
}

// BlacklistGuard 命中临时黑名单返回 403。
func BlacklistGuard(bl *ratelimit.Blacklist, log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := ClientIP(c)
		blocked, err := bl.Blocked(c.Request.Context(), ip)
		if err != nil {
			log.Warn("blacklist check failed", "ip", ip, "err", err)
		}
		if blocked {
			c.AbortWithStatusJSON(403, gin.H{"detail": "forbidden"})
			return
		}
		c.Next()
	}
}

// RateLimit 按 key 维度限流；超限记一次违规，累计到阈值即封禁。
func RateLimit(l *ratelimit.Limiter, bl *ratelimit.Blacklist, keyPrefix string, rate, burst int, log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := ClientIP(c)
		key := keyPrefix + ":" + ip
		allowed, err := l.Allow(c.Request.Context(), key, rate, burst)
		if err != nil {
			log.Warn("ratelimit error", "key", key, "err", err)
		}
		if !allowed {
			if blocked, _ := bl.RecordViolation(c.Request.Context(), ip); blocked {
				log.Warn("ip blacklisted", "ip", ip)
			}
			c.Header("Retry-After", "1")
			c.AbortWithStatusJSON(429, gin.H{"detail": "too many requests"})
			return
		}
		c.Next()
	}
}

// GlobalQPS 限制整个公开入口的总吞吐。
func GlobalQPS(l *ratelimit.Limiter, qps int, log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		allowed, err := l.Allow(c.Request.Context(), "rl:global", qps, qps*2)
		if err != nil {
			log.Warn("global ratelimit error", "err", err)
		}
		if !allowed {
			c.Header("Retry-After", "1")
			c.AbortWithStatusJSON(429, gin.H{"detail": "service busy"})
			return
		}
		c.Next()
	}
}

// RequireAdmin 校验 JWT，仅管理入口使用。
func RequireAdmin(iss *auth.Issuer, log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.GetHeader("Authorization")
		if !strings.HasPrefix(strings.ToLower(raw), "bearer ") {
			unauthorized(c)
			return
		}
		token := strings.TrimSpace(raw[7:])
		claims, err := iss.Parse(token)
		if err != nil {
			if errors.Is(err, auth.ErrExpired) {
				c.Header("X-Auth-Error", "expired")
			}
			log.Debug("jwt rejected", "err", err)
			unauthorized(c)
			return
		}
		c.Set(ctxUser, claims)
		c.Next()
	}
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", "Bearer")
	c.AbortWithStatusJSON(401, gin.H{"detail": "Not authenticated"})
}

// ClientIP 读取中间件写入的客户端 IP。
func ClientIP(c *gin.Context) string {
	if v, ok := c.Get(ctxClientIP); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// RequestIDOf 读取请求 ID。
func RequestIDOf(c *gin.Context) string {
	if v, ok := c.Get(ctxRequestID); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
