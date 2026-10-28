"""Tapiito Go 后端并发压测：验证缓存、Bloom、singleflight 在混合负载下的真实行为。

场景：
    1. 建 10 个不同类型的标签
    2. 冷启动：清空缓存后单请求 → 确认"先查库再写 Redis"
    3. 混合并发 300：10 标签 × 20 并发 + 100 个不存在 ID
    4. 热缓存 200 并发：确认命中缓存不再回源
    5. Bloom 过滤：100 个不存在 ID 的 MySQL 查询增量

指标：状态码分布、延迟分位、MySQL Com_select 增量、Redis 命中/未命中
"""
import os
from __future__ import annotations

import concurrent.futures
import json
import statistics
import subprocess
import time
import urllib.error
import urllib.parse
import urllib.request

PUB = "http://127.0.0.1:8080"
ADM = "http://127.0.0.1:8081"
ADMIN = (os.environ.get("TAPIITO_ADMIN_USER", ""), os.environ.get("TAPIITO_ADMIN_PASS", ""))

TYPES = [
    "business_card", "child_info", "recipe", "pet_tag", "company_website",
    "product_display", "nightclub", "anime_card", "emergency_contact", "custom",
]
PER_TAG_CONCURRENCY = 20
MISSING_TAG_CONCURRENCY = 100


def http(method: str, url: str, body=None, token=None, timeout=20):
    data = None
    headers = {}
    if body is not None:
        data = json.dumps(body).encode()
        headers["Content-Type"] = "application/json"
    if token:
        headers["Authorization"] = f"Bearer {token}"
    req = urllib.request.Request(url, data=data, headers=headers, method=method)
    try:
        with urllib.request.urlopen(req, timeout=timeout) as r:
            return r.status, json.loads(r.read().decode() or "{}")
    except urllib.error.HTTPError as e:
        try:
            return e.code, json.loads(e.read().decode() or "{}")
        except Exception:
            return e.code, {}
    except Exception as e:  # noqa: BLE001
        return -1, {"__error__": str(e)}


def sh(cmd: list[str]) -> str:
    return subprocess.run(cmd, capture_output=True, text=True).stdout.strip()


def mysql_selects() -> int:
    out = sh(["docker", "exec", "tapiito-mysql", "mysql", "-uroot", "-proot", "-N", "-B",
              "-e", "SHOW GLOBAL STATUS LIKE 'Com_select'"])
    return int(out.split()[-1]) if out else -1


def redis_stat(field: str) -> int:
    out = sh(["docker", "exec", "tapiito-redis-cache", "redis-cli", "INFO", "stats"])
    for line in out.splitlines():
        if line.startswith(field + ":"):
            return int(line.split(":")[1].strip())
    return -1


def redis_exists(key: str) -> bool:
    out = sh(["docker", "exec", "tapiito-redis-cache", "redis-cli", "EXISTS", key])
    return out.strip() == "1"


def login() -> str:
    _, body = http("POST", f"{ADM}/api/auth/loginAdmin",
                   {"username": ADMIN[0], "password": ADMIN[1]})
    return body.get("access_token", "")


def hit(nfc_id: str):
    """发起一次公开页请求，返回 (状态码, 耗时毫秒)。"""
    url = f"{PUB}/api/public/nfc/{urllib.parse.quote(nfc_id)}/page"
    t0 = time.perf_counter()
    code, _ = http("GET", url, timeout=30)
    return code, (time.perf_counter() - t0) * 1000


def pct(values: list[float], p: float) -> float:
    if not values:
        return 0.0
    ordered = sorted(values)
    idx = min(len(ordered) - 1, int(len(ordered) * p / 100))
    return ordered[idx]


