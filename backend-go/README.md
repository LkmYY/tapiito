# Tapiito Go 后端

FastAPI 后端的 Go/Gin 替代实现。目标是 400 QPS、10 万级名片、读多写少场景，保留 2 倍压测余量。

## 目录

```
cmd/
  api/            双入口：:8080 公开 / :8081 管理
  worker/         访问统计消费者（Redis Stream → MySQL）
  bloom-builder/  Bloom 轮换构建（cron 触发，跑完退出）
internal/
  config/         环境变量与生产校验
  db/             MySQL 连接池
  dbgate/         读写并发闸门（有界队列 + 503）
  cache/          Cache Aside + singleflight + 负缓存
  bloom/          双 Bloom 轮换 + delta 增量集合
  ratelimit/      Redis 令牌桶 + 黑名单 + 本地降级
  auth/           JWT（HS256）+ bcrypt
  tag/            名片领域：公开读、管理写
  stats/          访问统计异步写入
  http/           路由、中间件、handler
migrations/       golang-migrate SQL
deploy/           compose（dev）、nginx、proxy_params
scripts/          backup.sh（mysqldump 一致性备份）
```

## 本地启动

```bash
# 1. 依赖：MySQL 8 + Redis Stack（缓存/控制分开）
docker compose -f deploy/docker-compose.dev.yml up -d

# 2. 建表
migrate -path migrations -database "mysql://nfc_user:nfc_pass@tcp(127.0.0.1:3306)/nfc" up

# 3. 启动
cp .env.example .env    # 按需修改
go run ./cmd/api
```

- 公开 API：<http://127.0.0.1:8080>（`/healthz`、`/api/public/*`）
- 管理 API：<http://127.0.0.1:8081>（`/api/auth/loginAdmin`、`/api/nfc*`）

## 核心约束（写代码时不能破）

1. **绝不先写 Redis 再写 DB**：DB 提交后才失效缓存。
2. **Bloom 判不存在才 404**；Bloom 未就绪或 Redis 故障一律 503，不能误判。
3. **HTTP 并发 ≠ MySQL 并发**：所有 SQL 必须经过 `dbgate`，闸门满返回 503。
4. **公开 DTO 不含 phone / user_id / 审计字段**。
5. **管理接口只在 :8081**，公开入口不注册任何管理能力。
6. 缓存 key 带版本前缀（`nfc:page:v1:`），发布靠改前缀失效，不用 `FLUSHALL`。

## 已实现

| 模块 | 状态 |
|---|---|
| 双入口 + 中间件（RequestID / RealIP / 黑名单 / 限流 / JWT） | ✅ |
| Cache Aside + singleflight + 负缓存 + 抖动 TTL | ✅ |
| 双 Bloom 轮换 + delta（新增不误判） | ✅ |
| 读写闸门（有界队列、超时 503） | ✅ |
| 公开读：`/api/public/nfc/{id}/page`、`/api/public/gallery` | ✅ |
| 管理：登录、列表、详情、创建、更新、删除、批量、创建人 | ✅ |
| 访问统计 Stream 异步写入 + worker 批量落库 + 日聚合 | ✅ |
| migration（初始建表 + 索引 + 日聚合表） | ✅ |
| 备份脚本 / nginx 隔离 / Dockerfile | ✅ |

## 未完成（下一步）

- 公开写接口：`/api/public/nfc/{id}/note`、`photo`、`submit`、`/api/orders`、`/api/feedback` 目前返回 **501**，尚未迁移。
- 订单（列表/转换/删除）、反馈（提交/采纳/管理列表）、统计查询、OSS 上传、Twilio 短信：未迁移。
- 压测脚本（400 QPS 持续 10 分钟）尚未执行。
- 服务默认监听 `0.0.0.0`，生产需靠 nginx + 防火墙收敛，或增加 `BIND_HOST` 配置。

## 验证

```bash
go build ./...   # 编译
go vet ./...     # 静态检查
go test ./...    # 单测（闸门并发/队列/超时、生产配置校验）

# 端到端（需 MySQL + Redis Stack 已启动、api 正在运行）
bash tests/e2e.sh        # 19 项：登录/鉴权/CRUD/缓存失效/停用/删除/分页/展厅
bash tests/behavior.sh   #  7 项：singleflight、Bloom 拦截、限流 429、黑名单 403、TTL 抖动
bash tests/failover.sh   #  6 项：控制 Redis 挂掉 → 503 且不打爆 MySQL，恢复后自动回正常
python tests/ratelimit_check.py   # 真并发限流（80 并发）

# 契约对比：FastAPI 与 Go 连同一个 MySQL，逐字段比对结构与状态码
# 前置：FastAPI 起在 8002（DATABASE_URL 指向同一库），Go 起在 8080/8081
python tests/contract/contract_test.py
```

