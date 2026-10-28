#!/usr/bin/env bash
# 故障演练：控制 Redis（Bloom/限流）不可用时，服务必须快速失败且不打爆 MySQL
set -u
PUB=http://127.0.0.1:8080
PASS=0; FAIL=0
chk() { if [ "$2" = "$3" ]; then PASS=$((PASS+1)); echo "  PASS  $1"; else FAIL=$((FAIL+1)); echo "  FAIL  $1 expect=[$2] got=[$3]"; fi; }
mysql_q() { docker exec tapiito-mysql mysql -uroot -proot -N -B -e "SHOW GLOBAL STATUS LIKE 'Com_select'" 2>/dev/null | awk '{print $2}'; }

restore() {
  echo "=== 恢复控制 Redis ==="
  docker start tapiito-redis-control >/dev/null 2>&1
  for i in $(seq 1 20); do
    sleep 1
    H=$(docker inspect --format='{{.State.Health.Status}}' tapiito-redis-control 2>/dev/null)
    [ "$H" = "healthy" ] && break
  done
  echo "  control redis: $H"
}
trap restore EXIT

echo "=== 1. 控制 Redis 停掉后，公开读应 503（不是 404，也不能查库）==="
docker stop tapiito-redis-control >/dev/null 2>&1
sleep 2
BEFORE=$(mysql_q)
CODES=""
for i in $(seq 1 20); do
  CODES="$CODES$(curl -s --noproxy '*' -o /dev/null -w '%{http_code} ' "$PUB/api/public/nfc/ANY-ID-$i/page")"
done
AFTER=$(mysql_q)
DELTA=$((AFTER-BEFORE))
N503=$(echo "$CODES" | tr ' ' '\n' | grep -c '503')
N404=$(echo "$CODES" | tr ' ' '\n' | grep -c '404')
N200=$(echo "$CODES" | tr ' ' '\n' | grep -c '200')
echo "  info  503=$N503 404=$N404 200=$N200  Com_select增量=$DELTA"
[ "$N404" -eq 0 ] && chk "Bloom 不可用时不误判 404" "1" "1" || chk "Bloom 不可用时不误判 404" "1" "0"
[ "$N503" -ge 15 ] && chk "快速失败返回 503" "1" "1" || chk "快速失败返回 503" "1" "0"
[ "$DELTA" -le 3 ] && chk "MySQL 未被击穿（查询增量<=3）" "1" "1" || chk "MySQL 未被击穿" "1" "0"

echo "=== 2. 健康检查报告 degraded ==="
CODE=$(curl -s --noproxy '*' -o /dev/null -w '%{http_code}' "$PUB/healthz")
chk "healthz 返回 503" "503" "$CODE"

restore
sleep 2

echo "=== 3. 恢复后自动回到正常（Bloom 指针仍在）==="
CODE=$(curl -s --noproxy '*' -o /dev/null -w '%{http_code}' "$PUB/api/public/nfc/ANY-ID-RECOVER/page")
chk "恢复后不再 503" "404" "$CODE"
H=$(curl -s --noproxy '*' -o /dev/null -w '%{http_code}' "$PUB/healthz")
chk "healthz 恢复 200" "200" "$H"

echo
echo "结果：$PASS 通过 / $FAIL 失败"
[ "$FAIL" -eq 0 ] || exit 1