def main() -> int:
    token = login()
    if not token:
        print("登录失败")
        return 1

    print("=" * 62)
    print("步骤 1：清理旧数据并创建 10 个不同类型的标签")
    print("=" * 62)
    _, lst = http("GET", f"{ADM}/api/nfc?search=PERF-&page=1&page_size=50", token=token)
    for it in (lst.get("items") or []):
        http("DELETE", f"{ADM}/api/nfc/{it['id']}", token=token)

    ids: list[str] = []
    for i, ttype in enumerate(TYPES, 1):
        nfc_id = f"PERF-{i:02d}"
        code, body = http("POST", f"{ADM}/api/nfc", {
            "nfc_id": nfc_id, "name": f"压测标签{i}", "status": "enabled",
            "template_type": ttype,
            "template_payload": {"name": f"压测{i}", "title": ttype},
        }, token=token)
        if code != 200:
            print(f"  创建失败 {nfc_id}: {code} {body}")
            return 1
        ids.append(nfc_id)
    print(f"  已创建 {len(ids)} 个：{', '.join(ids)}")

    print()
    print("=" * 62)
    print("步骤 2：冷启动验证 —— 未访问过的标签应先查库再写 Redis")
    print("=" * 62)
    for k in ids:
        sh(["docker", "exec", "tapiito-redis-cache", "redis-cli", "DEL", f"nfc:page:v1:{k}"])
        sh(["docker", "exec", "tapiito-redis-cache", "redis-cli", "DEL", f"nfc:miss:v1:{k}"])
    before_sel = mysql_selects()
    before_hit = redis_stat("keyspace_hits")
    before_miss = redis_stat("keyspace_misses")

    code, ms = hit("PERF-01")
    after_sel = mysql_selects()
    cached = redis_exists("nfc:page:v1:PERF-01")
    print(f"  首次请求 PERF-01: status={code} 耗时={ms:.1f}ms")
    print(f"  MySQL Com_select 增量 = {after_sel - before_sel}  （>0 说明确实回源查库）")
    print(f"  Redis 已写入 nfc:page:v1:PERF-01 = {cached}  （Flase 则缓存未生效）")
    print(f"  Redis keyspace: hits+{redis_stat('keyspace_hits') - before_hit} "
          f"misses+{redis_stat('keyspace_misses') - before_miss}")
    ok_step2 = (after_sel - before_sel) >= 1 and cached

    print()
    print("=" * 62)
    print("步骤 3：再次请求同一标签 —— 应直接走缓存，不再查库")
    print("=" * 62)
    before_sel = mysql_selects()
    before_hit = redis_stat("keyspace_hits")
    for _ in range(5):
        code, ms = hit("PERF-01")
    after_sel = mysql_selects()
    print(f"  5 次请求 status={code} 平均耗时={ms:.1f}ms")
    print(f"  MySQL Com_select 增量 = {after_sel - before_sel}  （应为 0）")
    print(f"  Redis keyspace_hits 增量 = {redis_stat('keyspace_hits') - before_hit}  （应为 5）")
    ok_step3 = (after_sel - before_sel) <= 1 and (redis_stat("keyspace_hits") - before_hit) == 5

    print()
    print("=" * 62)
    print(f"步骤 4：混合并发 —— {len(ids)} 标签 × {PER_TAG_CONCURRENCY} 并发"
          f" + {MISSING_TAG_CONCURRENCY} 个不存在 ID = "
          f"{len(ids) * PER_TAG_CONCURRENCY + MISSING_TAG_CONCURRENCY} 请求同时打")
    print("=" * 62)
    for k in ids:
        sh(["docker", "exec", "tapiito-redis-cache", "redis-cli", "DEL", f"nfc:page:v1:{k}"])
    before_sel = mysql_selects()
    before_hit = redis_stat("keyspace_hits")
    before_miss = redis_stat("keyspace_misses")

    tasks: list[tuple[str, bool]] = []
    for k in ids:
        tasks.extend([(k, True)] * PER_TAG_CONCURRENCY)
    tasks.extend((f"MISS-{i:03d}", False) for i in range(MISSING_TAG_CONCURRENCY))

    t0 = time.perf_counter()
    with concurrent.futures.ThreadPoolExecutor(max_workers=len(tasks)) as pool:
        results = list(pool.map(lambda t: hit(t[0]), tasks))
    elapsed = (time.perf_counter() - t0) * 1000

    after_sel = mysql_selects()
    after_hit = redis_stat("keyspace_hits")
    after_miss = redis_stat("keyspace_misses")

    codes = [c for c, _ in results]
    lat = [ms for _, ms in results]
    exist_lat = [ms for (_, ms), (_, is_exist) in zip(results, tasks) if is_exist]
    miss_lat = [ms for (_, ms), (_, is_exist) in zip(results, tasks) if not is_exist]

    from collections import Counter
    dist = Counter(codes)
    print(f"  总请求 {len(tasks)}，总耗时 {elapsed:.0f}ms，QPS≈{len(tasks) / (elapsed / 1000):.0f}")
    print(f"  状态码分布: {dict(sorted(dist.items()))}")
    print(f"  延迟(全部): p50={pct(lat, 50):.0f}ms p95={pct(lat, 95):.0f}ms p99={pct(lat, 99):.0f}ms max={max(lat):.0f}ms")
    print(f"  延迟(存在): p50={pct(exist_lat, 50):.0f}ms p95={pct(exist_lat, 95):.0f}ms")
    print(f"  延迟(不存在): p50={pct(miss_lat, 50):.0f}ms p95={pct(miss_lat, 95):.0f}ms")
    print(f"  MySQL Com_select 增量 = {after_sel - before_sel}  （10 个冷标签各回源一次 → 期望≈10）")
    print(f"  Redis hits+{after_hit - before_hit}  misses+{after_miss - before_miss}")
    ok_step4 = dist.get(200, 0) == len(ids) * PER_TAG_CONCURRENCY and dist.get(404, 0) == MISSING_TAG_CONCURRENCY

    print()
    print("=" * 62)
    print("步骤 5：热缓存复压 —— 再打 200 个请求，应全部走 Redis 零回源")
    print("=" * 62)
    before_sel = mysql_selects()
    before_hit = redis_stat("keyspace_hits")
    tasks_hot = [(k, True) for k in ids for _ in range(PER_TAG_CONCURRENCY)]
    t0 = time.perf_counter()
    with concurrent.futures.ThreadPoolExecutor(max_workers=len(tasks_hot)) as pool:
        results_hot = list(pool.map(lambda t: hit(t[0]), tasks_hot))
    elapsed_hot = (time.perf_counter() - t0) * 1000
    after_sel = mysql_selects()
    after_hit = redis_stat("keyspace_hits")
    lat_hot = [ms for _, ms in results_hot]
    dist_hot = Counter(c for c, _ in results_hot)
    print(f"  总请求 {len(tasks_hot)}，总耗时 {elapsed_hot:.0f}ms，QPS≈{len(tasks_hot) / (elapsed_hot / 1000):.0f}")
    print(f"  状态码分布: {dict(sorted(dist_hot.items()))}")
    print(f"  延迟: p50={pct(lat_hot, 50):.0f}ms p95={pct(lat_hot, 95):.0f}ms p99={pct(lat_hot, 99):.0f}ms")
    print(f"  MySQL Com_select 增量 = {after_sel - before_sel}  （应为 0）")
    print(f"  Redis keyspace_hits 增量 = {after_hit - before_hit}")
    ok_step5 = (after_sel - before_sel) == 0 and dist_hot.get(200, 0) == len(tasks_hot)

    print()
    print("=" * 62)
    print("步骤 6：Bloom 过滤 —— 400 个不存在的 ID 不应产生任何 MySQL 查询")
    print("=" * 62)
    before_sel = mysql_selects()
    tasks_bloom = [(f"NOBLOOM-{i:04d}", False) for i in range(400)]
    t0 = time.perf_counter()
    with concurrent.futures.ThreadPoolExecutor(max_workers=100) as pool:
        results_bloom = list(pool.map(lambda t: hit(t[0]), tasks_bloom))
    elapsed_bloom = (time.perf_counter() - t0) * 1000
    after_sel = mysql_selects()
    dist_bloom = Counter(c for c, _ in results_bloom)
    lat_bloom = [ms for _, ms in results_bloom]
    print(f"  总请求 400，总耗时 {elapsed_bloom:.0f}ms，QPS≈{400 / (elapsed_bloom / 1000):.0f}")
    print(f"  状态码分布: {dict(sorted(dist_bloom.items()))}")
    print(f"  延迟: p50={pct(lat_bloom, 50):.1f}ms p95={pct(lat_bloom, 95):.1f}ms p99={pct(lat_bloom, 99):.1f}ms")
    print(f"  MySQL Com_select 增量 = {after_sel - before_sel}  （应为 0，全部被 Bloom 拦截）")
    ok_step6 = (after_sel - before_sel) == 0 and dist_bloom.get(404, 0) == 400

    print()
    print("=" * 62)
    print("结论")
    print("=" * 62)
    checks = [
        ("冷启动先查库再写缓存", ok_step2),
        ("已缓存数据直接命中 Redis", ok_step3),
        ("混合并发无崩溃、状态码全对", ok_step4),
        ("热缓存零回源", ok_step5),
        ("Bloom 拦截 400 个无效 ID", ok_step6),
    ]
    for name, ok in checks:
        print(f"  {'PASS' if ok else 'FAIL'}  {name}")

    print("\n  清理测试标签...")
    _, lst = http("GET", f"{ADM}/api/nfc?search=PERF-&page=1&page_size=50", token=token)
    for it in (lst.get("items") or []):
        http("DELETE", f"{ADM}/api/nfc/{it['id']}", token=token)
    for k in ids:
        sh(["docker", "exec", "tapiito-redis-cache", "redis-cli", "DEL", f"nfc:page:v1:{k}"])
    print("  完成")

    return 0 if all(ok for _, ok in checks) else 1


if __name__ == "__main__":
    raise SystemExit(main())
