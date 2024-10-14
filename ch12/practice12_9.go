package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gobook/ch12/practice12_9"
)

func main() {
	data := `((Title "Dr. Strangelove") (Subtitle "How I Learned to Stop Worrying and Love the Bomb") (Color nil) (Actor (("Dr. Strangelove" "Peter Sellers") ("Grp. Capt. Lionel Mandrake" "Peter Sellers") ("Pres. Merkin Muffley" "Peter Sellers") ("Gen. Buck Turgidson" "George C. Scott") ("Brig. Gen. Jack D. Ripper" "Strerling Handen"))) (Oscars ("Best Actor (Normin.)" "Best Adapt Screenplay (Normin.)" "Best Director (Normin.)" "Best Picture (Normin.)")) (Sequel nil))`
	r := strings.NewReader(data)
	decoder := practice12_9.NewDecoder(r)
	for {
		tok, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				return
			}

			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
		switch tok := tok.(type) {
		case practice12_9.StartList:
			fmt.Println("(")
		case practice12_9.EndList:
			fmt.Println(")")
		case practice12_9.Symbol:
			fmt.Println("Symbol:", tok.Value)
		case practice12_9.String:
			fmt.Println("String:", tok.Value)
		case practice12_9.Int:
			fmt.Println("Int:", tok.Value)
		default:
			_, _ = fmt.Fprintf(os.Stderr, "Unknown token %v\n", tok)
			return
		}
	}
}
