package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: practice10_4 package_name")
		os.Exit(1)
	}
	workSpace := args[1]
	pkgName := args[2]

	packages, err := getAllPackages(workSpace)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "practice10_4: %v\n", err)
		os.Exit(1)
	}

	results, err := filterDependedPackages(packages, pkgName)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "practice10_4: %v\n", err)
		os.Exit(1)
	}

	for _, result := range results {
		fmt.Println(result)
	}
}

func runListCommand(args ...string) ([]byte, error) {
	cmdArgs := []string{"list"}
	cmdArgs = append(cmdArgs, args...)
	list := exec.Command("go", cmdArgs...)

	output, err := list.Output()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return nil, errors.Join(err, errors.New(string(exitError.Stderr)))
		}
		return nil, err
	}

	return output, nil
}

func getAllPackages(workSpace string) ([]string, error) {
	output, err := runListCommand(workSpace + "/...")
	if err != nil {
		return nil, errors.Join(errors.New("failed to get work space list"), err)
	}

	packages := strings.Split(string(output), "\n")
	// last item is empty.
	packages = packages[:len(packages)-1]
	return packages, nil
}

type pkgInfo struct {
	Deps []string `json:"Deps"`
}

func (i pkgInfo) dependsOn(pkg string) bool {
	for _, dep := range i.Deps {
		if dep == pkg {
			return true
		}
	}

	return false
}

func getPackageInfo(pkg string) (*pkgInfo, error) {
	output, err := runListCommand("-json", pkg)
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to get package info of %s", pkg), err)
	}
	info := pkgInfo{}
	if err := json.Unmarshal(output, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

func filterDependedPackages(packages []string, depended string) ([]string, error) {
	var result []string
	for _, pkg := range packages {
		info, err := getPackageInfo(pkg)
		if err != nil {
			return nil, err
		}

		if info.dependsOn(depended) {
			result = append(result, pkg)
		}
	}

	return result, nil
}
