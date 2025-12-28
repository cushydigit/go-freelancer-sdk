package freelancer

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.RWMutex
	limit     int
	remaining int
	reset     time.Time
}

func newRateLimiter() *RateLimiter {
	return &RateLimiter{
		limit:     50,
		remaining: 50,
		reset:     time.Now(),
	}
}

func (r *RateLimiter) updateFromHeader(header http.Header) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Freelancer typical headers
	if rem := header.Get("X-RateLimit-Remaining"); rem != "" {
		if v, err := strconv.Atoi(rem); err == nil {
			r.remaining = v
		}
	}

	// Some APIs provide a 'Reset' timestamp (seconds since epoch)
	if res := header.Get("X-RateLimit-Reset"); res != "" {
		if v, err := strconv.ParseInt(res, 10, 64); err == nil {
			r.reset = time.Unix(v, 0)
		}
	}

}

func (r *RateLimiter) wait(ctx context.Context) error {
	r.mu.RLock()
	remaining := r.remaining
	reset := r.reset
	r.mu.RUnlock()

	if remaining > 0 {
		r.mu.Lock()
		r.remaining--
		r.mu.Unlock()
		return nil
	}

	waiteDuration := time.Until(reset)
	if waiteDuration <= 0 {
		return nil
	}

	time := time.NewTimer(waiteDuration)
	defer time.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.C:
		return nil
	}
}
