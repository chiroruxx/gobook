package surface

import (
	"math"
)

const (
	Height  = 320
	Width   = 600
	Cells   = 100
	XyRange = 30.0
	XyScale = Width / 2 / XyRange
	ZScale  = Height * 0.4
	Angle   = math.Pi / 6
)

var (
	Cos30 = math.Cos(Angle)
	Sin30 = math.Sin(Angle)
)

func Corner(i, j int) (float64, float64) {
	x := XyRange * (float64(i)/Cells - 0.5)
	y := XyRange * (float64(j)/Cells - 0.5)

	z := F(x, y)

	sx := Width/2 + (x-y)*Cos30*XyScale
	sy := Height/2 + (x+y)*Sin30*XyScale - z*ZScale
	return sx, sy
}

func F(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
