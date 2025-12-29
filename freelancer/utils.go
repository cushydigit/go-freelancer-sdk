package freelancer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
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

func addBool(q url.Values, key string, b *bool) {
	if b != nil {
		q.Set(key, strconv.FormatBool(*b))
	}
}

func addInt(q url.Values, key string, i *int) {
	if i != nil {
		q.Set(key, strconv.Itoa(*i))
	}
}

func addFloat(q url.Values, key string, f *float64) {
	if f != nil {
		// 'f' = no exponent
		// -1 = smallest necessary precision
		// 64 = float64
		q.Set(key, strconv.FormatFloat(*f, 'f', -1, 64))
	}
}

func addString(q url.Values, key string, s *string) {
	if s != nil {
		q.Set(key, *s)
	}
}

func addInt64(q url.Values, key string, i *int64) {
	if i != nil {
		q.Set(key, strconv.FormatInt(*i, 10))
	}
}

// StringTyped is a constraint for any type that is an underlying string
type StringTyped interface {
	~string
}

func addEnum[T StringTyped](q url.Values, key string, v *T) {
	if v != nil {
		q.Set(key, string(*v))
	}
}

func addInt64Slice(q url.Values, key string, values []int64) {
	for _, v := range values {
		q.Add(key, strconv.FormatInt(v, 10))
	}
}

func addStringSlice(q url.Values, key string, values []string) {
	for _, v := range values {
		q.Add(key, v)
	}
}

func addEnumSlice[T StringTyped](q url.Values, key string, values []T) {
	for _, v := range values {
		q.Add(key, string(v))
	}
}
