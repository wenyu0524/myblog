package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"myblog-api/response"

	"golang.org/x/time/rate"
)

type limiterStore struct {
	mu sync.Mutex
	// ponytail: process-local, unbounded client map; use an expiring shared store for many clients or replicas.
	limiters  map[string]*rate.Limiter
	perMinute int
}

func newLimiterStore(perMinute int) *limiterStore {
	return &limiterStore{
		limiters:  make(map[string]*rate.Limiter),
		perMinute: perMinute,
	}
}

func (s *limiterStore) allow(key string) bool {
	if s.perMinute <= 0 {
		return true
	}
	s.mu.Lock()
	limiter, ok := s.limiters[key]
	if !ok {
		limiter = rate.NewLimiter(rate.Every(time.Minute/time.Duration(s.perMinute)), s.perMinute)
		s.limiters[key] = limiter
	}
	s.mu.Unlock()
	return limiter.Allow()
}

func reject(w http.ResponseWriter) {
	response.Error(w, http.StatusTooManyRequests, "请求过于频繁")
}

func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}
	return r.RemoteAddr
}
