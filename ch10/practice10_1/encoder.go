package practice10_1

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type encoder interface {
	encode(out io.Writer, img image.Image) error
}

var encoders = make(map[Typ]encoder)

func init() {
	registerEncoder(JPEG, jpegEncoder{})
	registerEncoder(PNG, pngEncoder{})
}

func registerEncoder(typ Typ, enc encoder) {
	encoders[typ] = enc
}

func resolveEncoder(fileType Typ) (encoder, error) {
	enc, ok := encoders[fileType]
	if !ok {
		return nil, fmt.Errorf("unsupport filetype %s", string(fileType))
	}

	return enc, nil
}

type jpegEncoder struct{}

func (e jpegEncoder) encode(out io.Writer, img image.Image) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

type pngEncoder struct{}

func (e pngEncoder) encode(out io.Writer, img image.Image) error {
	return png.Encode(out, img)
}
