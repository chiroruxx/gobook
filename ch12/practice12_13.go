package main

import (
	"flag"
	"fmt"
	"os"

	"gobook/ch12/practice12_13"
)

func main() {
	t := flag.String("type", "encode", "encode or decode")
	flag.Parse()
	if *t == "encode" {
		encodeSexpr()
	} else {
		decodeSexpr()
	}
}

func encodeSexpr() {
	type Movie struct {
		Title    string            `sexpr:"title"`
		Subtitle string            `sexpr:"subtitle"`
		Year     int               `sexpr:"year"`
		Actor    map[string]string `sexpr:"actor"`
		Oscars   []string          `sexpr:"oscars"`
		Sequel   *string           `sexpr:"sequel"`
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
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

	s, err := practice12_13.Marshal(strangelove)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Printf("%s\n", s)
}

func decodeSexpr() {
	type Movie struct {
		Title    string            `sexpr:"title"`
		Subtitle string            `sexpr:"s"`
		Year     int               `sexpr:"year"`
		Color    bool              `sexpr:"color"`
		Actor    map[string]string `sexpr:"actor"`
		Oscars   []string          `sexpr:"oscars"`
		Sequel   *string           `sexpr:"sequel"`
	}

	data := `((title "Dr. Strangelove") (s "How I Learned to Stop Worrying and Love the Bomb") (Color nil) (Actor (("Dr. Strangelove" "Peter Sellers") ("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Strerling Handen"))) (Oscars ("Best Actor (Normin.)" "Best Adapt Screenplay (Normin.)" "Best Director (Normin.)" "Best Picture (Normin.)")) (Sequel nil))`
	var movie Movie
	err := practice12_13.Unmarshal([]byte(data), &movie)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Println("Title:", movie.Title)
	fmt.Println("Subtitle:", movie.Subtitle)
	fmt.Println("Year:", movie.Year)
	fmt.Println("Color:", movie.Color)
	fmt.Println("Actor:", movie.Actor)
	fmt.Println("Oscars:", movie.Oscars)
	fmt.Println("Sequel:", movie.Sequel)
}
