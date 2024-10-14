package practice12_3

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

func encode(buf *bytes.Buffer, v reflect.Value) error {
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
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		// (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')
	case reflect.Struct:
		// (name value ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			_, _ = fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
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
			if err := encode(buf, e); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
	case reflect.Map:
		// ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')
	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
