package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point2 struct{ X, Y float64 }

func (p *Point2) Distance(q Point2) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point2) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point2
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point2.X)
	cp.Point2.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	var p = ColoredPoint{Point2{1, 1}, red}
	var q = ColoredPoint{Point2{5, 4}, blue}
	fmt.Println(p.Distance(q.Point2))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point2))
}
