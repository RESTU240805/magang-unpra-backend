package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type visitorEntry struct {
	requests []time.Time
}

type RateLimiter struct {
	visitors   map[string]*visitorEntry
	mu         sync.Mutex
	maxReqs    int
	window     time.Duration
	cleanupInt time.Duration
}

func NewRateLimiter(maxRequests int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors:   make(map[string]*visitorEntry),
		maxReqs:    maxRequests,
		window:     window,
		cleanupInt: 5 * time.Minute,
	}
	go rl.cleanup()
	return rl
}

func (rl *RateLimiter) cleanup() {
	for {
		time.Sleep(rl.cleanupInt)
		rl.mu.Lock()
		cutoff := time.Now().Add(-rl.window)
		for ip, v := range rl.visitors {
			var active []time.Time
			for _, t := range v.requests {
				if t.After(cutoff) {
					active = append(active, t)
				}
			}
			if len(active) == 0 {
				delete(rl.visitors, ip)
			} else {
				v.requests = active
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()
		cutoff := now.Add(-rl.window)

		rl.mu.Lock()
		v, exists := rl.visitors[ip]
		if !exists {
			v = &visitorEntry{}
			rl.visitors[ip] = v
		}

		var active []time.Time
		for _, t := range v.requests {
			if t.After(cutoff) {
				active = append(active, t)
			}
		}
		v.requests = active

		if len(active) >= rl.maxReqs {
			rl.mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			return
		}

		v.requests = append(v.requests, now)
		rl.mu.Unlock()

		c.Next()
	}
}

func RateLimit(maxRequests int, window time.Duration) gin.HandlerFunc {
	return NewRateLimiter(maxRequests, window).Middleware()
}
