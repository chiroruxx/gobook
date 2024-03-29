package practice5_18

import (
	"io"
	"net/http"
	"os"
	"path"
)

func Fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
