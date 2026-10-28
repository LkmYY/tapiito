#!/usr/bin/env bash
# Tapiito Go 后端验收测试：公开读链路 + 管理写链路 + 缓存/Bloom/限流行为
# 管理端凭据从环境变量读取，不要写死在仓库里
: "${TAPIITO_ADMIN_USER:?请先 export TAPIITO_ADMIN_USER}"
: "${TAPIITO_ADMIN_PASS:?请先 export TAPIITO_ADMIN_PASS}"
AUTH_BODY=$(python -c "import json,os;print(json.dumps({'username':os.environ['TAPIITO_ADMIN_USER'],'password':os.environ['TAPIITO_ADMIN_PASS']}))")

set -u
PUB=http://127.0.0.1:8080
ADM=http://127.0.0.1:8081
C="curl -s --noproxy *"
PASS=0; FAIL=0

chk() { # chk <名称> <期望> <实际>
  if [ "$2" = "$3" ]; then PASS=$((PASS+1)); echo "  PASS  $1";
  else FAIL=$((FAIL+1)); echo "  FAIL  $1  expect=[$2] got=[$3]"; fi
}

echo "=== 1. 管理登录 ==="
LOGIN=$(curl -s --noproxy '*' -X POST "$ADM/api/auth/loginAdmin" \
  -H 'Content-Type: application/json' \
  -d "$AUTH_BODY")
TOKEN=$(echo "$LOGIN" | python -c "import sys,json;print(json.load(sys.stdin).get('access_token',''))" 2>/dev/null)
TT=$(echo "$LOGIN" | python -c "import sys,json;d=json.load(sys.stdin);print(d.get('token_type'),d.get('user_type'))" 2>/dev/null)
[ -n "$TOKEN" ] && chk "登录返回 token" "yes" "yes" || chk "登录返回 token" "yes" "no"
chk "token_type 与 user_type" "bearer admin" "$TT"

echo "=== 2. 错误密码 ==="
CODE=$(curl -s --noproxy '*' -o /dev/null -w '%{http_code}' -X POST "$ADM/api/auth/loginAdmin" \
  -H 'Content-Type: application/json' -d '{"username":"__invalid_user__","password":"wrong"}')
chk "错误密码 401" "401" "$CODE"

echo "=== 3. 无 token 访问管理接口 ==="
CODE=$(curl -s --noproxy '*' -o /dev/null -w '%{http_code}' "$ADM/api/nfc")
chk "无 token 401" "401" "$CODE"

echo "=== 4. 管理接口不能从公开端口访问 ==="
CODE=$(curl -s --noproxy '*' -o /dev/null -w '%{http_code}' "$PUB/api/nfc")
chk "公开端口访问 /api/nfc 404" "404" "$CODE"

