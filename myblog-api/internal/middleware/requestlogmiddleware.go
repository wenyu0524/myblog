package middleware

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// RequestLog records every API request with the caller and outcome.
func RequestLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := logx.WithFields(r.Context(),
			logx.Field("method", r.Method),
			logx.Field("path", r.URL.Path),
			logx.Field("client_ip", clientIP(r)),
		)
		if uid := ctx.Value("userId"); uid != nil {
			ctx = logx.WithFields(ctx, logx.Field("user_id", toString(uid)))
		}
		start := time.Now()
		logx.WithContext(ctx).Info("request started")
		next(w, r.WithContext(ctx))
		logx.WithContext(ctx).Infof("request completed duration_ms=%d", time.Since(start).Milliseconds())
	}
}
