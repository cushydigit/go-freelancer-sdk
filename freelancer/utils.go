package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

func execute[T any](ctx context.Context, c *Client, method, path string, query url.Values, body any) (T, error) {
	var result T

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return result, err
		}
		bodyReader = bytes.NewReader(b)
	}
	data, err := c.do(ctx, method, path, query, bodyReader)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("decode error: %w", err)
	}

	return result, nil
}

// Bool returns a pointer to the provided bool.
func Bool(v bool) *bool { return &v }

// String returns a pointer to the provided string.
func String(v string) *string { return &v }

// Int returns a pointer to the provided int.
func Int(v int) *int { return &v }

// Int64 returns a pointer to the provided int64.
func Int64(v int64) *int64 { return &v }

// Float64 returns a pointer to the provided float64.
func Float64(v float64) *float64 { return &v }

// Enum returns a pointer to the provided enum.
func Enum[T ~string](v T) *T { return &v }
