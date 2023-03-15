package main

import (
	"gobook/ch3/practice3_6"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, practice3_6.Width, practice3_6.Height))
	for py := 0; py < practice3_6.Height; py++ {
		for px := 0; px < practice3_6.Width; px++ {
			var colorValue color.Color
			if px-1 == 0 || py-1 == 0 || px+1 == practice3_6.Width || py+1 == practice3_6.Height {
				colorValue = practice3_6.GetColorValue(float64(px), float64(py))
			} else {
				x0 := float64(px) - 0.5
				x1 := float64(px) + 0.5
				y0 := float64(py) - 0.5
				y1 := float64(py) + 0.5

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
	png.Encode(os.Stdout, img)
}
