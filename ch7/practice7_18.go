package main

import (
	"encoding/xml"
	"fmt"
	"gobook/ch7/practice7_18"
	"io"
	"os"
	"strings"
)

func main() {
	//file, err := os.Open("./test2.xml")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "%v\n", err)
	//	os.Exit(1)
	//}
	//dec := xml.NewDecoder(file)
	dec := xml.NewDecoder(os.Stdin)
	var toks []xml.Token

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
		}

		switch tok := tok.(type) {
		case xml.StartElement, xml.EndElement:
			toks = append(toks, tok)
		case xml.CharData:
			data := strings.TrimSpace(string(tok))
			if data == "" {
				continue
			}

			toks = append(toks, xml.CharData(data))
		}
	}
	parser := practice7_18.NewParser()

	res := parser.Parse(&toks, nil)
	fmt.Println(res)
}
