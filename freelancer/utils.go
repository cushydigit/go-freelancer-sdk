package freelancer

import (
	"net/url"
	"strconv"
)

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
