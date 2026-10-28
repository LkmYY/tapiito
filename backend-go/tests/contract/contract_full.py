"""FastAPI ↔ Go 全量契约对比（26 个接口）。

覆盖：认证、标签 CRUD、订单（下单/列表/详情/转换/删除）、反馈（提交/已采纳/管理列表/采纳）、
统计（overview / 单标签 / 导出）、公开写（留言 / 照片列表）、公开读（落地页 / 展厅）。

两边连同一个 MySQL、用不同前缀的数据（FA-* / GO-*），逐字段比对状态码与结构。
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

ADMIN = (os.environ.get("TAPIITO_ADMIN_USER", ""), os.environ.get("TAPIITO_ADMIN_PASS", ""))

PASS = 0
FAIL = 0
DIFFS: list[str] = []


def call(base, method, path, token=None, body=None, raw=False):
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
        with urllib.request.urlopen(req, timeout=15) as r:
            payload = r.read().decode(errors="replace")
            ctype = r.headers.get("Content-Type", "")
            if raw:
                return r.status, {"ctype": ctype.split(";")[0], "has_body": len(payload) > 0}
            try:
                return r.status, json.loads(payload or "{}")
            except json.JSONDecodeError:
                return r.status, payload[:200]
    except urllib.error.HTTPError as e:
        payload = e.read().decode(errors="replace")
        try:
            return e.code, json.loads(payload or "{}")
        except json.JSONDecodeError:
            return e.code, payload[:200]
    except Exception as e:  # noqa: BLE001
        return -1, {"__error__": str(e)}


def flatten(obj, prefix=""):
    out = {}
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


def norm(shape):
    return {k: ("num" if v in ("int", "float") else v) for k, v in shape.items()}


def compare(name, fa, go):
    global PASS, FAIL
    fs, fb = fa
    gs, gb = go
    problems = []
    if fs != gs:
        problems.append(f"状态码 {fs} vs {gs}")

    fshape, gshape = norm(flatten(fb)), norm(flatten(gb))
    missing = sorted(set(fshape) - set(gshape))
    extra = sorted(set(gshape) - set(fshape))
    typed = sorted(k for k in set(fshape) & set(gshape) if fshape[k] != gshape[k])
    if missing:
        problems.append("Go 缺字段: " + ", ".join(missing[:6]))
    if extra:
        problems.append("Go 多字段: " + ", ".join(extra[:6]))
    if typed:
        problems.append("类型不一致: " + ", ".join(f"{k}({fshape[k]}→{gshape[k]})" for k in typed[:6]))

    if problems:
        FAIL += 1
        print(f"  FAIL  {name}")
        for p in problems:
            print(f"          {p}")
        DIFFS.append(f"{name}: " + "; ".join(problems))
        if fs != gs:
            print(f"          FastAPI={json.dumps(fb, ensure_ascii=False)[:180]}")
            print(f"          Go     ={json.dumps(gb, ensure_ascii=False)[:180]}")
    else:
        PASS += 1
        print(f"  PASS  {name}  (status={fs}, {len(fshape)} 字段)")


def login(base):
    _, b = call(base, "POST", "/api/auth/loginAdmin", body={"username": ADMIN[0], "password": ADMIN[1]})
    return b.get("access_token", "") if isinstance(b, dict) else ""


def cleanup(base, token, prefix):
    _, lst = call(base, "GET", f"/api/nfc?search={prefix}&page=1&page_size=50", token)
    if isinstance(lst, dict):
        for it in lst.get("items") or []:
            call(base, "DELETE", f"/api/nfc/{it['id']}", token)


def main():
    print("=== 登录与清理 ===")
    fa_t = login(FASTAPI)
    go_t = login(GO_ADMIN)
    if not fa_t or not go_t:
        print("FAIL 登录失败")
        return 1
    cleanup(FASTAPI, fa_t, "FA-")
    cleanup(GO_ADMIN, go_t, "GO-")

    print("\n=== 1. 认证（3）===")
    compare("登录成功", call(FASTAPI, "POST", "/api/auth/loginAdmin",
                          body={"username": ADMIN[0], "password": ADMIN[1]}),
            call(GO_ADMIN, "POST", "/api/auth/loginAdmin",
                 body={"username": ADMIN[0], "password": ADMIN[1]}))
    compare("登录失败", call(FASTAPI, "POST", "/api/auth/loginAdmin",
                          body={"username": ADMIN[0], "password": "x"}),
            call(GO_ADMIN, "POST", "/api/auth/loginAdmin",
                 body={"username": ADMIN[0], "password": "x"}))
    compare("缺字段 422", call(FASTAPI, "POST", "/api/auth/loginAdmin", body={"username": "a"}),
            call(GO_ADMIN, "POST", "/api/auth/loginAdmin", body={"username": "a"}))

    print("\n=== 2. 标签 CRUD（7）===")
    fa_c = call(FASTAPI, "POST", "/api/nfc", fa_t, {
        "nfc_id": "FA-TAG-01", "name": "契约A", "phone": "13800138000", "status": "enabled",
        "template_type": "business_card", "template_payload": {"name": "张三"}})
    go_c = call(GO_ADMIN, "POST", "/api/nfc", go_t, {
        "nfc_id": "GO-TAG-01", "name": "契约A", "phone": "13800138000", "status": "enabled",
        "template_type": "business_card", "template_payload": {"name": "张三"}})
    compare("创建标签", fa_c, go_c)
    fa_id = fa_c[1].get("id") if isinstance(fa_c[1], dict) else None
    go_id = go_c[1].get("id") if isinstance(go_c[1], dict) else None

    compare("标签详情", call(FASTAPI, "GET", f"/api/nfc/{fa_id}", fa_t),
            call(GO_ADMIN, "GET", f"/api/nfc/{go_id}", go_t))
    compare("标签列表", call(FASTAPI, "GET", "/api/nfc?page=1&page_size=5", fa_t),
            call(GO_ADMIN, "GET", "/api/nfc?page=1&page_size=5", go_t))
    compare("创建人列表", call(FASTAPI, "GET", "/api/nfc/creators", fa_t),
            call(GO_ADMIN, "GET", "/api/nfc/creators", go_t))
    compare("更新标签", call(FASTAPI, "PUT", f"/api/nfc/{fa_id}", fa_t, {"name": "改名"}),
            call(GO_ADMIN, "PUT", f"/api/nfc/{go_id}", go_t, {"name": "改名"}))
    compare("详情 404", call(FASTAPI, "GET", "/api/nfc/99999999", fa_t),
            call(GO_ADMIN, "GET", "/api/nfc/99999999", go_t))
    compare("批量停用", call(FASTAPI, "POST", "/api/nfc/batch", fa_t, {"ids": [fa_id], "action": "disable"}),
            call(GO_ADMIN, "POST", "/api/nfc/batch", go_t, {"ids": [go_id], "action": "disable"}))

    print("\n=== 3. 订单（5）===")
    fa_o = call(FASTAPI, "POST", "/api/orders", body={
        "template_type": "business_card", "custom_template": "", "template_payload": {"name": "订单测试"},
        "tag_name": "订单标签", "quantity": 3, "amount": 29.7,
        "contact_name": "李四", "contact_phone": "13900139000",
        "contact_address": "上海市徐汇区某路 1 号", "contact_note": "尽快"})
    go_o = call(GO_PUBLIC, "POST", "/api/orders", body={
        "template_type": "business_card", "custom_template": "", "template_payload": {"name": "订单测试"},
        "tag_name": "订单标签", "quantity": 3, "amount": 29.7,
        "contact_name": "李四", "contact_phone": "13900139000",
        "contact_address": "上海市徐汇区某路 1 号", "contact_note": "尽快"})
    compare("下单", fa_o, go_o)
    fa_oid = fa_o[1].get("id") if isinstance(fa_o[1], dict) else None
    go_oid = go_o[1].get("id") if isinstance(go_o[1], dict) else None

    compare("订单列表", call(FASTAPI, "GET", "/api/orders?page=1&page_size=5", fa_t),
            call(GO_ADMIN, "GET", "/api/orders?page=1&page_size=5", go_t))
    compare("订单详情", call(FASTAPI, "GET", f"/api/orders/{fa_oid}", fa_t),
            call(GO_ADMIN, "GET", f"/api/orders/{go_oid}", go_t))
    compare("订单转标签", call(FASTAPI, "POST", f"/api/orders/{fa_oid}/convert", fa_t, {"nfc_id": "FA-TAG-01"}),
            call(GO_ADMIN, "POST", f"/api/orders/{go_oid}/convert", go_t, {"nfc_id": "GO-TAG-01"}))
    compare("重复转换 400", call(FASTAPI, "POST", f"/api/orders/{fa_oid}/convert", fa_t, {}),
            call(GO_ADMIN, "POST", f"/api/orders/{go_oid}/convert", go_t, {}))

    print("\n=== 4. 反馈（4）===")
    fa_f = call(FASTAPI, "POST", "/api/feedback", body={"content": "这是一条测试意见内容"})
    go_f = call(GO_PUBLIC, "POST", "/api/feedback", body={"content": "这是一条测试意见内容"})
    compare("提交反馈", fa_f, go_f)
    fa_fid = fa_f[1].get("id") if isinstance(fa_f[1], dict) else None
    go_fid = go_f[1].get("id") if isinstance(go_f[1], dict) else None

    compare("内容太短 400", call(FASTAPI, "POST", "/api/feedback", body={"content": "短"}),
            call(GO_PUBLIC, "POST", "/api/feedback", body={"content": "短"}))
    compare("已采纳列表", call(FASTAPI, "GET", "/api/feedback/adopted?limit=5"),
            call(GO_PUBLIC, "GET", "/api/feedback/adopted?limit=5"))
    compare("采纳反馈", call(FASTAPI, "PATCH", f"/api/feedback/admin/{fa_fid}", fa_t, {"adopted": True}),
            call(GO_ADMIN, "PATCH", f"/api/feedback/admin/{go_fid}", go_t, {"adopted": True}))
    compare("反馈管理列表", call(FASTAPI, "GET", "/api/feedback/admin?page=1&page_size=5", fa_t),
            call(GO_ADMIN, "GET", "/api/feedback/admin?page=1&page_size=5", go_t))

    print("\n=== 5. 统计（3）===")
    compare("统计概览", call(FASTAPI, "GET", "/api/stats/overview", fa_t),
            call(GO_ADMIN, "GET", "/api/stats/overview", go_t))
    compare("单标签统计", call(FASTAPI, "GET", "/api/stats/nfc/FA-TAG-01", fa_t),
            call(GO_ADMIN, "GET", "/api/stats/nfc/GO-TAG-01", go_t))
    compare("导出 CSV", call(FASTAPI, "GET", "/api/stats/export?format=csv", fa_t, raw=True),
            call(GO_ADMIN, "GET", "/api/stats/export?format=csv", go_t, raw=True))

    print("\n=== 6. 公开读 / 公开写（5）===")
    compare("公开落地页", call(FASTAPI, "GET", "/api/public/nfc/FA-TAG-01/page"),
            call(GO_PUBLIC, "GET", "/api/public/nfc/GO-TAG-01/page"))
    compare("落地页 404", call(FASTAPI, "GET", "/api/public/nfc/NOPE-X/page"),
            call(GO_PUBLIC, "GET", "/api/public/nfc/NOPE-X/page"))
    compare("展厅列表", call(FASTAPI, "GET", "/api/public/gallery?limit=5"),
            call(GO_PUBLIC, "GET", "/api/public/gallery?limit=5"))
    compare("留言", call(FASTAPI, "POST", "/api/public/nfc/FA-TAG-01/note", body={"content": "你好"}),
            call(GO_PUBLIC, "POST", "/api/public/nfc/GO-TAG-01/note", body={"content": "你好"}))
    compare("照片列表", call(FASTAPI, "GET", "/api/public/nfc/FA-TAG-01/photos"),
            call(GO_PUBLIC, "GET", "/api/public/nfc/GO-TAG-01/photos"))

    print("\n=== 7. 清理 ===")
    call(FASTAPI, "DELETE", f"/api/orders/{fa_oid}", fa_t)
    call(GO_ADMIN, "DELETE", f"/api/orders/{go_oid}", go_t)
    call(FASTAPI, "DELETE", f"/api/nfc/{fa_id}", fa_t)
    call(GO_ADMIN, "DELETE", f"/api/nfc/{go_id}", go_t)
    print("  已清理")

    print(f"\n结果：{PASS} 一致 / {FAIL} 有差异")
    if DIFFS:
        print("\n差异清单：")
        for d in DIFFS:
            print("  -", d)
    return 0 if FAIL == 0 else 1


if __name__ == "__main__":
    sys.exit(main())
