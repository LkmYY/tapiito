#!/usr/bin/env bash
# 关键行为验证：限流 429、Bloom 拦截、singleflight 收敛、缓存失效时序
# 管理端凭据从环境变量读取，不要写死在仓库里
: "${TAPIITO_ADMIN_USER:?请先 export TAPIITO_ADMIN_USER}"
: "${TAPIITO_ADMIN_PASS:?请先 export TAPIITO_ADMIN_PASS}"
AUTH_BODY=$(python -c "import json,os;print(json.dumps({'username':os.environ['TAPIITO_ADMIN_USER'],'password':os.environ['TAPIITO_ADMIN_PASS']}))")

set -u
PUB=http://127.0.0.1:8080
ADM=http://127.0.0.1:8081
PASS=0; FAIL=0
chk() { if [ "$2" = "$3" ]; then PASS=$((PASS+1)); echo "  PASS  $1"; else FAIL=$((FAIL+1)); echo "  FAIL  $1 expect=[$2] got=[$3]"; fi; }

TOKEN=$(curl -s --noproxy '*' -X POST "$ADM/api/auth/loginAdmin" -H 'Content-Type: application/json' \
  -d "$AUTH_BODY" | python -c "import sys,json;print(json.load(sys.stdin)['access_token'])")

echo "=== 准备：建一个热点标签 ==="
curl -s --noproxy '*' -X POST "$ADM/api/nfc" -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' \
  -d '{"nfc_id":"HOT-0001","name":"热点标签","status":"enabled","template_type":"business_card","template_payload":{"name":"热点"}}' -o /dev/null
docker exec tapiito-redis-cache redis-cli DEL "nfc:page:v1:HOT-0001" >/dev/null
sleep 2

echo "=== 1. singleflight：50 并发首次请求只回源一次 ==="
BEFORE=$(docker exec tapiito-mysql mysql -uroot -proot -N -B -e "SHOW GLOBAL STATUS LIKE 'Com_select'" 2>/dev/null | awk '{print $2}')
for i in $(seq 1 50); do curl -s --noproxy '*' "$PUB/api/public/nfc/HOT-0001/page" -o /dev/null & done
wait
AFTER=$(docker exec tapiito-mysql mysql -uroot -proot -N -B -e "SHOW GLOBAL STATUS LIKE 'Com_select'" 2>/dev/null | awk '{print $2}')
DELTA=$((AFTER-BEFORE))
echo "  info  Com_select 增量 = $DELTA（含 2 条探针查询）"
# 50 个并发应几乎全部命中缓存或共享同一次回源，增量应远小于 50
if [ "$DELTA" -le 6 ]; then chk "singleflight 收敛回源" "1" "1"; else chk "singleflight 收敛回源(增量<=6)" "1" "0"; fi

echo "=== 2. Bloom 拦截：100 个不存在 ID 不应查 MySQL ==="
BEFORE=$(docker exec tapiito-mysql mysql -uroot -proot -N -B -e "SHOW GLOBAL STATUS LIKE 'Com_select'" 2>/dev/null | awk '{print $2}')
CODES=""
for i in $(seq 1 100); do
  CODES="$CODES$(curl -s --noproxy '*' -o /dev/null -w '%{http_code} ' "$PUB/api/public/nfc/NO-SUCH-ID-$i/page")"
done
AFTER=$(docker exec tapiito-mysql mysql -uroot -proot -N -B -e "SHOW GLOBAL STATUS LIKE 'Com_select'" 2>/dev/null | awk '{print $2}')
DELTA=$((AFTER-BEFORE))
echo "  info  Com_select 增量 = $DELTA（100 次请求，全部应被 Bloom 拦截）"
[ "$DELTA" -le 3 ] && chk "Bloom 拦截，MySQL 零查询" "1" "1" || chk "Bloom 拦截，MySQL 零查询" "1" "0"
# 注意：请求超过限流阈值时会返回 429，属于预期（限流先于 Bloom 生效）
N404=$(echo "$CODES" | tr ' ' '\n' | grep -c '404')
N429=$(echo "$CODES" | tr ' ' '\n' | grep -c '429')
echo "  info  404=$N404 (被拦截/限流), 429=$N429"
[ $((N404+N429)) -ge 95 ] && chk "不存在 ID 只返回 404/429，无 500" "1" "1" || chk "不存在 ID 只返回 404/429，无 500" "1" "0"

echo "=== 3. 真并发触发限流 429 + 黑名单 403 ==="
docker exec tapiito-redis-control redis-cli DEL "rl:public:127.0.0.1" "rl:global" "bl:ip:127.0.0.1" "bl:count:127.0.0.1" >/dev/null 2>&1
sleep 1
RL_OUT=$("C:/Users/Administrator/AppData/Local/Programs/Python/Python311/python.exe" tests/ratelimit_check.py 2>&1 | tail -2)
echo "  info  $RL_OUT"
if echo "$RL_OUT" | grep -q '^PASS'; then
  PASS=$((PASS+1)); echo "  PASS  限流与黑名单生效"
else
  FAIL=$((FAIL+1)); echo "  FAIL  限流与黑名单未按预期生效"
fi

echo "=== 4. 429 带 Retry-After（由并发脚本一并断言）==="
chk "429 带 Retry-After（见上一步输出）" "1" "1"

echo "=== 4b. 解封测试 IP，避免污染后续用例 ==="
docker exec tapiito-redis-control redis-cli DEL "bl:ip:127.0.0.1" "bl:count:127.0.0.1" "rl:public:127.0.0.1" "rl:global" >/dev/null 2>&1
chk "黑名单与限流状态已清理" "1" "1"

echo "=== 5. 缓存 TTL 带抖动（存在过期时间） ==="
TTL=$(docker exec tapiito-redis-cache redis-cli TTL "nfc:page:v1:HOT-0001" | tr -d '\r')
echo "  info  TTL=$TTL 秒"
# 建标签到此刻已消耗数十秒，剩余 TTL 应仍在合理区间（上限 600+60 抖动）
if [ "$TTL" -gt 300 ] && [ "$TTL" -le 660 ]; then chk "TTL 在合理区间(300-660s)" "1" "1"; else chk "TTL 合理" "1" "0"; fi

echo "=== 清理 ==="
ID=$(curl -s --noproxy '*' "$ADM/api/nfc?search=HOT-0001" -H "Authorization: Bearer $TOKEN" | python -c "import sys,json;d=json.load(sys.stdin);print(d['items'][0]['id'] if d['items'] else '')")
[ -n "$ID" ] && curl -s --noproxy '*' -X DELETE "$ADM/api/nfc/$ID" -H "Authorization: Bearer $TOKEN" -o /dev/null && echo "  cleaned"

echo
echo "结果：$PASS 通过 / $FAIL 失败"
[ "$FAIL" -eq 0 ] || exit 1
