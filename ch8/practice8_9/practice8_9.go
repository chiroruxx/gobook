package practice8_9

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type SizePerRoot struct {
	Number int
	Size   int64
}

type Total struct {
	Path   string
	Nfiles int64
	Nbytes int64
}

func WalkDir(dir string, number int, n *sync.WaitGroup, fileSizes chan<- SizePerRoot) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go WalkDir(subdir, number, n, fileSizes)
		} else {
			fileSizes <- SizePerRoot{
				Number: number,
				Size:   entry.Size(),
			}
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}

	infoList := make([]os.FileInfo, len(entries))
	for i, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Fprintf(os.Stderr, "du1: %v\n", err)
			return nil
		}
		infoList[i] = info
	}

	return infoList
}

func PrintDiskUsage(totals map[int]Total) {
	for _, total := range totals {
		fmt.Printf("%s:	%d files	%.1f GB\n", total.Path, total.Nfiles, float64(total.Nbytes)/1e9)
	}
	fmt.Println()
}
