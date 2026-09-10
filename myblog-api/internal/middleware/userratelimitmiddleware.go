// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package middleware

import (
	"net/http"
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
			reject(w)
			return
		}
		next(w, r)
	}
}
