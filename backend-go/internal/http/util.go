package httpapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


// marshalPayload 把任意 JSON 值固化为入库用的 []byte；nil 返回 nil。
func marshalPayload(v any) ([]byte, error) {
	if v == nil {
		return nil, nil
	}
	return json.Marshal(v)
}

// bindJSON 在解析前缓存原始请求体，供错误响应回填 Pydantic 风格的 input 字段。
func bindJSON(c *gin.Context, req any) error {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(body))
	c.Set("raw_body", body)
	return c.ShouldBindJSON(req)
}

// bindError 把 gin/validator 的错误转换成 FastAPI(Pydantic) 的 422 结构：
//
//	{"detail": [{"type": "missing", "loc": ["body", "password"], "msg": "Field required"}]}
//
// 前端按 detail 做统一提示，结构对齐可以避免替换后端时出现行为差异。
func bindError(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "invalid request body"})
		return
	}
	details := make([]gin.H, 0, len(ve))
	for _, fe := range ve {
		typ, msg := "value_error", "Invalid value"
		switch fe.Tag() {
		case "required":
			typ, msg = "missing", "Field required"
		case "max":
			typ, msg = "string_too_long", "String should have at most "+fe.Param()+" characters"
		case "min":
			typ, msg = "too_short", "String should have at least "+fe.Param()+" characters"
		case "oneof":
			typ, msg = "enum", "Input should be "+fe.Param()
		case "gte", "gt", "lte", "lt":
			typ, msg = "less_than_equal", "Input should be within range"
		}
		item := gin.H{
			"type": typ,
			"loc":  []string{"body", toSnake(fe.Field())},
			"msg":  msg,
		}
		if input := rawBodyMap(c); input != nil {
			item["input"] = input
		}
		details = append(details, item)
	}
	c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": details})
}

// toSnake 把 Go 字段名转成 JSON 风格的蛇形名（NfcID → nfc_id）。
func toSnake(s string) string {
	var b strings.Builder
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i > 0 {
				prev := rune(s[i-1])
				if prev >= 'a' && prev <= 'z' {
					b.WriteByte('_')
				} else if i+1 < len(s) {
					next := rune(s[i+1])
					if next >= 'a' && next <= 'z' {
						b.WriteByte('_')
					}
				}
			}
			b.WriteRune(r - 'A' + 'a')
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// rawBodyMap 把缓存的原始请求体解析成 map，用于回填 input 字段。
func rawBodyMap(c *gin.Context) map[string]any {
	v, ok := c.Get("raw_body")
	if !ok {
		return nil
	}
	raw, ok := v.([]byte)
	if !ok || len(raw) == 0 {
		return nil
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil
	}
	return m
}

// validationError 输出 Pydantic 风格的单项校验错误（422），
// 用于业务层校验（如字符串长度按字符计数），与 FastAPI 完全一致。
func validationError(c *gin.Context, field, typ, msg string, input any, ctx map[string]any) {
	item := gin.H{"type": typ, "loc": []string{"body", field}, "msg": msg}
	if input != nil {
		item["input"] = input
	}
	if ctx != nil {
		item["ctx"] = ctx
	}
	c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": []gin.H{item}})
}
