// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package middleware

import (
	"net/http"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserRateLimitMiddleware struct {
	store *limiterStore
}

func NewUserRateLimitMiddleware(perMinute int) *UserRateLimitMiddleware {
	return &UserRateLimitMiddleware{store: newLimiterStore(perMinute)}
}

func (m *UserRateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := clientIP(r)
		if userID := r.Context().Value("userId"); userID != nil {
			key += ":" + toString(userID)
		}
		if !m.store.allow(key) {
			logx.WithContext(r.Context()).WithFields(logx.Field("method", r.Method), logx.Field("path", r.URL.Path), logx.Field("client_ip", clientIP(r)), logx.Field("limit_per_minute", m.store.perMinute)).Error("request rate limit exceeded")
			reject(w)
			return
		}
		next(w, r)
	}
}
