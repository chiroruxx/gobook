package practice3_3

import (
	"gobook/ch3/surface"
	"math"
	"strconv"
)

var zMin = 1000.0
var zMax = 0.0

func Corner(i, j int) (float64, float64, float64, bool) {
	x := surface.XyRange * (float64(i)/surface.Cells - 0.5)
	y := surface.XyRange * (float64(j)/surface.Cells - 0.5)

	z := surface.F(x, y) * surface.ZScale

	if math.IsNaN(z) {
		return 0, 0, 0, true
	}

	zMin = math.Min(z, zMin)
	zMax = math.Max(z, zMax)

	sx := surface.Width/2 + (x-y)*surface.Cos30*surface.XyScale
	sy := surface.Height/2 + (x+y)*surface.Sin30*surface.XyScale - z
	return sx, sy, z, false
}

func GetColorFromItem(item []float64) string {
	z := (item[2] + item[5] + item[8] + item[11]) / 4
	return getColor(z)
}

func getColor(value float64) string {
	red := int64((value - zMin) / (zMax - zMin) * 255)
	blue := 255 - red

	redPrefix := ""
	if red < 16 {
		redPrefix = "0"
	}

	bluePrefix := ""
	if blue < 16 {
		bluePrefix = "0"
	}

	return redPrefix + strconv.FormatInt(red, 16) + "00" + bluePrefix + strconv.FormatInt(blue, 16)
}
