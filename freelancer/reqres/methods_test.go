package reqres

import (
	"math"
	"strings"
	"testing"
)

func TestBool(t *testing.T) {
	v := true
	p := Bool(v)
	if *p != v {
		t.Errorf("expected %t, got %t", v, *p)
	}
}

func TestString(t *testing.T) {
	v := "hello"
	p := String(v)
	if *p != v {
		t.Errorf("expected '%s', got '%s'", v, *p)
	}
}

func TestInt(t *testing.T) {
	v := 42
	p := Int(v)
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestInt64(t *testing.T) {
	v := int64(42)
	p := Int64(v)
	if *p != v {
		t.Errorf("expected %d, got %d", v, *p)
	}
}

func TestFloat64(t *testing.T) {
	v := 3.14
	p := Float64(v)
	if math.Abs(*p-v) > 1e-6 {
		t.Errorf("expected %.2f, got %.2f", v, *p)
	}
}

func TestEnum(t *testing.T) {
	type Color string
	const (
		Red   Color = "red"
		Green Color = "green"
	)
	enum := Enum[Color]("blue")
	if !strings.Contains(string(*enum), "blue") {
		t.Errorf("expected 'blue', got '%s'", *enum)
	}
}
