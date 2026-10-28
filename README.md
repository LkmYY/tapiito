# Tapiito

Tapiito 是一个以 NFC 标签为入口的数字信息展示平台。用户只需提供展示内容，由平台管理员创建并维护 NFC 页面；手机碰触标签后打开 `/p/{nfc_id}`，即可查看对应的数字名片、求助信息、企业官网、产品介绍等内容。

平台不提供用户自助注册或用户端登录，所有标签、订单和展示内容统一由管理后台管理。

## 核心能力

- **NFC 标签管理**：创建、编辑、启停、批量操作、复制公开链接，按 NFC ID / 手机号 / 标签名 / 订单号搜索。
- **300 套展示模板**：10 种模板类型，每种类型 30 种独立风格，覆盖个人、家庭、企业、餐饮、宠物等场景。
- **公开展示页**：`/p/{nfc_id}`，自动记录访问统计；支持访客留言、照片上传，并可通过 Twilio 向绑定手机号发送短信/彩信通知。
- **模板展厅**：`/showroom` 按类型浏览全部模板风格，直接进入预览或下单。
- **硬件订单**：`/order` 选择模板、填写展示内容与邮寄信息；管理员在 `/admin/orders` 将订单转换为 NFC 标签。
- **意见反馈**：`/feedback` 提交意见，管理后台可查看并将意见标记为“已采纳”，公开页展示已采纳内容。
- **访问统计**：按标签查看访问量、设备类型与来源信息，支持统计导出。
- **图片存储**：默认保存到本地 `uploads/`，也可切换阿里云 OSS。
- **法务与关于页面**：`/about`、`/privacy`、`/disclaimer`。

## 模板体系

当前模板类型：

| 类型 | 说明 |
| --- | --- |
| 名片 | 个人/商务数字名片 |
| 求助信息卡 | 儿童、老人等求助信息 |
| 紧急联系人 | 医疗、SOS、家庭联系人 |
| 企业官网 | 企业介绍、业务展示 |
| 产品展示 | 商品详情、价格与卖点 |
| 菜谱 | 餐饮菜单与菜品介绍 |
| 自定义企业 | 企业自定义内容模块 |
| 夜店交际 | 夜店/社交场景 |
| 二次元闪卡 | 动漫、卡牌风格 |
| 宠物铭牌 | 宠物身份与联系人信息 |

模板注册入口：`frontend/src/templates/registry.js`。每种类型通过 manifest + `extensions/` 扩展风格，每种类型最终提供 30 种风格。新增或调整模板前请先阅读 [`docs/template-style-contract.md`](docs/template-style-contract.md)。

## 技术栈

**前端**

- Vue 3 + Vite
- Vue Router + Pinia
- Element Plus + Tailwind CSS
- `@vueuse/motion` 页面动效
- Lucide 图标

**后端**

- Go 1.24+ / Gin
- MySQL 8（唯一数据源）
- Redis Stack：缓存实例（Cache Aside）+ 控制实例（Bloom / 限流 / 黑名单）
- JWT（HS256）+ bcrypt
- 双入口隔离：`:8080` 公开 / `:8081` 管理
- 阿里云 OSS（可选，未启用时落本地）
- Twilio 短信通知（可选）

**部署**

- Docker / Docker Compose
- Nginx

## 目录结构

```text
.
├── backend-go/
│   ├── cmd/
│   │   ├── api/            # 双入口服务（公开 :8080 / 管理 :8081）
│   │   ├── worker/         # 访问统计消费者
│   │   ├── bloom-builder/  # Bloom 轮换构建
│   │   ├── devgateway/     # 本地联调网关（分流公开/管理）
│   │   └── loadtest/       # 压测工具
│   ├── internal/
│   │   ├── tag/            # 标签领域（公开读、管理写）
│   │   ├── order/          # 订单与订单转标签
│   │   ├── feedback/       # 意见反馈
│   │   ├── stats/          # 访问统计查询
│   │   ├── cache/          # Cache Aside + singleflight
│   │   ├── bloom/          # 双 Bloom 轮换 + delta
│   │   ├── dbgate/         # MySQL 读写并发闸门
│   │   ├── ratelimit/      # 令牌桶 + 黑名单
│   │   ├── auth/           # JWT / bcrypt
│   │   ├── upload/         # OSS / 本地存储
│   │   └── sms/            # Twilio 发送
│   ├── migrations/         # golang-migrate SQL
│   ├── tests/              # e2e / behavior / failover / contract / perf
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── pages/          # 首页、展厅、订单、公开页、管理后台
│   │   ├── pages/templates/# 模板展示组件
│   │   ├── templates/      # 模板注册表、字段 schema、风格扩展
│   │   ├── components/     # 导航、页脚等公共组件
│   │   ├── api/            # Axios 接口封装
│   │   └── stores/         # Pinia 状态
│   ├── public/showroom/    # 展厅 WebP 预览图
│   └── package.json
├── docs/
│   └── template-style-contract.md
├── deploy/
│   └── nginx.conf
├── docker-compose.yml
└── README.md
```

