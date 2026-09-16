package execution

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu           sync.Mutex
	rate         time.Duration
	lastExecuted time.Time
}

func NewRateLimiter(interval time.Duration) *RateLimiter {
	return &RateLimiter{
		rate: interval,
	}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.Sub(r.lastExecuted) < r.rate {
		return false
	}

	r.lastExecuted = now
	return true
}
