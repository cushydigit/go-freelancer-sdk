package query

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

var timeType = reflect.TypeOf(time.Time{})

func Values(v any) (url.Values, error) {
	q := url.Values{}
	val := reflect.ValueOf(v)

	// handler pointers
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return q, nil
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return q, fmt.Errorf("query options must be a struct, got %s", val.Kind())
	}
	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		key := field.Tag.Get("url")

		if key == "" || key == "-" {
			continue
		}

		if err := encodeValue(q, key, val.Field(i)); err != nil {
			return nil, fmt.Errorf("%s: %w", field.Name, err)
		}

	}
	return q, nil
}

func encodeValue(q url.Values, key string, v reflect.Value) error {
	// if it's a pointer, check for nil and dereference
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		return encodeValue(q, key, v.Elem())
	}

	// time.Time
	if v.Type() == timeType {
		t := v.Interface().(time.Time)
		q.Set(key, strconv.FormatInt(t.Unix(), 10))
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		q.Set(key, v.String())

	case reflect.Bool:
		q.Set(key, strconv.FormatBool(v.Bool()))

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		q.Set(key, strconv.FormatInt(v.Int(), 10))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		q.Set(key, strconv.FormatUint(v.Uint(), 10))

	case reflect.Float32, reflect.Float64:
		// ignore zero value for float
		if v.Float() == 0 {
			return nil
		}
		q.Set(key, strconv.FormatFloat(v.Float(), 'f', -1, v.Type().Bits()))

	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			value, err := scalarValue(v.Index(i))
			if err != nil {
				return err
			}
			q.Add(key, value)
		}
	default:
		return fmt.Errorf("unsupported type %s", v.Type())
	}

	return nil
}

func scalarValue(v reflect.Value) (string, error) {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return "", nil
		}
		return scalarValue(v.Elem())
	}

	if v.Type() == timeType {
		t := v.Interface().(time.Time)
		return strconv.FormatInt(t.Unix(), 10), nil
	}

	switch v.Kind() {
	case reflect.String:
		return v.String(), nil

	case reflect.Bool:
		return strconv.FormatBool(v.Bool()), nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10), nil

	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, v.Type().Bits()), nil

	default:
		return "", fmt.Errorf("unsupported type %s", v.Type())
	}
}
