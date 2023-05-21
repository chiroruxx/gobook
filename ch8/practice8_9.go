package main

import (
	"flag"
	"gobook/ch8/practice8_9"
	"sync"
	"time"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan practice8_9.SizePerRoot)
	totals := make(map[int]practice8_9.Total, len(roots))
	var n sync.WaitGroup
	for rootNumber, root := range roots {
		totals[rootNumber] = practice8_9.Total{Path: root}
		n.Add(1)
		go practice8_9.WalkDir(root, rootNumber, &n, fileSizes)
	}

	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	tick = time.Tick(500 * time.Millisecond)

loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			total := totals[size.Number]
			total.Nfiles++
			total.Nbytes += size.Size
			totals[size.Number] = total
		case <-tick:
			practice8_9.PrintDiskUsage(totals)
		}
	}
	practice8_9.PrintDiskUsage(totals)
}
