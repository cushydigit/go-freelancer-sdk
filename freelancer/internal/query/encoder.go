package query

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
)

func Values(v any) url.Values {
	q := url.Values{}
	val := reflect.ValueOf(v)

	// handler pointers
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return q // not covered
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return q // not covered
	}
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("url")

		if tag == "" || tag == "-" {
			continue
		}

		fVal := val.Field(i)
		writeField(q, tag, fVal)

	}
	return q
}

func writeField(q url.Values, key string, v reflect.Value) {
	// if it's a pointer, check for nil and dereference
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.String:
		q.Set(key, v.String())
	case reflect.Bool:
		q.Set(key, strconv.FormatBool(v.Bool()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		q.Set(key, strconv.FormatInt(v.Int(), 10))
	case reflect.Float64:
		if v.Float() == 0 {
			return
		}
		q.Set(key, strconv.FormatFloat(v.Float(), 'f', -1, 64))
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			q.Add(key, fmt.Sprintf("%v", v.Index(i).Interface()))
		}
	default:
		log.Printf("warning unhandled typed")
	}
}