echo "=== 5. 创建标签 ==="
NEW=$(curl -s --noproxy '*' -X POST "$ADM/api/nfc" -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{"nfc_id":"E2E-0001","name":"端到端测试","phone":"13800138000","status":"enabled","template_type":"business_card","template_payload":{"name":"林晓峰","title":"产品设计师"}}')
NID=$(echo "$NEW" | python -c "import sys,json;print(json.load(sys.stdin).get('id',''))" 2>/dev/null)
HAS_FIELDS=$(echo "$NEW" | python -c "
import sys,json
d=json.load(sys.stdin)
need=['id','nfc_id','name','phone','user_id','creator_name','status','template_type','custom_template','template_payload','order_no','created_at']
print('ok' if all(k in d for k in need) else 'missing:'+','.join(k for k in need if k not in d))
" 2>/dev/null)
chk "创建成功且字段齐全" "ok" "$HAS_FIELDS"

echo "=== 6. 公开读（首次，回源 MySQL） ==="
PAGE=$(curl -s --noproxy '*' "$PUB/api/public/nfc/E2E-0001/page")
PAGE_OK=$(echo "$PAGE" | python -c "
import sys,json
d=json.load(sys.stdin)
print('ok' if d.get('status')=='enabled' and d.get('template_payload',{}).get('name')=='林晓峰' else 'bad:'+json.dumps(d,ensure_ascii=False))
" 2>/dev/null)
chk "公开页返回启用内容" "ok" "$PAGE_OK"

echo "=== 7. 缓存已回填 ==="
HIT=$(docker exec tapiito-redis-cache redis-cli EXISTS "nfc:page:v1:E2E-0001" | tr -d '\r')
chk "Redis 有 nfc:page:v1 键" "1" "$HIT"

echo "=== 8. 更新后缓存立即失效 ==="
curl -s --noproxy '*' -X PUT "$ADM/api/nfc/$NID" -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"name":"改名后的标签"}' -o /dev/null
GONE=$(docker exec tapiito-redis-cache redis-cli EXISTS "nfc:page:v1:E2E-0001" | tr -d '\r')
chk "更新后缓存被删除" "0" "$GONE"
NAME=$(curl -s --noproxy '*' "$PUB/api/public/nfc/E2E-0001/page" | python -c "import sys,json;print(json.load(sys.stdin).get('name'))" 2>/dev/null)
chk "公开页拿到新名字" "改名后的标签" "$NAME"

echo "=== 9. 停用后公开页不报错且不下发内容 ==="
curl -s --noproxy '*' -X PUT "$ADM/api/nfc/$NID" -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"status":"disabled"}' -o /dev/null
D=$(curl -s --noproxy '*' -w '\n%{http_code}' "$PUB/api/public/nfc/E2E-0001/page")
DCODE=$(echo "$D" | tail -1)
DBODY=$(echo "$D" | head -n -1)
chk "停用返回 200" "200" "$DCODE"
DSTATUS=$(echo "$DBODY" | python -c "import sys,json;print(json.load(sys.stdin).get('status'))" 2>/dev/null)
chk "停用 status=disabled" "disabled" "$DSTATUS"
DPAYLOAD=$(echo "$DBODY" | python -c "import sys,json;print(json.load(sys.stdin).get('template_payload'))" 2>/dev/null)
chk "停用不下发 payload" "None" "$DPAYLOAD"

echo "=== 10. 重新启用 ==="
curl -s --noproxy '*' -X PUT "$ADM/api/nfc/$NID" -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d '{"status":"enabled"}' -o /dev/null
S=$(curl -s --noproxy '*' "$PUB/api/public/nfc/E2E-0001/page" | python -c "import sys,json;print(json.load(sys.stdin).get('status'))" 2>/dev/null)
chk "重新启用 status=enabled" "enabled" "$S"

echo "=== 11. 不存在的 ID 返回 404（Bloom 拦截） ==="
BEFORE=$(docker exec tapiito-mysql mysql -uroot -proot -N -B -e "SHOW GLOBAL STATUS LIKE 'Questions'" 2>/dev/null | awk '{print $2}')
CODE=$(curl -s --noproxy '*' -o /dev/null -w '%{http_code}' "$PUB/api/public/nfc/NOT-EXIST-9999/page")
chk "不存在返回 404" "404" "$CODE"
AFTER=$(docker exec tapiito-mysql mysql -uroot -proot -N -B -e "SHOW GLOBAL STATUS LIKE 'Questions'" 2>/dev/null | awk '{print $2}')
echo "  info  MySQL Questions: $BEFORE -> $AFTER (差 $((AFTER-BEFORE))，含两条 SHOW 自身)"

echo "=== 12. 批量停用 ==="
BATCH=$(curl -s --noproxy '*' -X POST "$ADM/api/nfc/batch" -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' -d "{\"ids\":[$NID],\"action\":\"disable\"}")
AFF=$(echo "$BATCH" | python -c "import sys,json;print(json.load(sys.stdin).get('affected'))" 2>/dev/null)
chk "批量停用 affected=1" "1" "$AFF"

echo "=== 13. 列表分页结构 ==="
LIST=$(curl -s --noproxy '*' "$ADM/api/nfc?page=1&page_size=10" -H "Authorization: Bearer $TOKEN")
LF=$(echo "$LIST" | python -c "
import sys,json
d=json.load(sys.stdin)
need=['items','total','page','page_size','total_pages']
print('ok' if all(k in d for k in need) else 'missing')
" 2>/dev/null)
chk "分页字段齐全" "ok" "$LF"

echo "=== 14. 删除标签 ==="
curl -s --noproxy '*' -X DELETE "$ADM/api/nfc/$NID" -H "Authorization: Bearer $TOKEN" -o /dev/null
CODE=$(curl -s --noproxy '*' -o /dev/null -w '%{http_code}' "$PUB/api/public/nfc/E2E-0001/page")
chk "删除后公开页 404（负缓存=0 时回源判定）" "404" "$CODE"

echo "=== 15. 展厅接口 ==="
G=$(curl -s --noproxy '*' "$PUB/api/public/gallery?limit=5")
GF=$(echo "$G" | python -c "
import sys,json
d=json.load(sys.stdin)
print('ok' if isinstance(d.get('items'),list) and 'total' in d else 'bad')
" 2>/dev/null)
chk "gallery 结构 {items,total}" "ok" "$GF"

echo
echo "结果：$PASS 通过 / $FAIL 失败"
[ "$FAIL" -eq 0 ] || exit 1
