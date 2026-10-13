package query

import (
	"strconv"
	"testing"
	"time"

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

	values, err := Values(v)

	assert.Equal(t, "golang", values.Get("query"))
	assert.Equal(t, "10", values.Get("limit"))
	assert.Equal(t, "true", values.Get("active"))
	assert.Equal(t, "100.5", values.Get("amount"))
	assert.Nil(t, err)
}

func TestValues_NilPointersAreIgnored(t *testing.T) {
	type opts struct {
		Query *string `url:"query"`
		Limit *int    `url:"limit"`
	}

	values, err := Values(opts{})
	assert.Empty(t, values)
	assert.Equal(t, 0, len(values))
	assert.Nil(t, err)
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

	values, err := Values(v)

	assert.ElementsMatch(t, []string{"1", "2"}, values["jobs[]"])
	assert.ElementsMatch(t, []string{"US", "DE"}, values["countries[]"])
	assert.Nil(t, err)
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

	values, err := Values(v)

	assert.Empty(t, values)
	assert.Nil(t, err)
}

func TestValues_PointerToStruct(t *testing.T) {
	type opts struct {
		Query string `url:"query"`
	}

	values, err := Values(&opts{Query: "test"})

	assert.Equal(t, "test", values.Get("query"))
	assert.Nil(t, err)
}

func TestValues_NilPointerInput(t *testing.T) {
	var opts *struct {
		Query string `url:"query"`
	}

	values, err := Values(opts)

	assert.NotNil(t, values)
	assert.Empty(t, values)
	assert.Nil(t, err)
}

func TestValues_IgnoreMissingOrDashTags(t *testing.T) {
	type opts struct {
		Query  string `url:"query"`
		Ignore string
		Skip   string `url:"-"`
	}

	values, err := Values(opts{
		Query:  "ok",
		Ignore: "no",
		Skip:   "maybe",
	})

	assert.Equal(t, "ok", values.Get("query"))
	assert.Equal(t, 1, len(values))
	assert.Nil(t, err)
}

func TestValues_UnsupportedTypesDoNotPanic(t *testing.T) {
	type opts struct {
		Map map[string]string `url:"map"`
	}

	assert.NotPanics(t, func() {
		vs, err := Values(opts{
			Map: map[string]string{"a": "b"},
		})
		assert.Nil(t, vs)
		assert.NotNil(t, err)
	})
}

func TestValues_EncodedURL(t *testing.T) {
	type opts struct {
		Jobs []int `url:"jobs[]"`
	}

	values, err := Values(opts{Jobs: []int{1, 2}})
	encoded := values.Encode()

	assert.Contains(t, encoded, "jobs%5B%5D=1")
	assert.Contains(t, encoded, "jobs%5B%5D=2")
	assert.Nil(t, err)
}

func TestValues_NonStructInput_Primitive(t *testing.T) {
	values, err := Values(1)
	assert.NotNil(t, values)
	assert.Empty(t, values)
	assert.Error(t, err)
}

func TestValues_NonStructInput_Slice(t *testing.T) {
	values, err := Values([]int{1, 2, 4})
	assert.NotNil(t, values)
	assert.Empty(t, values)
	assert.Error(t, err)
}

func TestValues_TimeFields(t *testing.T) {
	type opts struct {
		FromTime *time.Time `url:"from_time"`
		ToTime   *time.Time `url:"to_time"`
	}

	from := time.Date(2026, 9, 1, 12, 0, 0, 0, time.UTC)
	to := time.Date(2026, 9, 7, 12, 0, 0, 0, time.UTC)

	values, err := Values(opts{
		FromTime: &from,
		ToTime:   &to,
	})

	assert.NoError(t, err)
	assert.Equal(t, strconv.FormatInt(from.Unix(), 10), values.Get("from_time"))
	assert.Equal(t, strconv.FormatInt(to.Unix(), 10), values.Get("to_time"))
}

func TestValues_NilTimePointersAreIgnored(t *testing.T) {
	type opts struct {
		FromTime *time.Time `url:"from_time"`
		ToTime   *time.Time `url:"to_time"`
	}

	values, err := Values(opts{})

	assert.NoError(t, err)
	assert.Empty(t, values)
}
