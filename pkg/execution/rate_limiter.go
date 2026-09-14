package execution

import (
	"context"
	"golang.org/x/time/rate"
	"time"
)

type APIThrottle struct {
	Limiter *rate.Limiter
}

func NewAPIThrottle(r rate.Limit, b int) *APIThrottle {
	return &APIThrottle{Limiter: rate.NewLimiter(r, b)}
}

func (at *APIThrottle) Wait(ctx context.Context) error {
	return at.Limiter.Wait(ctx)
}