`frontend/dist/` 为构建产物，不纳入版本管理；Docker Compose 会在 `frontend-build` 服务中生成。

## 本地开发

### 1. 启动依赖（MySQL + Redis Stack）

```bash
docker compose -f backend-go/deploy/docker-compose.dev.yml up -d
```

### 2. 建表

```bash
migrate -path backend-go/migrations \
  -database "mysql://nfc_user:nfc_pass@tcp(127.0.0.1:3306)/nfc" up
```

### 3. 启动后端

```bash
cd backend-go
cp .env.example .env    # 按需修改
go run ./cmd/api
```

- 公开 API：<http://127.0.0.1:8080>（`/healthz`、`/api/public/*`）
- 管理 API：<http://127.0.0.1:8081>（`/api/auth/loginAdmin`、`/api/nfc*` 等）

前端本地联调需要一个合并入口（公开与管理分属两个端口）：

```bash
go run ./cmd/devgateway     # :8090，按路径+方法分流
```

> `.env.example` 中的 Twilio / OSS 配置为空的占位符；生产环境必须显式设置
> `JWT_SECRET` 与 `DEFAULT_ADMIN_PASSWORD`，否则服务拒绝启动。

### 2. 启动前端

```bash
cd frontend
npm install
npm run dev
```

前端默认地址：<http://localhost:5173>

Vite 已将 `/api` 和 `/uploads` 代理到 `http://localhost:8000`。

### 3. 登录管理后台

访问 <http://localhost:5173/login>。

管理员账号由环境变量 `DEFAULT_ADMIN_USERNAME` / `DEFAULT_ADMIN_PASSWORD` 配置，首次登录时会自动创建或同步。生产环境请务必修改默认账号密码，并设置足够强的 `JWT_SECRET`。

## 环境变量

| 变量 | 说明 | 默认值 |
| --- | --- | --- |
| `APP_ENV` | 运行环境，生产环境请设为 `prod` | `dev` |
| `JWT_SECRET` | JWT 签名密钥，生产环境必须修改 | `change-this-in-production` |
| `JWT_ACCESS_TOKEN_EXPIRE_MINUTES` | 登录令牌有效期（分钟） | `60` |
| `DEFAULT_ADMIN_USERNAME` | 默认管理员账号 | 见 `config.py` |
| `DEFAULT_ADMIN_PASSWORD` | 默认管理员密码 | 见 `config.py` |
| `DB_URL` / `DATABASE_URL` | 数据库连接串 | `sqlite:///./nfc.db` |
| `DB_ECHO` | 是否打印 SQL | `0` |
| `OSS_ENABLE` | 是否启用阿里云 OSS | `0` |
| `OSS_ENDPOINT` | OSS Endpoint | 空 |
| `OSS_ACCESS_KEY_ID` | OSS AccessKey ID | 空 |
| `OSS_ACCESS_KEY_SECRET` | OSS AccessKey Secret | 空 |
| `OSS_BUCKET_NAME` | OSS Bucket | 空 |
| `UPLOAD_DIR` | 本地上传目录 | `uploads` |
| `MAX_UPLOAD_SIZE` | 单文件大小上限（字节） | `10485760` |
| `ALLOWED_UPLOAD_TYPES` | 允许的 MIME 类型 | `image/jpeg,image/png,image/webp` |
| `TWILIO_ACCOUNT_SID` | Twilio Account SID | 空 |
| `TWILIO_AUTH_TOKEN` | Twilio Auth Token | 空 |
| `TWILIO_PHONE_NUMBER` | Twilio 发送号码 | 空 |
| `PUBLIC_BASE_URL` | 公网访问地址，用于拼接图片 URL | 空 |
| `CORS_ORIGINS` | 允许的前端来源，逗号分隔 | `*` |
| `RATE_LIMIT_ENABLED` | 是否开启限流 | `0` |
| `RATE_LIMIT_PER_MINUTE` | 每分钟请求上限 | `60` |
| `LOG_LEVEL` | 日志级别 | `INFO` |

