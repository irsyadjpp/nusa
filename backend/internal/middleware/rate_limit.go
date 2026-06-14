package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimiter implements a simple in-memory rate limiter
type RateLimiter struct {
	visitors map[string]*visitor
	mu       sync.RWMutex
	rate     int           // requests per minute
	burst    int           // burst size
}

type visitor struct {
	requests  []time.Time
	lastReset time.Time
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(rate, burst int) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		burst:    burst,
	}
	
	// Start cleanup goroutine
	go rl.cleanup()
	
	return rl
}

// cleanup removes old visitor entries
func (rl *RateLimiter) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	
	for range ticker.C {
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastReset) > 5*time.Minute {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// Allow checks if a request from the given IP is allowed
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	
	now := time.Now()
	v, exists := rl.visitors[ip]
	
	if !exists {
		rl.visitors[ip] = &visitor{
			requests:  []time.Time{now},
			lastReset: now,
		}
		return true
	}
	
	// Reset if minute has passed
	if now.Sub(v.lastReset) >= time.Minute {
		v.requests = []time.Time{now}
		v.lastReset = now
		return true
	}
	
	// Check burst limit
	if len(v.requests) >= rl.burst {
		return false
	}
	
	// Check rate limit
	if len(v.requests) >= rl.rate {
		return false
	}
	
	v.requests = append(v.requests, now)
	return true
}

// GetRateLimit returns the rate limit headers
func (rl *RateLimiter) GetRateLimit(ip string) (remaining, reset int) {
	rl.mu.RLock()
	defer rl.mu.RUnlock()
	
	v, exists := rl.visitors[ip]
	if !exists {
		return rl.rate, int(time.Now().Add(time.Minute).Unix())
	}
	
	remaining = rl.rate - len(v.requests)
	if remaining < 0 {
		remaining = 0
	}
	
	reset = int(v.lastReset.Add(time.Minute).Unix())
	return
}

// Global rate limiter instance
var globalRateLimiter *RateLimiter

// InitRateLimiter initializes the global rate limiter
func InitRateLimiter(rate, burst int) {
	globalRateLimiter = NewRateLimiter(rate, burst)
}

// RateLimit middleware applies rate limiting based on IP address
func RateLimit() gin.HandlerFunc {
	if globalRateLimiter == nil {
		// Default rate limit if not initialized
		InitRateLimiter(100, 20)
	}
	
	return func(c *gin.Context) {
		ip := c.ClientIP()
		
		if !globalRateLimiter.Allow(ip) {
			remaining, reset := globalRateLimiter.GetRateLimit(ip)
			
			c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", globalRateLimiter.rate))
			c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
			c.Header("X-RateLimit-Reset", fmt.Sprintf("%d", reset))
			c.Header("Retry-After", "60")
			
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Please try again later.",
			})
			c.Abort()
			return
		}
		
		remaining, reset := globalRateLimiter.GetRateLimit(ip)
		c.Header("X-RateLimit-Limit", fmt.Sprintf("%d", globalRateLimiter.rate))
		c.Header("X-RateLimit-Remaining", fmt.Sprintf("%d", remaining))
		c.Header("X-RateLimit-Reset", fmt.Sprintf("%d", reset))
		
		c.Next()
	}
}
