// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginRateLimitMiddleware struct {
	store *limiterStore
}

func NewLoginRateLimitMiddleware(perMinute int) *LoginRateLimitMiddleware {
	return &LoginRateLimitMiddleware{store: newLimiterStore(perMinute)}
}

func (m *LoginRateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewReader(body))
		var req struct {
			Username string `json:"username"`
		}
		_ = json.Unmarshal(body, &req)
		key := clientIP(r) + ":" + req.Username
		if !m.store.allow(key) {
			logx.WithContext(r.Context()).WithFields(logx.Field("method", r.Method), logx.Field("path", r.URL.Path), logx.Field("client_ip", clientIP(r)), logx.Field("username", req.Username), logx.Field("limit_per_minute", m.store.perMinute)).Error("login rate limit exceeded")
			reject(w)
			return
		}
		next(w, r)
	}
}

func toString(value any) string {
	if text, ok := value.(string); ok {
		return text
	}
	if number, ok := value.(float64); ok {
		return strconv.FormatInt(int64(number), 10)
	}
	return "unknown"
}
