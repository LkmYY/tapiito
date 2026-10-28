"""FastAPI ↔ Go 契约对比：同一份 MySQL 数据，两边发相同请求，逐字段比对结构与状态码。

用法：
    python tests/contract/contract_test.py

对比维度：
    1. HTTP 状态码（必须完全一致）
    2. JSON 结构签名（键路径集合 + 值类型），忽略 id/token/时间戳等动态值
    3. 明确列出差异路径，便于定位

不需要 pytest，直接运行即可。
"""
import os
from __future__ import annotations

import json
import sys
import urllib.error
import urllib.parse
import urllib.request

FASTAPI = "http://127.0.0.1:8002"
GO_PUBLIC = "http://127.0.0.1:8080"
GO_ADMIN = "http://127.0.0.1:8081"

ADMIN_USER = os.environ.get("TAPIITO_ADMIN_USER", "")
ADMIN_PASS = os.environ.get("TAPIITO_ADMIN_PASS", "")

# 值不参与比较，只比类型；时间与 id 在两库天然不同
PASS = 0
FAIL = 0
DIFFS: list[str] = []


def call(base: str, method: str, path: str, token: str | None = None, body=None):
    # 只编码查询串里的中文，保留 URL 结构字符
    url = base + urllib.parse.quote(path, safe="/?=&:%")
    data = None
    headers = {}
    if body is not None:
        data = json.dumps(body).encode()
        headers["Content-Type"] = "application/json"
    if token:
        headers["Authorization"] = f"Bearer {token}"
    req = urllib.request.Request(url, data=data, headers=headers, method=method)
    try:
        with urllib.request.urlopen(req, timeout=10) as r:
            raw = r.read().decode() or "{}"
            try:
                return r.status, json.loads(raw)
            except json.JSONDecodeError:
                return r.status, raw
    except urllib.error.HTTPError as e:
        raw = e.read().decode() or "{}"
        try:
            return e.code, json.loads(raw)
        except json.JSONDecodeError:
            return e.code, raw
    except Exception as e:  # noqa: BLE001
        return -1, {"__error__": str(e)}


def flatten(obj, prefix="") -> dict[str, str]:
    """把 JSON 展开成 {路径: 类型}，列表只看首个元素的结构。"""
    out: dict[str, str] = {}
    if isinstance(obj, dict):
        if not obj:
            out[prefix or "$"] = "empty-object"
            return out
        for k, v in obj.items():
            out.update(flatten(v, f"{prefix}.{k}" if prefix else k))
        return out
    if isinstance(obj, list):
        if not obj:
            out[prefix or "$"] = "empty-array"
            return out
        out.update(flatten(obj[0], f"{prefix}[]"))
        return out
    if obj is None:
        out[prefix or "$"] = "null"
    elif isinstance(obj, bool):
        out[prefix or "$"] = "bool"
    elif isinstance(obj, int):
        out[prefix or "$"] = "int"
    elif isinstance(obj, float):
        out[prefix or "$"] = "float"
    elif isinstance(obj, str):
        out[prefix or "$"] = "str"
    else:
        out[prefix or "$"] = type(obj).__name__
    return out


def normalize(shape: dict[str, str]) -> dict[str, str]:
    """int/float 视为数值，避免 0 与 0.0 产生噪声差异。"""
    return {k: ("num" if v in ("int", "float") else v) for k, v in shape.items()}


def compare(name: str, fa: tuple, go: tuple):
    global PASS, FAIL
    fa_status, fa_body = fa
    go_status, go_body = go

    problems = []
    if fa_status != go_status:
        problems.append(f"状态码 {fa_status} vs {go_status}")

    fa_shape = normalize(flatten(fa_body))
    go_shape = normalize(flatten(go_body))

    missing = sorted(set(fa_shape) - set(go_shape))
    extra = sorted(set(go_shape) - set(fa_shape))
    typed = sorted(k for k in set(fa_shape) & set(go_shape) if fa_shape[k] != go_shape[k])

    if missing:
        problems.append("Go 缺少字段: " + ", ".join(missing))
    if extra:
        problems.append("Go 多出字段: " + ", ".join(extra))
    if typed:
        problems.append("类型不一致: " + ", ".join(f"{k}({fa_shape[k]}→{go_shape[k]})" for k in typed))

    if problems:
        FAIL += 1
        print(f"  FAIL  {name}")
        for p in problems:
            print(f"          {p}")
        DIFFS.append(f"{name}: " + "; ".join(problems))
        if fa_status != go_status or missing:
            print(f"          FastAPI={json.dumps(fa_body, ensure_ascii=False)[:220]}")
            print(f"          Go     ={json.dumps(go_body, ensure_ascii=False)[:220]}")
    else:
        PASS += 1
        print(f"  PASS  {name}  (status={fa_status}, {len(fa_shape)} 个字段)")


def login(base: str) -> str:
    _, body = call(base, "POST", "/api/auth/loginAdmin", body={"username": ADMIN_USER, "password": ADMIN_PASS})
    if isinstance(body, dict):
        return body.get("access_token", "")
    return ""


