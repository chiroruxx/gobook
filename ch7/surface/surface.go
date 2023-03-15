package surface

import (
	"fmt"
	"gobook/ch7/eval"
	"io"
	"math"
	"net/http"
)

import sflib "gobook/ch3/surface"

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		fmt.Println("err", s)
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func Plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+html")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", sflib.Width, sflib.Height)
	for i := 0; i < sflib.Cells; i++ {
		for j := 0; j < sflib.Cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64) {
	x := sflib.XyRange * (float64(i)/sflib.Cells - 0.5)
	y := sflib.XyRange * (float64(j)/sflib.Cells - 0.5)

	z := f(x, y)

	sx := sflib.Width/2 + (x-y)*sflib.Cos30*sflib.XyScale
	sy := sflib.Height/2 + (x+y)*sflib.Sin30*sflib.XyScale - z*sflib.ZScale
	return sx, sy
}
