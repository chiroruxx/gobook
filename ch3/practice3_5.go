package main

import (
	"gobook/ch3/practice3_5"
	"image"
	"image/png"
	"os"
)

func main() {
	const (
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(yMax-yMin) + yMin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xMax-xMin) + xMin
			z := complex(x, y)
			img.Set(px, py, practice3_5.Mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}
