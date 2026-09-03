package middleware

import (
	"net"
	"net/http"

	"github.com/zeromicro/go-zero/core/limit"
)

func LimitByIP(limiter *limit.PeriodLimit) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				http.Error(w, "无效的客户端地址", http.StatusBadRequest)
				return
			}

			state, err := limiter.TakeCtx(r.Context(), ip)
			if err != nil || state == limit.OverQuota {
				w.Header().Set("Retry-After", "60")
				http.Error(w, "请求过于频繁", http.StatusTooManyRequests)
				return
			}

			next(w, r)
		}
	}
}
