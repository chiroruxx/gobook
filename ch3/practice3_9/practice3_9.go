package practice3_9

import (
	"gobook/ch3/practice3_6"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	centerX := practice3_6.Width / 2
	x := r.FormValue("x")
	if x != "" {
		x, err := strconv.Atoi(x)
		if err == nil {
			centerX = x
		}
	}

	centerY := practice3_6.Height / 2
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

	img := image.NewRGBA(image.Rect(0, 0, practice3_6.Width, practice3_6.Height))
	for py := 0; py < practice3_6.Height; py++ {
		for px := 0; px < practice3_6.Width; px++ {
			var colorValue color.Color
			x := float64(centerX) + float64(px-practice3_6.Width/2.0)*unitPerCell
			y := float64(centerY) + float64(py-practice3_6.Width/2.0)*unitPerCell

			if px-1 == 0 || py-1 == 0 || px+1 == practice3_6.Width || py+1 == practice3_6.Height {
				colorValue = practice3_6.GetColorValue(x, y)
			} else {
				halfPerCell := unitPerCell / 2
				x0 := x - halfPerCell
				x1 := x + halfPerCell
				y0 := y - halfPerCell
				y1 := y + halfPerCell

				c0r, c0b, c0g, c0a := practice3_6.GetColorValue(x0, y0).RGBA()
				c1r, c1b, c1g, c1a := practice3_6.GetColorValue(x0, y1).RGBA()
				c2r, c2b, c2g, c2a := practice3_6.GetColorValue(x1, y0).RGBA()
				c3r, c3b, c3g, c3a := practice3_6.GetColorValue(x1, y1).RGBA()

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
