package main

import (
	"fmt"
	"gobook/ch6/coloredpoint"
	"gobook/ch6/geometry"
	"image/color"
)

func main() {
	var cp coloredpoint.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	var p = coloredpoint.ColoredPoint{Point: coloredpoint.Point{Point: geometry.Point{X: 1, Y: 1}}, Color: red}
	var q = coloredpoint.ColoredPoint{Point: coloredpoint.Point{Point: geometry.Point{X: 5, Y: 4}}, Color: blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}
