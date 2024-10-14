package main

import (
	"fmt"
	"os"

	"gobook/ch12/sexpr"
)

func main() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	data := `((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Color nil) (Actor (("Dr. Strangelove" "Peter Sellers") ("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Strerling Handen"))) (Oscars ("Best Actor (Normin.)" "Best Adapt Screenplay (Normin.)" "Best Director (Normin.)" "Best Picture (Normin.)")) (Sequel nil))`
	var movie Movie
	err := sexpr.Unmarshal([]byte(data), &movie)
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
