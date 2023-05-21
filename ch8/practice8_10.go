package main

import (
	"gobook/ch8/practice8_10"
	"os"
)

func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)

	// 入力があったらキャンセルするゴルーチン
	go practice8_10.CancelFunc()()

	// ワークリストに挿入するゴルーチンの起動(秒で終わるはず)
	go func() {
		workList <- os.Args[1:]
	}()

	// クロールする20個のゴルーチンの起動
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := practice8_10.Crawl(link)
				go func() {
					workList <- foundLinks
				}()
			}
		}()
	}

	// ワークリストにあるリンクをひとつずつ unseenLinkに入れる
	seen := make(map[string]bool)
	for list := range workList {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
