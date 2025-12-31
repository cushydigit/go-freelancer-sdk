package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues_BasicTypesAndPointers(t *testing.T) {
	type opts struct {
		Query  *string  `url:"query"`
		Limit  *int     `url:"limit"`
		Active *bool    `url:"active"`
		Amount *float64 `url:"amount"`
	}

	q := "golang"
	limit := 10
	active := true
	amount := 100.5

	v := opts{
		Query:  &q,
		Limit:  &limit,
		Active: &active,
		Amount: &amount,
	}

	values := Values(v)

	assert.Equal(t, "golang", values.Get("query"))
	assert.Equal(t, "10", values.Get("limit"))
	assert.Equal(t, "true", values.Get("active"))
	assert.Equal(t, "100.5", values.Get("amount"))
}

func TestValues_NilPointersAreIgnored(t *testing.T) {
	type opts struct {
		Query *string `url:"query"`
		Limit *int    `url:"limit"`
	}

	values := Values(opts{})
	assert.Empty(t, values)
	assert.Equal(t, 0, len(values))
}

func TestValues_SliceFields(t *testing.T) {
	type opts struct {
		Jobs      []int64  `url:"jobs[]"`
		Countries []string `url:"countries[]"`
	}

	v := opts{
		Jobs:      []int64{1, 2},
		Countries: []string{"US", "DE"},
	}

	values := Values(v)

	assert.ElementsMatch(t, []string{"1", "2"}, values["jobs[]"])
	assert.ElementsMatch(t, []string{"US", "DE"}, values["countries[]"])
}

func TestValues_ZeroFloatIsIgnored(t *testing.T) {
	type opts struct {
		MinPrice *float64 `url:"min_price"`
		MaxPrice *float64 `url:"max_price"`
	}

	zero := 0.0
	v := opts{
		MinPrice: &zero,
		MaxPrice: &zero,
	}

	values := Values(v)

	assert.Empty(t, values)
}

func TestValues_PointerToStruct(t *testing.T) {
	type opts struct {
		Query string `url:"query"`
	}

	values := Values(&opts{Query: "test"})

	assert.Equal(t, "test", values.Get("query"))
}

func TestValues_NilPointerInput(t *testing.T) {
	var opts *struct {
		Query string `url:"query"`
	}

	values := Values(opts)

	assert.NotNil(t, values)
	assert.Empty(t, values)
}

func TestValues_IgnoreMissingOrDashTags(t *testing.T) {
	type opts struct {
		Query  string `url:"query"`
		Ignore string
		Skip   string `url:"-"`
	}

	values := Values(opts{
		Query:  "ok",
		Ignore: "no",
		Skip:   "maybe",
	})

	assert.Equal(t, "ok", values.Get("query"))
	assert.Equal(t, 1, len(values))
}

func TestValues_UnsupportedTypesDoNotPanic(t *testing.T) {
	type opts struct {
		Map map[string]string `url:"map"`
	}
	assert.NotPanics(t, func() {
		_ = Values(opts{
			Map: map[string]string{"a": "b"},
		})
	})
}

func TestValues_EncodedURL(t *testing.T) {
	type opts struct {
		Jobs []int `url:"jobs[]"`
	}

	values := Values(opts{Jobs: []int{1, 2}})
	encoded := values.Encode()

	assert.Contains(t, encoded, "jobs%5B%5D=1")
	assert.Contains(t, encoded, "jobs%5B%5D=2")
}

func TestValues_NonStructInput_Primitive(t *testing.T) {
	values := Values(1)
	assert.NotNil(t, values)
	assert.Empty(t, values)
}

func TestValues_NonStructInput_Slice(t *testing.T) {
	values := Values([]int{1, 2, 4})
	assert.NotNil(t, values)
	assert.Empty(t, values)
}
