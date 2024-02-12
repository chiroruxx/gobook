package main

import (
	"flag"
	"fmt"
	_ "image/png"
	"os"
	"strings"

	"gobook/ch10/practice10_1"
)

func main() {
	var fileType practice10_1.Typ

	flag.Func("type", "type of output image", func(val string) error {
		typ := practice10_1.Typ(val)
		if !practice10_1.ValidateType(typ) {
			var supports []string
			for _, support := range practice10_1.Supports() {
				supports = append(supports, string(support))
			}

			return fmt.Errorf("unsupported image type %s, supports %s", val, strings.Join(supports, ","))
		}

		fileType = typ

		return nil
	})
	flag.Parse()

	if fileType == "" {
		_, _ = fmt.Fprintln(os.Stderr, "jpeg: flag type is required")
		os.Exit(1)
	}

	if err := practice10_1.Convert(os.Stdin, os.Stdout, fileType); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}
