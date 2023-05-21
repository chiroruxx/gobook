package practice8_2

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type command interface {
	action(conn net.Conn, u *user) (shouldClose bool)
}

func newCommand(name string, args []string) (command, error) {
	var c command
	switch name {
	case "ls":
		var relativePath string
		if len(args) > 0 {
			relativePath = args[0]
		}

		c = lsCommand{
			relativePath: relativePath,
		}
	case "cd":
		var newPath string
		if len(args) > 0 {
			newPath = args[0]
		}

		c = cdCommand{newPath}
	case "get":
		if len(args) == 0 {
			return nil, fmt.Errorf("usage: get <filename>")
		}

		c = getCommand{path: args[0]}
	case "close":
		c = closeCommand{}
	default:
		return nil, fmt.Errorf("command " + name + " is not defined.\n")
	}

	return c, nil
}

type lsCommand struct {
	relativePath string
}

func (c lsCommand) action(conn net.Conn, u *user) (shouldClose bool) {
	basePath := u.pwd()

	relativePath := c.relativePath
	if relativePath != "" && !strings.HasPrefix(relativePath, "/") {
		relativePath = "/" + relativePath
	}

	path := basePath + relativePath
	fileInfos, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	for _, fileInfo := range fileInfos {
		names = append(names, fileInfo.Name())
	}
	mustWriteConn(strings.Join(names, " "), conn)
	return
}

type cdCommand struct {
	path string
}

func (c cdCommand) action(conn net.Conn, u *user) (shouldClose bool) {
	userPathList := make([]string, len(u.currentDir))
	copy(userPathList, u.currentDir)

	path := c.path
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	pathList := strings.Split(path, "/")
	for _, path = range pathList {
		switch path {
		case "":
			continue
		case "..":
			if len(userPathList) == 0 {
				mustWriteConn("directory is not found: "+c.path, conn)
				return
			}
			userPathList = userPathList[:len(userPathList)-1]
		default:
			userPathList = append(userPathList, path)
		}
	}

	newRelativePath := strings.Join(userPathList, "/")
	if !fileExists(ftpDir + "/" + newRelativePath) {
		mustWriteConn("directory is not found: "+c.path, conn)
		return
	}

	u.currentDir = userPathList
	mustWriteConn("current directory is changed: "+"/"+newRelativePath, conn)
	return
}

type getCommand struct {
	path string
}

func (c getCommand) action(conn net.Conn, u *user) (shouldClose bool) {
	path := c.path
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	filepath := u.pwd() + "/" + c.path
	if !fileExists(filepath) {
		mustWriteConn("file is not found: "+c.path, conn)
		return
	}
	if isDirectory(filepath) {
		mustWriteConn("file is a directory: "+c.path, conn)
		return
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		mustWriteConn("cannot open file: "+c.path, conn)
		return
	}

	mustWriteConn(string(content), conn)
	return
}

type closeCommand struct{}

func (c closeCommand) action(net.Conn, *user) (shouldClose bool) {
	return true
}
