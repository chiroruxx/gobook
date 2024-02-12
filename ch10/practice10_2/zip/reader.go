package zip

import (
	"archive/zip"

	"gobook/ch10/practice10_2"
)

// Reader is a zip reader.
type Reader struct{}

// Read file from filepath and return as Archive
func (r Reader) Read(path string) (practice10_2.Archive, error) {
	zr, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	f := &File{
		files: zr.File,
	}
	return f, nil
}

func init() {
	practice10_2.RegisterReader(".zip", &Reader{})
}

// File is a zip file.
type File struct {
	files []*zip.File
}

// Names are file names in a zip.
func (f *File) Names() []string {
	names := make([]string, len(f.files))
	for i := 0; i < len(names); i++ {
		names[i] = f.files[i].Name
	}

	return names
}

// Size is total size of files in zip.
func (f *File) Size() int {
	var size uint64
	for _, file := range f.files {
		size += file.UncompressedSize64
	}

	return int(size)
}
