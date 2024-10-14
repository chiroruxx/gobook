package format2

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var indentSize = 2

func Any(value any) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr:
		if v.IsNil() {
			return "nil"
		}
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}

func formatComplex(v reflect.Value, indent int) string {
	switch v.Kind() {
	case reflect.Slice:
		res := "slice [\n"
		for i := 0; i < v.Len(); i++ {
			res += fmt.Sprintf("%s%d: %s\n", indentString(indent+1), i, formatComplex(v.Index(i), indent+1))
		}
		res += fmt.Sprintf("%s]", indentString(indent))
		return res
	case reflect.Array:
		res := "array [\n"
		for i := 0; i < v.Len(); i++ {
			res += fmt.Sprintf("%s%d: %s\n", indentString(indent+1), i, formatComplex(v.Index(i), indent+1))
		}
		res += fmt.Sprintf("%s]", indentString(indent))
		return res
	case reflect.Map:
		res := "map {\n"
		for _, key := range v.MapKeys() {
			keyString := formatComplex(key, indent+1)
			valueString := formatComplex(v.MapIndex(key), indent+1)
			res += fmt.Sprintf("%s%s: %s\n", indentString(indent+1), keyString, valueString)
		}
		res += fmt.Sprintf("%s}", indentString(indent))
		return res
	case reflect.Struct:
		res := "struct {\n"
		for i := 0; i < v.NumField(); i++ {
			nameString := v.Type().Field(i).Name
			valueString := formatComplex(v.Field(i), indent+1)
			res += fmt.Sprintf("%s%s: %s\n", indentString(indent+1), nameString, valueString)
		}
		res += fmt.Sprintf("%s}", indentString(indent))
		return res
	case reflect.Pointer:
		if v.IsNil() {
			return formatAtom(v)
		}
		return fmt.Sprintf("%s*%s: %s", indentString(indent), v.Elem().Type(), formatComplex(v.Elem(), indent))
	default:
		return formatAtom(v)
	}
}

func indentString(indent int) string {
	return strings.Repeat(" ", indentSize*indent)
}
