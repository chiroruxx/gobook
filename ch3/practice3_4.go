package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyRange       = 30.0
	xyScale       = width / 2 / xyRange
	zScale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var zMin = 1000.0
var zMax = 0.0

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, getSVG())
}

func getSVG() string {
	var items [][]float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, aFailed := corner(i+1, j)
			bx, by, bz, bFailed := corner(i, j)
			cx, cy, cz, cFailed := corner(i, j+1)
			dx, dy, dz, dFailed := corner(i+1, j+1)

			if aFailed || bFailed || cFailed || dFailed {
				continue
			}

			var item []float64
			item = append(item, ax)
			item = append(item, ay)
			item = append(item, az)
			item = append(item, bx)
			item = append(item, by)
			item = append(item, bz)
			item = append(item, cx)
			item = append(item, cy)
			item = append(item, cz)
			item = append(item, dx)
			item = append(item, dy)
			item = append(item, dz)

			items = append(items, item)
		}
	}

	svg := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	for _, item := range items {
		svg += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='#%s'/>\n",
			item[0], item[1], item[3], item[4], item[6], item[7], item[9], item[10], getColorFromItem(item))
	}
	svg += "</svg>"

	return svg
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyRange * (float64(i)/cells - 0.5)
	y := xyRange * (float64(j)/cells - 0.5)

	z := f(x, y) * zScale

	if math.IsNaN(z) {
		return 0, 0, 0, true
	}

	zMin = math.Min(z, zMin)
	zMax = math.Max(z, zMax)

	sx := width/2 + (x-y)*cos30*xyScale
	sy := height/2 + (x+y)*sin30*xyScale - z
	return sx, sy, z, false
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func getColorFromItem(item []float64) string {
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
