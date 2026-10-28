"""真并发限流验证：用线程池同时发出请求，避免 shell 启动进程的延迟稀释限流效果。

同时验证 429 响应带 Retry-After 头、以及累计超限后触发黑名单 403。
"""
import collections
import concurrent.futures
import sys
import urllib.error
import urllib.request

URL = "http://127.0.0.1:8080/api/public/nfc/RATE-PROBE/page"
N = 80


def hit(_):
    try:
        with urllib.request.urlopen(URL, timeout=5) as r:
            return r.status, r.headers.get("Retry-After")
    except urllib.error.HTTPError as e:
        return e.code, e.headers.get("Retry-After")
    except Exception:
        return -1, None


with concurrent.futures.ThreadPoolExecutor(max_workers=N) as pool:
    results = list(pool.map(hit, range(N)))

codes = [c for c, _ in results]
stat = collections.Counter(codes)
retry_after = next((h for c, h in results if c == 429 and h), None)

print(f"total={len(codes)} " + " ".join(f"{k}:{v}" for k, v in sorted(stat.items())))
print(f"retry_after={retry_after or ''}")

fail = []
if stat.get(429, 0) == 0:
    fail.append("未触发 429（限流未生效或并发不足）")
if stat.get(403, 0) == 0:
    fail.append("未触发 403（黑名单未生效）")
if not retry_after:
    fail.append("429 响应缺少 Retry-After 头")
if stat.get(-1, 0) or stat.get(500, 0):
    fail.append("出现连接错误或 500")

if fail:
    print("FAIL " + "；".join(fail))
    sys.exit(1)

print(f"PASS 429={stat[429]} 403={stat[403]} 正常={stat.get(404, 0)}，Retry-After={retry_after}")
