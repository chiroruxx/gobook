package practice10_1

import (
	"fmt"
	"image"
	"io"
	"os"
)

// Typ is a file type of output
type Typ string

const (
	JPEG Typ = "jpeg"
	PNG  Typ = "png"
)

// Convert input to an image of given file type.
func Convert(in io.Reader, out io.Writer, fileType Typ) error {
	enc, err := resolveEncoder(fileType)
	if err != nil {
		return err
	}

	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(os.Stderr, "Input format=", kind)

	return enc.encode(out, img)
}

// Supports returns all supported output image types.
func Supports() []Typ {
	types := make([]Typ, len(encoders))
	i := 0
	for typ := range encoders {
		types[i] = typ
		i++
	}

	return types
}

// ValidateType returns given image type is supported.
func ValidateType(target Typ) bool {
	_, ok := encoders[target]
	return ok
}
