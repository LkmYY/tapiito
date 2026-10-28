// Command devgateway 是本地开发用的合并入口，模拟生产环境 nginx 的分流职责。
//
// 生产上公开域名只反代 :8080、管理域名只反代 :8081；本地前端只有一个 base URL，
// 因此需要一个按「路径 + 方法」分流的网关，否则前端无法同时访问两个入口。
//
// 分流规则（与 nginx 等价）：
//   /api/public/*            → 公开 :8080（只读 + 公开写）
//   POST /api/orders         → 公开 :8080（下单是公开能力）
//   POST /api/feedback       → 公开 :8080（反馈是公开能力）
//   GET  /api/feedback/adopted → 公开 :8080（已采纳意见公开展示）
//   /uploads/*               → 公开 :8080（静态资源）
//   其余 /api/*              → 管理 :8081（需要 JWT）
//
// 仅用于本地联调，不要部署到生产。
package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	listen := flag.String("listen", ":8090", "网关监听地址")
	public := flag.String("public", "http://127.0.0.1:8080", "公开 API")
	admin := flag.String("admin", "http://127.0.0.1:8081", "管理 API")
	flag.Parse()

	pubURL, err := url.Parse(*public)
	if err != nil {
		log.Fatalf("bad public url: %v", err)
	}
	admURL, err := url.Parse(*admin)
	if err != nil {
		log.Fatalf("bad admin url: %v", err)
	}

	pubProxy := httputil.NewSingleHostReverseProxy(pubURL)
	admProxy := httputil.NewSingleHostReverseProxy(admURL)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target := admProxy
		label := "admin"
		switch {
		case strings.HasPrefix(r.URL.Path, "/api/public/"):
			target, label = pubProxy, "public"
		case strings.HasPrefix(r.URL.Path, "/uploads/"):
			target, label = pubProxy, "public"
		case r.Method == http.MethodPost && r.URL.Path == "/api/orders":
			target, label = pubProxy, "public"
		case r.Method == http.MethodPost && r.URL.Path == "/api/feedback":
			target, label = pubProxy, "public"
		case r.Method == http.MethodGet && r.URL.Path == "/api/feedback/adopted":
			target, label = pubProxy, "public"
		}
		w.Header().Set("X-Dev-Gateway", label)
		target.ServeHTTP(w, r)
	})

	log.Printf("dev gateway on %s  public=%s  admin=%s", *listen, *public, *admin)
	if err := http.ListenAndServe(*listen, handler); err != nil {
		log.Fatal(err)
	}
}