生产环境如果 `APP_ENV` 不是 `dev` 且仍使用默认 `JWT_SECRET`，后端会拒绝启动。

## API 概览

管理端接口需要 `Authorization: Bearer <token>`；标注为公开的接口无需登录。

**认证**

- `POST /api/auth/loginAdmin`：管理员登录

**标签管理**

- `GET /api/nfc`：分页查询标签
- `POST /api/nfc`：创建标签
- `GET /api/nfc/{id}`：标签详情
- `PUT /api/nfc/{id}`：更新标签
- `DELETE /api/nfc/{id}`：删除标签
- `POST /api/nfc/batch`：批量启用/禁用/删除
- `GET /api/nfc/creators`：创建人筛选选项

**访问统计**

- `GET /api/stats/overview`：统计概览
- `GET /api/stats/nfc/{nfc_id}`：单个标签统计
- `GET /api/stats/export`：导出统计

**订单**

- `POST /api/orders`：提交硬件订单（公开）
- `GET /api/orders`：订单列表（管理端）
- `GET /api/orders/{order_id}`：订单详情（管理端）
- `POST /api/orders/{order_id}/convert`：订单转换为 NFC 标签（管理端）
- `DELETE /api/orders/{order_id}`：删除订单（管理端）

**意见反馈**

- `POST /api/feedback`：提交意见（公开）
- `GET /api/feedback/adopted`：已采纳意见（公开）
- `GET /api/feedback/admin`：管理端意见列表（管理端）
- `PATCH /api/feedback/admin/{feedback_id}`：采纳/取消采纳意见（管理端）

**公开接口**

- `GET /api/public/gallery`：模板展厅数据
- `GET /api/public/nfc/{nfc_id}/page`：公开标签页数据
- `POST /api/public/nfc/{nfc_id}/note`：提交留言
- `POST /api/public/nfc/{nfc_id}/photo`：上传照片
- `GET /api/public/nfc/{nfc_id}/photos`：照片列表
- `POST /api/public/nfc/{nfc_id}/submit`：提交联系/通知信息

完整请求与响应结构以 `/docs` 的 Swagger 文档为准。

## 测试与构建

```bash
# 后端测试
cd backend
pytest

# 前端生产构建
cd frontend
npm run build
```

前端没有单独的 `npm test` 脚本；提交前至少运行一次 `npm run build`，确认模板与路由可正常打包。

## Docker 部署

```bash
# 在项目根目录创建 .env，覆盖数据库密码、JWT_SECRET、默认管理员等配置
docker compose up -d --build
```

Compose 会启动 MySQL、Redis Stack（缓存 + 控制）、Go 后端、统计 worker、前端构建容器和 Nginx。
Nginx 配置位于 `deploy/nginx.conf`：单域名下按「方法 + 路径」把 `/api/*` 分流到 `api:8080`（公开）
或 `api:8081`（管理）；生产建议拆成 `api.*` 与 `admin.*` 两个域名并给管理域名加 IP 白名单。

生产部署前请重点确认：

- `JWT_SECRET` 已替换为随机强密钥；
- `DEFAULT_ADMIN_PASSWORD` 已修改；
- MySQL 密码、OSS 凭证、Twilio 凭证通过环境变量注入，不写入仓库；
- `CORS_ORIGINS` 限制为实际前端域名；
- 对外域名与 HTTPS 证书配置完成。

## 相关文档

- [模板与风格契约](docs/template-style-contract.md)
- [后端接口文档](http://localhost:8000/docs)（本地启动后）
- [Nginx 配置](deploy/nginx.conf)

## License

MIT
