package httpapi

import (
	"encoding/json"
	"html"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ShareHandler 为需要「分享卡片」的路由动态注入 meta。
//
// 背景：前端是 SPA，所有路由共用同一个 index.html，
// 里面的 title/description 写死为首页文案。
// 微信、QQ、钉钉等抓取分享卡片时不执行 JS，
// 只能读到 HTML 里的静态 meta——所以必须在服务端替换。
type ShareHandler struct {
	indexPath string
	metaPath  string

	mu       sync.RWMutex
	tmpl     string    // index.html 原文
	tmplMod  time.Time // 文件修改时间，变化则重新读取
	typeName map[string]string
}

// NewShareHandler 创建分享卡片处理器。
func NewShareHandler(indexPath, metaPath string) *ShareHandler {
	h := &ShareHandler{indexPath: indexPath, metaPath: metaPath}
	h.loadTypeNames()
	return h
}

func (h *ShareHandler) loadTypeNames() {
	h.typeName = map[string]string{}
	raw, err := os.ReadFile(h.metaPath)
	if err != nil {
		return
	}
	var meta struct {
		Types map[string]string `json:"types"`
	}
	if json.Unmarshal(raw, &meta) == nil {
		h.typeName = meta.Types
	}
}

// template 读取并缓存 index.html；文件更新后自动重载。
func (h *ShareHandler) template() string {
	st, err := os.Stat(h.indexPath)
	if err != nil {
		return ""
	}

	h.mu.RLock()
	if h.tmpl != "" && h.tmplMod.Equal(st.ModTime()) {
		tmpl := h.tmpl
		h.mu.RUnlock()
		return tmpl
	}
	h.mu.RUnlock()

	raw, err := os.ReadFile(h.indexPath)
	if err != nil {
		return ""
	}

	h.mu.Lock()
	h.tmpl = string(raw)
	h.tmplMod = st.ModTime()
	h.mu.Unlock()

	return h.tmpl
}

// TypeLabel 返回模板类型的中文名。
func (h *ShareHandler) TypeLabel(key string) string {
	if v, ok := h.typeName[key]; ok && v != "" {
		return v
	}
	return ""
}

// prettify 把 gacha_ur 这类 slug 变成 Gacha UR，用于风格名兜底。
func prettify(slug string) string {
	if slug == "" {
		return ""
	}
	parts := strings.Split(slug, "_")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p == "" {
			continue
		}
		if len(p) <= 3 { // ur / vr / 3d 这类缩写保持大写
			out = append(out, strings.ToUpper(p))
		} else {
			out = append(out, strings.ToUpper(p[:1])+p[1:])
		}
	}
	return strings.Join(out, " ")
}

// Preview 处理 /preview 的分享抓取。
func (h *ShareHandler) Preview(c *gin.Context) {
	tmpl := h.template()
	if tmpl == "" {
		c.String(http.StatusNotFound, "index template not found")
		return
	}

	typ := c.Query("type")
	style := c.Query("style")

	typeLabel := h.TypeLabel(typ)
	styleLabel := prettify(style)

	parts := make([]string, 0, 2)
	if typeLabel != "" {
		parts = append(parts, typeLabel)
	} else if typ != "" {
		parts = append(parts, prettify(typ))
	}
	if styleLabel != "" {
		parts = append(parts, styleLabel)
	}

	title := "Tapiito 模板预览"
	desc := "碰一下，内容就跳出来。9 种标签类型 × 30 种设计风格，内容随时可改。"
	if len(parts) > 0 {
		title = strings.Join(parts, " · ") + " ｜ Tapiito"
		desc = "Tapiito 模板预览：" + strings.Join(parts, " · ") +
			"。碰一下标签即可展示，9 种类型 × 30 种风格，内容随时可改。"
	}

	// 缩略图用现成的 jpg（微信对 webp 支持不稳）
	image := "https://tapiito.com/docs/assets/logo.png"
	if typ != "" && style != "" {
		image = "https://tapiito.com/share/" + typ + "__" + style + ".jpg"
	} else if typ != "" {
		image = "https://tapiito.com/share/" + typ + "__.jpg"
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(replaceMeta(tmpl, title, desc, image, c.Request.URL.String())))
}

// replaceMeta 替换 HTML 里的 title / description / og 标签。
func replaceMeta(tmpl, title, desc, image, path string) string {
	safeTitle := html.EscapeString(title)
	safeDesc := html.EscapeString(desc)

	out := tmpl

	// <title>
	if i := strings.Index(out, "<title>"); i >= 0 {
		if j := strings.Index(out[i:], "</title>"); j >= 0 {
			out = out[:i] + "<title>" + safeTitle + "</title>" + out[i+j+len("</title>"):]
		}
	}

	// description
	out = replaceAttr(out, `name="description"`, "content", safeDesc)
	// og 系列
	out = replaceAttr(out, `property="og:title"`, "content", safeTitle)
	out = replaceAttr(out, `property="og:description"`, "content", safeDesc)
	out = replaceAttr(out, `property="og:image"`, "content", image)
	out = replaceAttr(out, `property="og:url"`, "content", "https://tapiito.com"+path)

	// 没有 og:image 就补一个（很多页面模板里没写）
	if !strings.Contains(out, `property="og:image"`) {
		inject := `<meta property="og:image" content="` + image + `">` +
			`<meta property="og:title" content="` + safeTitle + `">` +
			`<meta property="og:description" content="` + safeDesc + `">`
		out = strings.Replace(out, "</head>", inject+"</head>", 1)
	}

	return out
}

// replaceAttr 把形如 <meta name="description" content="..."> 的 content 换掉。
func replaceAttr(s, selector, attr, value string) string {
	idx := strings.Index(s, selector)
	if idx < 0 {
		return s
	}
	// 只在同一个标签内查找 content
	tagStart := strings.LastIndex(s[:idx], "<")
	tagEnd := strings.Index(s[idx:], ">")
	if tagStart < 0 || tagEnd < 0 {
		return s
	}
	tagEnd += idx
	tag := s[tagStart:tagEnd]

	if i := strings.Index(tag, attr+`="`); i >= 0 {
		j := strings.Index(tag[i+len(attr)+2:], `"`)
		if j >= 0 {
			newTag := tag[:i] + attr + `="` + value + `"` + tag[i+len(attr)+2+j+1:]
			return s[:tagStart] + newTag + s[tagEnd:]
		}
	}
	// 标签里没有该属性则补上
	newTag := tag + " " + attr + `="` + value + `"`
	return s[:tagStart] + newTag + s[tagEnd:]
}
