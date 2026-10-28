// Package db 管理 MySQL 连接池。
package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/lkmyy/tapiito-go/internal/config"
)

// Open 创建连接池并按配置收敛参数。
func Open(cfg config.DBConfig) (*sql.DB, error) {
	conn, err := sql.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(cfg.MaxOpenConns)
	conn.SetMaxIdleConns(cfg.MaxIdleConns)
	conn.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	conn.SetConnMaxIdleTime(2 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := conn.PingContext(ctx); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("ping mysql: %w", err)
	}
	return conn, nil
}
