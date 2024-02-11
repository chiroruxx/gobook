package practice9_4

import "fmt"

var print bool

func SetPrint(p bool) {
	print = p
}

func println(s string) {
	if !print {
		return
	}

	fmt.Println(s)
}

func printf(format string, args ...any) {
	if !print {
		return
	}

	fmt.Printf(format, args...)
}
