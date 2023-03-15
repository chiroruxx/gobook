package practice3_4

import (
	"fmt"
	"gobook/ch3/practice3_3"
	"gobook/ch3/surface"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, getSVG())
}

func getSVG() string {
	var items [][]float64
	for i := 0; i < surface.Cells; i++ {
		for j := 0; j < surface.Cells; j++ {
			ax, ay, az, aFailed := practice3_3.Corner(i+1, j)
			bx, by, bz, bFailed := practice3_3.Corner(i, j)
			cx, cy, cz, cFailed := practice3_3.Corner(i, j+1)
			dx, dy, dz, dFailed := practice3_3.Corner(i+1, j+1)

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
		"width='%d' height='%d'>\n", surface.Width, surface.Height)

	for _, item := range items {
		svg += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='#%s'/>\n",
			item[0], item[1], item[3], item[4], item[6], item[7], item[9], item[10], practice3_3.GetColorFromItem(item))
	}
	svg += "</svg>"

	return svg
}
