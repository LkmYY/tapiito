// Package auth 提供 JWT 签发/校验与口令哈希校验。
package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalid = errors.New("invalid token")
	ErrExpired = errors.New("token expired")
)

type Issuer struct {
	secret []byte
	expire time.Duration
	issuer string
}

func NewIssuer(secret string, expire time.Duration, issuer string) *Issuer {
	return &Issuer{secret: []byte(secret), expire: expire, issuer: issuer}
}

// Claims 是校验通过后的最小身份信息。
type Claims struct {
	Subject string
	Role    string
}

// Issue 签发管理员令牌，载荷与 FastAPI 保持兼容（sub + exp），并附加 role/iat/jti。
func (i *Issuer) Issue(username string) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(i.expire)
	claims := jwt.MapClaims{
		"sub":  username,
		"role": "admin",
		"iss":  i.issuer,
		"iat":  now.Unix(),
		"exp":  exp.Unix(),
		"jti":  newJTI(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(i.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, exp, nil
}

// Parse 校验签名与有效期。
func (i *Issuer) Parse(raw string) (*Claims, error) {
	parsed, err := jwt.Parse(raw, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return i.secret, nil
	}, jwt.WithExpirationRequired(), jwt.WithIssuer(i.issuer))

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpired
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalid, err)
	}

	mc, ok := parsed.Claims.(jwt.MapClaims)
	if !ok || !parsed.Valid {
		return nil, ErrInvalid
	}
	sub, _ := mc["sub"].(string)
	role, _ := mc["role"].(string)
	if sub == "" {
		return nil, ErrInvalid
	}
	if role == "" {
		role = "admin" // 兼容旧令牌（FastAPI 只发了 sub/exp）
	}
	return &Claims{Subject: sub, Role: role}, nil
}

// CheckPassword 比对 bcrypt 哈希。
func CheckPassword(hash, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}

// HashPassword 生成 bcrypt 哈希（用于初始化管理员）。
func HashPassword(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(b), err
}

func newJTI() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(b)
}
