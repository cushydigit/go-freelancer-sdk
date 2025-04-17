package freelancer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

type params map[string]interface{}

type request struct {
	method   string
	endpoint string
	query    url.Values
	header   http.Header
	body     io.Reader
	fullURL  string
}

func (r *request) addParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Add(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		v, err := json.Marshal(value)
		if err == nil {
			value = string(v)
		}
	}

	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (r *request) validate() error {
	if r.method == "" {
		panic("method cannot be empty")
	}
	if r.endpoint == "" {
		panic("endpoint cannot be empty")
	}
	if r.query == nil {
		r.query = url.Values{}
	}
	return nil
}
