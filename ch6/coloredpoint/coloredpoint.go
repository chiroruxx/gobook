package coloredpoint

import (
	"gobook/ch6/geometry"
	"image/color"
)

type Point struct{ geometry.Point }

func (p *Point) Distance(q Point) float64 {
	return p.Point.Distance(q.Point)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}
