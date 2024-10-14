package practice12_5

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
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
		buf.WriteString("null")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		_, _ = fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		_, _ = fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.Float32, reflect.Float64:
		_, _ = fmt.Fprintf(buf, "%v", v.Float())
	case reflect.String:
		_, _ = fmt.Fprintf(buf, "%q", v.String())
	case reflect.Bool:
		_, _ = fmt.Fprintf(buf, "%v", v.Bool())
	case reflect.Ptr:
		return encode(buf, v.Elem(), indent)
	case reflect.Array, reflect.Slice:
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte(']')
	case reflect.Struct:
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			var innerBuf bytes.Buffer
			if err := encode(&innerBuf, v.Field(i), indent); err != nil {
				return err
			}
			if innerBuf.Len() != 0 {
				if i != 0 {
					buf.WriteByte(',')
				}
				_, _ = fmt.Fprintf(buf, "%q:%s", v.Type().Field(i).Name, innerBuf.String())
			}
		}
		buf.WriteByte('}')
	case reflect.Interface:
		// do nothing
	case reflect.Map:
		buf.WriteByte('{')
		keys := v.MapKeys()
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})

		for i, key := range keys {
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteByte(':')
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			if i != len(keys)-1 {
				buf.WriteByte(',')
			}
		}
		buf.WriteByte('}')
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
