package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyRange       = 30.0
	xyScale       = width / 2 / xyRange
	zScale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aFailed := corner(i+1, j)
			bx, by, bFailed := corner(i, j)
			cx, cy, cFailed := corner(i, j+1)
			dx, dy, dFailed := corner(i+1, j+1)

			if aFailed || bFailed || cFailed || dFailed {
				continue
			}

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	z := f(x, y)

	if math.IsNaN(z) {
		return 0, 0, true
	}

	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z*zScale
	return sx, sy, false
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
