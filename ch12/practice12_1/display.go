package practice12_1

import (
	"fmt"
	"reflect"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	fmt.Printf(formatComposite(path, v, true))
}

func formatComposite(path string, v reflect.Value, newLine bool) (res string) {
	switch v.Kind() {
	case reflect.Invalid:
		return fmt.Sprintf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			res += formatComposite(fmt.Sprintf("%s[%d]", path, i), v.Index(i), newLine)
		}
		return
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			res += formatComposite(fieldPath, v.Field(i), newLine)
		}
		return
	case reflect.Map:
		for _, key := range v.MapKeys() {
			res += formatComposite(fmt.Sprintf("%s[%s]", path, formatComposite(key.Type().String(), key, false)), v.MapIndex(key), newLine)
		}
		return
	case reflect.Ptr:
		if v.IsNil() {
			return fmt.Sprintf("%s = nil\n", path)
		} else {
			return formatComposite(fmt.Sprintf("(*%s)", path), v.Elem(), newLine)
		}
	case reflect.Interface:
		if v.IsNil() {
			return fmt.Sprintf("%s = nil\n", path)
		} else {
			res = fmt.Sprintf("%s.type = %s\n", path, v.Elem().Type())
			res += formatComposite(path+".value", v.Elem(), newLine)
			return
		}
	default:
		res = fmt.Sprintf("%s = %s", path, formatAtom(v))
		if newLine {
			res += "\n"
		}
		return res
	}
}
