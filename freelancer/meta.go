package freelancer

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ResponseMeta struct {
	StatusCode int
	RateLimit  RateLimitInfo
}

type RateLimitInfo struct {
	Remaining    int
	Limits       []RateLimitWindow
	RawRemaining string
	RawLimit     string
}

type RateLimitWindow struct {
	Limit  int
	Window time.Duration
}

func parseResponseMeta(resp *http.Response) *ResponseMeta {

	meta := &ResponseMeta{
		StatusCode: resp.StatusCode,
	}
	meta.RateLimit = parseRateLimit(resp.Header)

	return meta
}

func parseRateLimit(header http.Header) RateLimitInfo {
	info := RateLimitInfo{}

	if remaining := header.Get("RateLimit-Remaining"); remaining != "" {
		if value, err := strconv.Atoi(strings.TrimSpace(remaining)); err == nil {
			info.Remaining = value
		}
	}
	info.RawRemaining = header.Get("RateLimit-Remaining")

	if limits := header.Get("RateLimit-Limit"); limits != "" {
		info.Limits = parseRateLimitLimits(limits)
	}
	info.RawLimit = header.Get("RateLimit-Limit")

	return info
}

func parseRateLimitLimits(value string) []RateLimitWindow {
	parts := strings.Split(value, ",")

	limits := make([]RateLimitWindow, 0, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		fields := strings.Split(part, ";")

		if len(fields) == 0 {
			continue
		}

		limit, err := strconv.Atoi(strings.TrimSpace(fields[0]))
		if err != nil {
			continue
		}

		window := time.Duration(0)

		for _, field := range fields[1:] {
			field = strings.TrimSpace(field)

			key, value, ok := strings.Cut(field, "=")
			if !ok {
				continue
			}

			if key != "window" {
				continue
			}

			seconds, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil {
				continue
			}

			window = time.Duration(seconds) * time.Second
		}

		limits = append(limits, RateLimitWindow{
			Limit:  limit,
			Window: window,
		})
	}

	return limits
}
