package format2

import (
	"fmt"
	"reflect"
	"strings"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	fmt.Println(formatComplex(reflect.ValueOf(x), 0))
}

func display(v reflect.Value, indent int) {
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		printfWithIndent(indent, "slice {\n")
		for i := 0; i < v.Len(); i++ {
			printfWithIndent(indent+1, "%d: ", i)
			display(v.Index(i), indent+1)
		}
		printfWithIndent(indent, "}\n")
	case reflect.Struct:
		printfWithIndent(indent, "struct {\n")
		for i := 0; i < v.NumField(); i++ {
			printfWithIndent(indent+1, "%s: ", v.Type().Field(i).Name)
			display(v.Field(i), indent+1)
		}
		printfWithIndent(indent, "}\n")
	case reflect.Map:
		printfWithIndent(indent, "map {\n")
		for _, key := range v.MapKeys() {
			printfWithIndent(indent+1, "%s: ", key)
			display(v.MapIndex(key), indent+1)
		}
		printfWithIndent(indent, "}\n")
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("nil\n")
		} else {
			display(v.Elem(), indent)
		}
	case reflect.Interface:
		//if v.IsNil() {
		//	fmt.Printf("%s = nil\n", path)
		//} else {
		//	fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
		//	display(path+".value", v.Elem())
		//}
	default:
		fmt.Printf("%s\n", formatAtom(v))
	}
}

func printfWithIndent(indent int, format string, args ...any) {
	if indent > 0 {
		fmt.Printf(strings.Repeat(" ", indent*2))
	}
	fmt.Printf(format, args...)
}
