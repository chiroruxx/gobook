package main

import (
	"fmt"
	"math"
	"os"

	"gobook/ch12/practice12_5"
)

func main() {
	testInf := practice12_5.Test{}

	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
		Float           float64
		inf             practice12_5.Inf
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Strerling Handen",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Normin.)",
			"Best Adapt Screenplay (Normin.)",
			"Best Director (Normin.)",
			"Best Picture (Normin.)",
		},
		Float: math.Pi,
		inf:   testInf,
	}

	s, err := practice12_5.Marshal(strangelove)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Printf("%s\n", s)
}
