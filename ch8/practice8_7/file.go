package practice8_7

import (
	"fmt"
	"os"
	"strings"
)

func writeFile(content []byte, u URL) error {
	path := filePath(u)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := makeDirectory(path); err != nil {
			return err
		}
	}

	return os.WriteFile(path, content, 0644)
}

func filePath(u URL) string {
	root, _ := os.Getwd()
	path := fmt.Sprintf("%s/mirror/%s", root, u.Host())
	uPath := u.Path()
	if uPath != "" {
		path = fmt.Sprintf("%s/%s", path, uPath)
	}
	if u.Extension() == "" {
		path += "/index.html"
	}

	return path
}

func fileLink(u URL) string {
	return fmt.Sprintf("file:/%s", filePath(u))
}

func makeDirectory(path string) error {
	pathList := strings.Split(path, "/")
	dir := strings.Join(pathList[:len(pathList)-1], "/")
	return os.MkdirAll(dir, 0755)
}
