package practice12_6

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
		i := v.Int()
		if i == 0 {
			return nil
		}
		_, _ = fmt.Fprintf(buf, "%d", i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		u := v.Uint()
		if u == 0 {
			return nil
		}
		_, _ = fmt.Fprintf(buf, "%d", u)
	case reflect.Float32, reflect.Float64:
		f := v.Float()
		if f == 0 {
			return nil
		}
		_, _ = fmt.Fprintf(buf, "%v", f)
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		if c == 0 {
			return nil
		}
		_, _ = fmt.Fprintf(buf, "#C(%v %v)", real(c), imag(c))
	case reflect.String:
		s := v.String()
		if s == "" {
			return nil
		}
		_, _ = fmt.Fprintf(buf, "%q", s)
	case reflect.Bool:
		if v.Bool() {
			buf.WriteByte('t')
		}
	case reflect.Ptr:
		return encode(buf, v.Elem())
	case reflect.Array, reflect.Slice:
		// (value ...)
		var buf1 bytes.Buffer
		for i := 0; i < v.Len(); i++ {
			var buf2 bytes.Buffer
			if err := encode(&buf2, v.Index(i)); err != nil {
				return err
			}
			if buf1.Len() != 0 {
				buf1.WriteByte(' ')
			}
			buf1.WriteString(buf2.String())
		}
		if buf1.Len() != 0 {
			_, _ = fmt.Fprintf(buf, "(%s)", buf1.String())
		}
	case reflect.Struct:
		// (name value ...)
		var buf1 bytes.Buffer
		for i := 0; i < v.NumField(); i++ {
			var buf2 bytes.Buffer
			if err := encode(&buf2, v.Field(i)); err != nil {
				return err
			}
			if buf2.Len() == 0 {
				continue
			}
			if buf1.Len() != 0 {
				buf1.WriteByte(' ')
			}

			_, _ = fmt.Fprintf(&buf1, "(%s %s)", v.Type().Field(i).Name, buf2.String())
		}
		if buf1.Len() != 0 {
			_, _ = fmt.Fprintf(buf, "(%s)", buf1.String())
		}
	case reflect.Interface:
		// ("actual type" actual value)
		if !v.IsNil() {
			var buf1 bytes.Buffer
			e := v.Elem()
			if err := encode(&buf1, e); err != nil {
				return err
			}
			if buf1.Len() != 0 {
				_, _ = fmt.Fprintf(buf, "(%s)", buf1.String())
			}
		}
	case reflect.Map:
		// ((key value) ...)
		var buf1 bytes.Buffer
		for _, key := range v.MapKeys() {
			var buf2 bytes.Buffer
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			if buf2.Len() == 0 {
				continue
			}
			if buf1.Len() != 0 {
				buf1.WriteByte(' ')
			}
			buf1.WriteByte('(')
			if err := encode(&buf1, key); err != nil {
				return err
			}
			_, _ = fmt.Fprintf(&buf1, " %s)", buf2.String())
		}
		if buf1.Len() != 0 {
			_, _ = fmt.Fprintf(buf, "(%s)", buf1.String())
		}
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
