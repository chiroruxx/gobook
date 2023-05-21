package practice8_2

import (
	"path/filepath"
	"strings"
)

const ftpDir = "ftp"

type user struct {
	currentDir []string
}

func newUser() *user {
	var currentDir []string
	u := user{currentDir: currentDir}

	return &u
}

func (u *user) pwd() string {
	baseDir, _ := filepath.Abs(ftpDir)

	if len(u.currentDir) == 0 {
		return baseDir
	}

	return baseDir + "/" + strings.Join(u.currentDir, "/")
}