实测结果（2026-10-22，MySQL 8 + Redis Stack）：

| 项目 | 结果 |
|---|---|
| 单测 | ✅ 全部通过 |
| e2e | 19/19 |
| 行为验证 | 7/7 |
| 故障演练 | 6/6 |
| **契约对比（FastAPI ↔ Go）** | **17/17 一致**（状态码 + 全部字段结构） |
| 100 次不存在 ID | 全部 404，MySQL `Com_select` 增量 **1**（Bloom 全部拦截） |
| 50 并发同一热点 ID | `Com_select` 增量 **2**（singleflight 收敛） |
| 80 并发读 | 20 通过 / 29×429 / 31×403（黑名单按设计触发） |
| 控制 Redis 停机 | 20/20 返回 503，`Com_select` 增量 **1**（未击穿） |

### 并发压测（2026-10-22，Go 压测工具，服务与压测同机）

工具：`go run ./cmd/loadtest -ids LOAD-1,...,LOAD-10 -missing 100 -c 200 -d 10s`
（Go 原生并发，避免 Python GIL 让客户端成为瓶颈）

| 场景 | QPS | p50 | p95 | p99 | 状态码 |
|---|---|---|---|---|---|
| 热缓存 · 50 并发 | 3112 | 15ms | 22ms | 29ms | 200 × 24936 |
| 热缓存 · 100 并发 | 3615 | 27ms | 36ms | 41ms | 200 × 28969 |
| 热缓存 · 200 并发 | 3631 | 54ms | 69ms | 84ms | 200 × 29219 |
| 混合流量（91% 是不存在 ID）· 200 并发 | 3466 | 56ms | 74ms | 87ms | 200:3071 / 404:31696 |
| 冷缓存突刺 · 200 并发 | 3493 | 54ms | 75ms | 89ms | 200 × 10634 |

目标 400 QPS，实测 **3600+**，约为目标 9 倍（同机压测，客户端也抢 CPU）。

**关键指标（MySQL general log 精确统计）**

| 验证点 | 结果 |
|---|---|
| 热缓存 11256 请求，对 `nfc_tag` 的业务 SQL | **0 次** |
| 10617 个不存在 ID，对 `nfc_tag` 的业务 SQL | **0 次**（Bloom 全拦） |
| 冷缓存 10634 请求 → 数据库回源次数 | **10 次**（singleflight 收敛 99.9%） |
| 未缓存标签首次访问 | `Com_select` +1，随后 `nfc:page:v1:{id}` 出现 → 先查库再写缓存 |
| 已缓存标签再次访问 | `keyspace_hits` 精确 +1/次，`Com_select` 不变 |

**压测工具本身踩的坑**：`resp.Body.Close()` 前必须 `io.Copy(io.Discard, resp.Body)` 读干响应体，
否则连接无法复用，短连接会迅速堆出上万个 TIME_WAIT，在 Windows 上耗尽动态端口，
表现为「大量连接错误 + p99 飙到 8 秒」——很容易被误判成服务端崩溃，实则压测端问题。

### 联调中发现并修复的问题

1. **`BF.EXISTS` 返回值类型**：RedisBloom 在 RESP3 下返回 boolean，用 `.Int()` 解析会报
   `unexpected type=bool`，导致公开读全部 503。改为 `.Result()` + 类型兼容。
2. **限流 Lua 脚本语法错误**：`tokens = burst ts = now` 单行赋值在 Lua 沙箱中报
   `nonexistent global variable 'burst'`，异常被降级逻辑吞掉，表现为「限流完全不生效且无日志」。
   已改为逐行赋值，并把返回值解析改为类型兼容。
3. **singleflight 吞掉错误**：`GetOrLoad` 丢弃了 `sf.Do` 的 error，删除标签后返回 500 而不是 404。
4. **健康检查用了启动期 context**：60 秒后所有探活永久 503，K8s 会反复重启 Pod。改为独立 context。

> 第 2 条最值得注意：**降级逻辑会把真实错误掩盖成"正常放行"**，联调时一定要用真并发验证限流，
> 串行 curl 永远测不出来（令牌被逐步补满）。

