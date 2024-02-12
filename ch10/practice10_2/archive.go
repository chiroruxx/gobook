package practice10_2

import (
	"fmt"
	"path/filepath"
)

// Reader is an archive file reader.
type Reader interface {
	Read(path string) (Archive, error)
}

var readers = make(map[string]Reader)

// RegisterReader register a reader.
func RegisterReader(ext string, r Reader) {
	readers[ext] = r
}

func resolveReader(ext string) (Reader, error) {
	r, ok := readers[ext]
	if !ok {
		return nil, fmt.Errorf("unsupported extension %s", ext)
	}

	return r, nil
}

// Archive is an archive file.
type Archive interface {
	Size() int
	Names() []string
}

// Read file from filepath and return as Archive
func Read(path string) (Archive, error) {
	ext := filepath.Ext(path)
	r, err := resolveReader(ext)
	if err != nil {
		return nil, err
	}
	return r.Read(path)
}
