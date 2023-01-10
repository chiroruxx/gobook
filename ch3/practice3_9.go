package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	xMin, yMin, xMax, yMax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	centerX := width / 2
	x := r.FormValue("x")
	if x != "" {
		x, err := strconv.Atoi(x)
		if err == nil {
			centerX = x
		}
	}

	centerY := height / 2
	y := r.FormValue("y")
	if y != "" {
		y, err := strconv.Atoi(y)
		if err == nil {
			centerY = y
		}
	}

	scale := 100
	s := r.FormValue("s")
	if s != "" {
		s, err := strconv.Atoi(s)
		if err == nil {
			scale = s
		}
	}

	img := draw(centerX, centerY, scale)
	png.Encode(w, img)
}

func draw(centerX int, centerY int, scale int) *image.RGBA {
	unitPerCell := 1 / (float64(scale) / 100.0)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var colorValue color.Color
			x := float64(centerX) + float64(px-width/2.0)*unitPerCell
			y := float64(centerY) + float64(py-height/2.0)*unitPerCell

			if px-1 == 0 || py-1 == 0 || px+1 == width || py+1 == height {
				colorValue = getColorValue(x, y)
			} else {
				halfPerCell := unitPerCell / 2
				x0 := x - halfPerCell
				x1 := x + halfPerCell
				y0 := y - halfPerCell
				y1 := y + halfPerCell

				c0r, c0b, c0g, c0a := getColorValue(x0, y0).RGBA()
				c1r, c1b, c1g, c1a := getColorValue(x0, y1).RGBA()
				c2r, c2b, c2g, c2a := getColorValue(x1, y0).RGBA()
				c3r, c3b, c3g, c3a := getColorValue(x1, y1).RGBA()

				colorValue = color.RGBA{
					R: uint8((c0r + c1r + c2r + c3r) / 4),
					G: uint8((c0g + c1g + c2g + c3g) / 4),
					B: uint8((c0b + c1b + c2b + c3b) / 4),
					A: uint8((c0a + c1a + c2a + c3a) / 4)}
			}
			img.Set(px, py, colorValue)
		}
	}
	return img
}

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

func getColorValue(px, py float64) color.Color {
	x := px/width*(xMax-xMin) + xMin
	y := py/height*(yMax-yMin) + yMin
	z := complex(x, y)
	return mandelbrot(z)
}
