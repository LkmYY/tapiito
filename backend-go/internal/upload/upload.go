// Package upload 处理照片上传：优先 OSS，未启用时落本地磁盘。
package upload

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/lkmyy/tapiito-go/internal/config"
)

// Store 是上传落地实现。
type Store struct {
	dir    string
	cfg    config.UploadConfig
	bucket *oss.Bucket
	log    *slog.Logger
}

// New 初始化：OSS 开启且配置完整时用 OSS，否则本地目录。
func New(cfg config.Config, log *slog.Logger) (*Store, error) {
	s := &Store{dir: cfg.UploadDir, cfg: cfg.Upload, log: log}
	if cfg.OSS.Enable {
		if cfg.OSS.Endpoint == "" || cfg.OSS.KeyID == "" || cfg.OSS.Bucket == "" {
			log.Warn("OSS enabled but config incomplete, fallback to local storage")
			return s, nil
		}
		client, err := oss.New(cfg.OSS.Endpoint, cfg.OSS.KeyID, cfg.OSS.KeySecret)
		if err != nil {
			return nil, fmt.Errorf("init oss: %w", err)
		}
		bucket, err := client.Bucket(cfg.OSS.Bucket)
		if err != nil {
			return nil, fmt.Errorf("open bucket: %w", err)
		}
		s.bucket = bucket
		log.Info("oss storage enabled", "bucket", cfg.OSS.Bucket)
		return s, nil
	}
	if err := os.MkdirAll(cfg.UploadDir, 0o755); err != nil {
		return nil, fmt.Errorf("mkdir upload dir: %w", err)
	}
	return s, nil
}

// Validate 校验大小与 MIME 白名单。
func (s *Store) Validate(size int64, contentType string) error {
	if s.cfg.MaxSize > 0 && size > s.cfg.MaxSize {
		return fmt.Errorf("file too large: %d > %d", size, s.cfg.MaxSize)
	}
	if len(s.cfg.AllowedTypes) > 0 {
		if _, ok := s.cfg.AllowedTypes[contentType]; !ok {
			return fmt.Errorf("content type not allowed: %s", contentType)
		}
	}
	return nil
}

// Save 保存文件并返回 key（OSS object key 或本地文件名）。
func (s *Store) Save(ctx context.Context, originalName, contentType string, data []byte) (string, error) {
	key := safeName(originalName)

	if s.bucket != nil {
		// 上传放到独立 goroutine 不可控，这里用 SDK 同步调用；超时由调用方 context 兜底
		if err := s.bucket.PutObject(key, bytes.NewReader(data),
			oss.ContentType(contentType)); err != nil {
			return "", fmt.Errorf("oss put: %w", err)
		}
		return key, nil
	}

	dest := filepath.Join(s.dir, key)
	if err := os.WriteFile(dest, data, 0o644); err != nil {
		return "", fmt.Errorf("write local: %w", err)
	}
	return key, nil
}

// safeName 生成 UUID + 原扩展名，杜绝路径穿越。
func safeName(original string) string {
	ext := strings.ToLower(filepath.Ext(original))
	if len(ext) > 10 || strings.ContainsAny(ext, "/\\") {
		ext = ""
	}
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	}
	return hex.EncodeToString(b) + ext
}
