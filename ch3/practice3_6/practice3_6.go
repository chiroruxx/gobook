package practice3_6

import (
	"image/color"
	"math/cmplx"
)

const (
	xMax, yMax, xMin, yMin = +2, +2, -2, -2
	Height                 = 1024
	Width                  = 1024
)

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 30

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{R: 255 - contrast*n, G: 120, B: contrast * n, A: 255}
			//return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

func GetColorValue(px, py float64) color.Color {
	x := px/Width*(xMax-xMin) + xMin
	y := py/Height*(yMax-yMin) + yMin
	z := complex(x, y)
	return mandelbrot(z)
}
