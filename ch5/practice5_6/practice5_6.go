package practice5_6

import (
	"gobook/ch3/surface"
)

func Corner(i, j int) (sx float64, sy float64) {
	x := surface.XyRange * (float64(i)/surface.Cells - 0.5)
	y := surface.XyRange * (float64(j)/surface.Cells - 0.5)

	z := surface.F(x, y)

	sx = surface.Width/2 + (x-y)*surface.Cos30*surface.XyScale
	sy = surface.Height/2 + (x+y)*surface.Sin30*surface.XyScale - z*surface.ZScale
	return
}
