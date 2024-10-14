package practice12_4

import (
	"bytes"
	"fmt"
	"reflect"
)

type Inf interface {
	hasNextTitle(limit int) bool
}

type Test struct {
}

func (t Test) hasNextTitle(limit int) bool {
	return true
}

func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, _ = fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_, _ = fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		_, _ = fmt.Fprintf(buf, "%v", v.Float())
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		_, _ = fmt.Fprintf(buf, "#C(%v %v)", real(c), imag(c))
	case reflect.String:
		_, _ = fmt.Fprintf(buf, "%q", v.String())
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}
	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)
	case reflect.Array, reflect.Slice:
		// (value ...)
		buf.WriteByte('(')
		indent++
		for i := 0; i < v.Len(); i++ {
			ret(buf, indent)
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
			if i == v.Len()-1 {
				indent--
				ret(buf, indent)
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		// (name value ...)
		buf.WriteByte('(')
		indent++
		for i := 0; i < v.NumField(); i++ {
			ret(buf, indent)
			_, _ = fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i == v.NumField()-1 {
				indent--
				ret(buf, indent)
			}
		}
		buf.WriteByte(')')
	case reflect.Interface:
		// ("actual type" actual value)
		if v.IsNil() {
			buf.WriteString("nil")
		} else {
			buf.WriteByte('(')
			e := v.Elem()
			_, _ = fmt.Fprintf(buf, "%q ", e.Type())
			if err := encode(buf, e, indent); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
	case reflect.Map:
		// ((key value) ...)
		buf.WriteByte('(')
		indent++
		for i, key := range v.MapKeys() {
			ret(buf, indent)
			buf.WriteByte('(')
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			buf.WriteByte(')')
			if i == len(v.MapKeys())-1 {
				indent--
				ret(buf, indent)
			}
		}
		buf.WriteByte(')')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func ret(buf *bytes.Buffer, indent int) {
	buf.WriteByte('\n')
	for i := 0; i < indent; i++ {
		buf.WriteByte(' ')
	}
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