def main() -> int:
    print("=== 准备：双方各自登录 ===")
    fa_token = login(FASTAPI)
    go_token = login(GO_ADMIN)
    if not fa_token or not go_token:
        print("FAIL 登录失败，无法继续", fa_token[:20], go_token[:20])
        return 1

    print("=== 清理上次残留 ===")
    for base, token, prefix in ((FASTAPI, fa_token, "CT-FA-"), (GO_ADMIN, go_token, "CT-GO-")):
        _, lst = call(base, "GET", f"/api/nfc?search={prefix}&page=1&page_size=50", token)
        if isinstance(lst, dict):
            for it in lst.get("items") or []:
                call(base, "DELETE", f"/api/nfc/{it['id']}", token)
                print(f"  清理 {base} -> {it.get('nfc_id')}")

    print("\n=== 1. 认证 ===")
    compare("登录成功",
            call(FASTAPI, "POST", "/api/auth/loginAdmin", body={"username": ADMIN_USER, "password": ADMIN_PASS}),
            call(GO_ADMIN, "POST", "/api/auth/loginAdmin", body={"username": ADMIN_USER, "password": ADMIN_PASS}))
    compare("登录密码错误",
            call(FASTAPI, "POST", "/api/auth/loginAdmin", body={"username": ADMIN_USER, "password": "bad"}),
            call(GO_ADMIN, "POST", "/api/auth/loginAdmin", body={"username": ADMIN_USER, "password": "bad"}))
    compare("缺字段登录",
            call(FASTAPI, "POST", "/api/auth/loginAdmin", body={"username": ADMIN_USER}),
            call(GO_ADMIN, "POST", "/api/auth/loginAdmin", body={"username": ADMIN_USER}))
    compare("无 token 访问管理接口",
            call(FASTAPI, "GET", "/api/nfc"),
            call(GO_ADMIN, "GET", "/api/nfc"))

    print("\n=== 2. 标签 CRUD ===")
    fa_create = call(FASTAPI, "POST", "/api/nfc", fa_token, {
        "nfc_id": "CT-FA-001", "name": "契约测试A", "phone": "13800138000",
        "status": "enabled", "template_type": "business_card",
        "template_payload": {"name": "张三", "title": "工程师"},
    })
    go_create = call(GO_ADMIN, "POST", "/api/nfc", go_token, {
        "nfc_id": "CT-GO-001", "name": "契约测试A", "phone": "13800138000",
        "status": "enabled", "template_type": "business_card",
        "template_payload": {"name": "张三", "title": "工程师"},
    })
    compare("创建标签", fa_create, go_create)

    fa_id = fa_create[1].get("id") if isinstance(fa_create[1], dict) else None
    go_id = go_create[1].get("id") if isinstance(go_create[1], dict) else None

    compare("标签详情", call(FASTAPI, "GET", f"/api/nfc/{fa_id}", fa_token),
            call(GO_ADMIN, "GET", f"/api/nfc/{go_id}", go_token))
    compare("标签列表分页", call(FASTAPI, "GET", "/api/nfc?page=1&page_size=5", fa_token),
            call(GO_ADMIN, "GET", "/api/nfc?page=1&page_size=5", go_token))
    compare("标签列表搜索", call(FASTAPI, "GET", "/api/nfc?search=契约&page=1&page_size=5", fa_token),
            call(GO_ADMIN, "GET", "/api/nfc?search=契约&page=1&page_size=5", go_token))
    compare("标签列表状态筛选", call(FASTAPI, "GET", "/api/nfc?status=enabled", fa_token),
            call(GO_ADMIN, "GET", "/api/nfc?status=enabled", go_token))
    compare("创建人列表", call(FASTAPI, "GET", "/api/nfc/creators", fa_token),
            call(GO_ADMIN, "GET", "/api/nfc/creators", go_token))
    compare("更新标签", call(FASTAPI, "PUT", f"/api/nfc/{fa_id}", fa_token, {"name": "改名B"}),
            call(GO_ADMIN, "PUT", f"/api/nfc/{go_id}", go_token, {"name": "改名B"}))
    compare("详情不存在", call(FASTAPI, "GET", "/api/nfc/99999999", fa_token),
            call(GO_ADMIN, "GET", "/api/nfc/99999999", go_token))
    compare("批量停用", call(FASTAPI, "POST", "/api/nfc/batch", fa_token, {"ids": [fa_id], "action": "disable"}),
            call(GO_ADMIN, "POST", "/api/nfc/batch", go_token, {"ids": [go_id], "action": "disable"}))

    print("\n=== 3. 公开接口 ===")
    compare("公开页（启用）", call(FASTAPI, "GET", "/api/public/nfc/CT-FA-001/page"),
            call(GO_PUBLIC, "GET", "/api/public/nfc/CT-GO-001/page"))
    compare("公开页（停用）", call(FASTAPI, "GET", "/api/public/nfc/CT-FA-001/page"),
            call(GO_PUBLIC, "GET", "/api/public/nfc/CT-GO-001/page"))
    compare("公开页不存在", call(FASTAPI, "GET", "/api/public/nfc/NO-SUCH-CODE/page"),
            call(GO_PUBLIC, "GET", "/api/public/nfc/NO-SUCH-CODE/page"))
    compare("展厅列表", call(FASTAPI, "GET", "/api/public/gallery?limit=5"),
            call(GO_PUBLIC, "GET", "/api/public/gallery?limit=5"))

    print("\n=== 4. 清理 ===")
    call(FASTAPI, "DELETE", f"/api/nfc/{fa_id}", fa_token)
    call(GO_ADMIN, "DELETE", f"/api/nfc/{go_id}", go_token)
    print("  已删除测试标签")

    print(f"\n结果：{PASS} 一致 / {FAIL} 有差异")
    if DIFFS:
        print("\n差异清单：")
        for d in DIFFS:
            print("  -", d)
    return 0 if FAIL == 0 else 1


if __name__ == "__main__":
    sys.exit(main())
