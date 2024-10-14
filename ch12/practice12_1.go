package main

import (
	"gobook/ch12/practice12_1"
)

func main() {
	type MovieID struct {
		Title string
	}
	type Movie struct {
		MovieID
		Subtitle string
		Year     int
		Color    bool
		Actor    map[string]string
		Oscars   []string
		Sequel   *string
	}

	strangeloveID := MovieID{
		Title: "Dr. Strangelove",
	}
	strangelove := Movie{
		MovieID:  strangeloveID,
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
	}

	evaluates := map[MovieID]uint{
		strangeloveID: 3,
	}

	practice12_1.Display("strangelove info", strangelove)
	practice12_1.Display("strangelove evaluates", evaluates)
}
