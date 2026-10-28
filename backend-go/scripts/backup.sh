#!/usr/bin/env bash
# MySQL 一致性备份：InnoDB 快照 + 校验和 + 保留策略
# 用法：DB_HOST=127.0.0.1 DB_NAME=nfc DB_USER=backup BACKUP_DIR=/backup ./backup.sh
set -euo pipefail

DB_HOST="${DB_HOST:-127.0.0.1}"
DB_PORT="${DB_PORT:-3306}"
DB_NAME="${DB_NAME:-nfc}"
DB_USER="${DB_USER:-root}"
BACKUP_DIR="${BACKUP_DIR:-/backup}"
RETENTION_DAYS="${RETENTION_DAYS:-7}"
LOCK_FILE="/tmp/tapiito-backup.lock"

mkdir -p "$BACKUP_DIR"

# flock 防止多个备份任务重叠
exec 9>"$LOCK_FILE"
if ! flock -n 9; then
  echo "[backup] another backup is running, exit" >&2
  exit 0
fi

# 先确认所有表都是 InnoDB，--single-transaction 对 MyISAM 无效
ENGINES=$(mysql -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" -N -B -e "
  SELECT DISTINCT ENGINE FROM information_schema.TABLES
  WHERE TABLE_SCHEMA='$DB_NAME' AND TABLE_TYPE='BASE TABLE';")
if echo "$ENGINES" | grep -qvi '^InnoDB$'; then
  echo "[backup] WARNING: non-InnoDB tables found: $ENGINES" >&2
fi

TS=$(date +%F-%H%M)
OUT="$BACKUP_DIR/${DB_NAME}-${TS}.sql.gz"

# --single-transaction：一致性快照；--quick 避免一次性缓冲全部行
# 不加 --master-data（会引入额外锁），需要 binlog 位点时另行处理
mysqldump \
  -h"$DB_HOST" -P"$DB_PORT" -u"$DB_USER" \
  --single-transaction \
  --quick \
  --skip-lock-tables \
  --routines --triggers --events \
  --default-character-set=utf8mb4 \
  --no-tablespaces \
  "$DB_NAME" | gzip > "$OUT"

sha256sum "$OUT" > "$OUT.sha256"
SIZE=$(du -h "$OUT" | cut -f1)
echo "[backup] done $OUT ($SIZE)"

# 保留策略：日备 7 天
find "$BACKUP_DIR" -name "${DB_NAME}-*.sql.gz" -mtime "+$RETENTION_DAYS" -print -delete
find "$BACKUP_DIR" -name "${DB_NAME}-*.sql.gz.sha256" -mtime "+$RETENTION_DAYS" -print -delete

# 恢复演练（每月至少一次手工执行）：
#   gunzip -c nfc-YYYY-MM-DD-HHMM.sql.gz | mysql -h$DB_HOST -u$DB_USER nfc_restore_test
#   mysql -h$DB_HOST -u$DB_USER -e "SELECT COUNT(*) FROM nfc_restore_test.nfc_tag;"
