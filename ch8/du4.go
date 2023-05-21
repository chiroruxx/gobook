package main

import (
	"flag"
	"gobook/ch8/du1"
	"gobook/ch8/du4"
	"os"
	"sync"
	"time"
)

var verbose3 = flag.Bool("v", false, "show verbose progress message")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go du4.WalkDir(root, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(du4.Done)
	}()

	var tick <-chan time.Time
	if *verbose3 {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-du4.Done:
			for range fileSizes {
				// do nothing
			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			du1.PrintDiskUsage(nfiles, nbytes)
		}
	}
	du1.PrintDiskUsage(nfiles, nbytes)
}
