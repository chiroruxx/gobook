package params

import (
	"fmt"
	"net/http"
	"net/mail"
	"reflect"
	"strconv"
	"strings"
)

type validationRule func(value string) bool

var validateRules = map[string]validationRule{
	"mail": func(value string) bool {
		_, err := mail.ParseAddress(value)
		return err == nil
	},
}

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	type fieldMapValue struct {
		val       reflect.Value
		validator validationRule
	}

	// 名前と値のマッピング
	fieldMap := make(map[string]fieldMapValue)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag

		httpName := tag.Get("http")
		if httpName == "" {
			httpName = strings.ToLower(fieldInfo.Name)
		}

		var validator validationRule
		validateName := tag.Get("validate")
		if validateName != "" {
			var ok bool
			if validator, ok = validateRules[validateName]; !ok {
				return fmt.Errorf("invalid validate rule: %s", validateName)
			}
		}
		fieldMap[httpName] = fieldMapValue{
			v.Field(i),
			validator,
		}
	}

	// 値の設定
	for name, values := range req.Form {
		f := fieldMap[name].val
		if !f.IsValid() {
			continue
		}
		validator := fieldMap[name].validator
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value, validator); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value, validator); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}

	return nil
}

func populate(v reflect.Value, value string, validator validationRule) error {
	if validator != nil && !validator(value) {
		return fmt.Errorf("invalid value: %s", value)
	}

	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
