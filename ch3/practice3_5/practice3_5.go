package practice3_5

import (
	"image/color"
	"math/cmplx"
)

func Mandelbrot(z complex128) color.Color {
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
