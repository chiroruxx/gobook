package practice12_11

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func Pack(given *url.URL, ptr any) (*url.URL, error) {
	var m = map[string]reflect.Value{}
	v := reflect.ValueOf(ptr)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.IsValid() {
			continue
		}
		tag := getTagName(v, i)
		m[tag] = field
	}

	queries := given.Query()
	for key, value := range m {
		valStr, err := valueToString(value)
		if err != nil {
			return nil, err
		}
		queries.Add(key, valStr)
	}
	given.RawQuery = queries.Encode()

	return given, nil
}

func valueToString(value reflect.Value) (string, error) {
	switch value.Kind() {
	case reflect.String:
		return value.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatInt(value.Int(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'g', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(value.Bool()), nil
	case reflect.Ptr:
		return valueToString(value.Elem())
	default:
		return "", fmt.Errorf("unsupported type: %s", value.Type())
	}
}

func getTagName(v reflect.Value, i int) string {
	field := v.Type().Field(i)
	tag := field.Tag
	name := tag.Get("http")
	if name == "" {
		return strings.ToLower(field.Name)
	}

	return name
}
