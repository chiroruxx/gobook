package main

import (
	"fmt"
	"gobook/ch3/practice3_1"
	"gobook/ch3/surface"
)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", surface.Width, surface.Height)
	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {
			ax, ay, aFailed := practice3_1.Corner(i+1, j)
			bx, by, bFailed := practice3_1.Corner(i, j)
			cx, cy, cFailed := practice3_1.Corner(i, j+1)
			dx, dy, dFailed := practice3_1.Corner(i+1, j+1)

			if aFailed || bFailed || cFailed || dFailed {
				continue
			}

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}
