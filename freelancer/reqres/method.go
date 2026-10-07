package reqres

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
