package freelancer

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu           sync.Mutex
	limit        int
	remaining    int
	window       time.Duration
	lastCheck    time.Time
	sleepOnLimit bool
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limit:     50,
		remaining: 50,
		window:    60 * time.Second,
	}
}

func (r *RateLimiter) UpdateFromHeader(resp *http.Response) {
	r.mu.Lock()
	defer r.mu.Unlock()

	limitHeader := resp.Header.Get("RateLimit-Limit")
	log.Printf("Rate-Limit: %s", limitHeader)
	// if limitHeader != "" {
	// 	parts := strings.Split(limitHeader, ";")
	// 	if len(parts) > 1 {
	// 		mainPart := strings.Split(parts[0], ",")[0]
	// 		if l, err := strconv.Atoi(strings.TrimSpace(mainPart)); err == nil {
	// 			r.limit = l
	// 		}
	// 	}
	// }

	limitRemainingHeader := resp.Header.Get("RateLimit-Remaining")

	log.Printf("RateLimit-Remaining: %s", limitRemainingHeader)
}

func (r *RateLimiter) WaitIfNeeded() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.remaining <= 1 {
		sleepTime := r.window - time.Since(r.lastCheck)
		if sleepTime > 0 {
			time.Sleep(sleepTime)
			r.remaining = r.limit
			r.lastCheck = time.Now()
		}
	}
}
