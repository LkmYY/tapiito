package httpapi

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lkmyy/tapiito-go/internal/auth"
	"github.com/lkmyy/tapiito-go/internal/config"
	"github.com/lkmyy/tapiito-go/internal/ratelimit"
)

// Deps 是构造路由所需的全部依赖。
type Deps struct {
	Cfg       *config.Config
	Log       *slog.Logger
	Limiter   *ratelimit.Limiter
	Blacklist *ratelimit.Blacklist
	Issuer    *auth.Issuer
	Public    *PublicHandler
	Admin     *AdminHandler
	Order     *OrderHandler
	Feedback  *FeedbackHandler
	Stats     *StatsHandler
	Health    []func() error
}

// NewPublicRouter 构建公开入口：
// 只注册公开读、公开写（留言/照片/短信/下单/反馈）与展厅/落地页，绝不注册管理能力。
func NewPublicRouter(d Deps) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), RequestID(), RealIP(), SecurityHeaders(),
		CORSMiddleware(d.Cfg.CORSOrigins), BlacklistGuard(d.Blacklist, d.Log))

	r.GET("/healthz", Health(d.Health...))

	api := r.Group("/api")
	api.Use(GlobalQPS(d.Limiter, d.Cfg.Limit.GlobalQPS, d.Log))
	{
		// 公开读
		read := RateLimit(d.Limiter, d.Blacklist, "rl:public",
			d.Cfg.Limit.PublicReadRPS, d.Cfg.Limit.PublicReadBurst, d.Log)
		pub := api.Group("/public")
		pub.GET("/gallery", read, d.Public.Gallery)
		pub.GET("/nfc/:nfc_id/page", read, d.Public.Page)
		pub.GET("/nfc/:nfc_id/photos", read, d.Public.Photos)

		// 公开写
		write := RateLimit(d.Limiter, d.Blacklist, "rl:public:write",
			d.Cfg.Limit.PublicWriteRPS, d.Cfg.Limit.PublicWriteBurst, d.Log)
		pub.POST("/nfc/:nfc_id/note", write, d.Public.Note)
		pub.POST("/nfc/:nfc_id/photo", write, d.Public.Photo)
		pub.POST("/nfc/:nfc_id/submit", write, d.Public.Submit)
		// 暗号换 ticket：限流比普通写更严，防爆破主要靠下方 Redis 计数
		pub.POST("/nfc/:nfc_id/unlock", write, d.Public.Unlock)

		// 下单与反馈：路径不带 /public 前缀（与 FastAPI 一致）
		api.POST("/orders", write, d.Order.Create)
		api.POST("/feedback", write, d.Feedback.Create)
		api.GET("/feedback/adopted", read, d.Feedback.Adopted)
	}

	return r
}

// NewAdminRouter 构建管理入口：独立端口，除登录外全部需要 JWT。
func NewAdminRouter(d Deps) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), RequestID(), RealIP(), SecurityHeaders(),
		BlacklistGuard(d.Blacklist, d.Log))

	r.GET("/healthz", Health(d.Health...))

	api := r.Group("/api")
	// 登录单独限流，防止撞库
	api.POST("/auth/loginAdmin",
		RateLimit(d.Limiter, d.Blacklist, "rl:login", d.Cfg.Limit.LoginPerMinute,
			d.Cfg.Limit.LoginPerMinute*2, d.Log),
		d.Admin.Login)

	authed := api.Group("")
	authed.Use(RequireAdmin(d.Issuer, d.Log),
		RateLimit(d.Limiter, d.Blacklist, "rl:admin",
			d.Cfg.Limit.AdminRPS, d.Cfg.Limit.AdminBurst, d.Log))
	{
		// 标签
		authed.GET("/nfc", d.Admin.ListNfc)
		authed.POST("/nfc", d.Admin.CreateNfc)
		authed.POST("/nfc/batch", d.Admin.BatchNfc)
		authed.GET("/nfc/creators", d.Admin.Creators)
		authed.GET("/nfc/:id", d.Admin.GetNfc)
		authed.PUT("/nfc/:id", d.Admin.UpdateNfc)
		authed.DELETE("/nfc/:id", d.Admin.DeleteNfc)

		// 订单
		authed.GET("/orders", d.Order.List)
		authed.GET("/orders/:order_id", d.Order.Get)
		authed.POST("/orders/:order_id/convert", d.Order.Convert)
		authed.DELETE("/orders/:order_id", d.Order.Delete)

		// 反馈管理
		authed.GET("/feedback/admin", d.Feedback.AdminList)
		authed.PATCH("/feedback/admin/:feedback_id", d.Feedback.Adopt)

		// 统计
		authed.GET("/stats/overview", d.Stats.Overview)
		authed.GET("/stats/nfc/:nfc_id", d.Stats.ByNfc)
		authed.GET("/stats/export", d.Stats.Export)
	}

	return r
}

// NotFound 是统一的 404 响应，保持 FastAPI 的 detail 结构。
func NotFound(c *gin.Context) { c.JSON(http.StatusNotFound, gin.H{"detail": "Not Found"}) }
