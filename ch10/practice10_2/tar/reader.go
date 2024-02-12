package tar

import (
	"archive/tar"
	"io"
	"os"

	"gobook/ch10/practice10_2"
)

type Reader struct{}

// Read file from filepath and return as Archive
func (r Reader) Read(path string) (practice10_2.Archive, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var headers []*tar.Header
	tr := tar.NewReader(f)
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		headers = append(headers, h)
	}

	res := &File{
		headers: headers,
	}

	return res, nil
}

func init() {
	practice10_2.RegisterReader(".tar", &Reader{})
}

// File is a tar file.
type File struct {
	headers []*tar.Header
}

// Size is total size of files in tar.
func (f File) Size() int {
	var n int64
	for _, h := range f.headers {
		n += h.Size
	}

	return int(n)
}

// Names are file names in tar.
func (f File) Names() []string {
	names := make([]string, len(f.headers))
	for i := 0; i < len(names); i++ {
		names[i] = f.headers[i].Name
	}

	return names
}
