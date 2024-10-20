package practice13_2

import (
	"reflect"
	"unsafe"
)

type comparison struct {
	ptr unsafe.Pointer
	typ reflect.Type
}

func IsLoopStruct(x any) bool {
	seen := make(map[comparison]bool)
	return isLoopStruct(reflect.ValueOf(x), seen)
}

func isLoopStruct(x reflect.Value, seen map[comparison]bool) bool {
	if x.CanAddr() {
		k := comparison{
			unsafe.Pointer(x.UnsafeAddr()),
			x.Type(),
		}
		if seen[k] {
			return true
		}
		seen[k] = true
	}

	switch x.Kind() {
	case reflect.Ptr:
		return isLoopStruct(x.Elem(), seen)
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			field := x.Field(i)
			if isLoopStruct(field, seen) {
				return true
			}
			if !field.CanAddr() {
				continue
			}

			k := comparison{
				unsafe.Pointer(field.UnsafeAddr()),
				field.Type(),
			}
			seen[k] = true
		}
		return false
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isLoopStruct(x.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isLoopStruct(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	default:
		return false
	}
}
