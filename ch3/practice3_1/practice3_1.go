package practice3_1

import (
	"gobook/ch3/surface"
	"math"
)

func Corner(i, j int) (float64, float64, bool) {
	x := surface.XyRange * (float64(i)/surface.Cells - 0.5)
	y := surface.XyRange * (float64(j)/surface.Cells - 0.5)

	z := surface.F(x, y)

	if math.IsNaN(z) {
		return 0, 0, true
	}

	sx := surface.Width/2 + (x-y)*surface.Cos30*surface.XyScale
	sy := surface.Height/2 + (x+y)*surface.Sin30*surface.XyScale - z*surface.ZScale
	return sx, sy, false
}
