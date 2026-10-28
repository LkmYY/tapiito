package config

import (
	"os"
	"strings"
	"testing"
)

func TestLoadDevDefaults(t *testing.T) {
	clearEnv(t)
	cfg, err := Load()
	if err != nil {
		t.Fatalf("dev load failed: %v", err)
	}
	if cfg.PublicAddr != ":8080" || cfg.AdminAddr != ":8081" {
		t.Fatalf("unexpected addrs: %s %s", cfg.PublicAddr, cfg.AdminAddr)
	}
	if cfg.DB.ReadConcurrency != 80 || cfg.DB.WriteConcurrency != 20 {
		t.Fatalf("unexpected gate defaults: read=%d write=%d",
			cfg.DB.ReadConcurrency, cfg.DB.WriteConcurrency)
	}
}

// 生产环境必须拒绝弱配置：默认 JWT_SECRET、短密钥、缺管理员密码。
func TestProdRejectsWeakConfig(t *testing.T) {
	cases := []struct {
		name string
		env  map[string]string
		want string
	}{
		{
			name: "missing jwt secret",
			env:  map[string]string{"APP_ENV": "prod", "DEFAULT_ADMIN_PASSWORD": "x"},
			want: "JWT_SECRET",
		},
		{
			name: "short jwt secret",
			env:  map[string]string{"APP_ENV": "prod", "JWT_SECRET": "short", "DEFAULT_ADMIN_PASSWORD": "x"},
			want: "长度",
		},
		{
			name: "dev default secret in prod",
			env: map[string]string{
				"APP_ENV":         "prod",
				"JWT_SECRET":     "dev-only-secret-do-not-use-in-prod",
				"DEFAULT_ADMIN_PASSWORD": "x",
			},
			want: "开发默认",
		},
		{
			name: "missing admin password",
			env: map[string]string{
				"APP_ENV":    "prod",
				"JWT_SECRET": strings.Repeat("a", 40),
			},
			want: "DEFAULT_ADMIN_PASSWORD",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			clearEnv(t)
			for k, v := range tc.env {
				t.Setenv(k, v)
			}
			_, err := Load()
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tc.want)
			}
			if !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("expected error containing %q, got %q", tc.want, err.Error())
			}
		})
	}
}

func TestProdAcceptsStrongConfig(t *testing.T) {
	clearEnv(t)
	t.Setenv("APP_ENV", "prod")
	t.Setenv("JWT_SECRET", strings.Repeat("k", 48))
	t.Setenv("DEFAULT_ADMIN_PASSWORD", "a-strong-password")
	cfg, err := Load()
	if err != nil {
		t.Fatalf("prod load failed: %v", err)
	}
	if !cfg.IsProd() {
		t.Fatal("expected prod")
	}
}

func TestZeroConcurrencyRejected(t *testing.T) {
	clearEnv(t)
	t.Setenv("DB_READ_CONCURRENCY", "0")
	if _, err := Load(); err == nil {
		t.Fatal("expected error for zero read concurrency")
	}
}

func clearEnv(t *testing.T) {
	t.Helper()
	for _, k := range []string{
		"APP_ENV", "JWT_SECRET", "DEFAULT_ADMIN_PASSWORD",
		"DB_READ_CONCURRENCY", "DB_WRITE_CONCURRENCY", "RL_GLOBAL_QPS",
	} {
		if v, ok := os.LookupEnv(k); ok {
			t.Cleanup(func() { _ = os.Setenv(k, v) })
			_ = os.Unsetenv(k)
		}
	}
}
