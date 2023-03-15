package defer1

import "fmt"

func F(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	F(x - 1)
}
