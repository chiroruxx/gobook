package main

import (
	"fmt"
	"gobook/ch9/practice9_1"
)

var finish = make(chan struct{})

func main() {
	fmt.Println("start")
	go func() {
		fmt.Println("start 1")
		practice9_1.Deposit(100)
		res := practice9_1.Withdraw(10)
		fmt.Printf("100 - 10: %v: %v\n", res, practice9_1.Balance())
		finish <- struct{}{}
	}()

	go func() {
		fmt.Println("start 2")
		//practice9_1.Deposit(1000)
		res := practice9_1.Withdraw(100)
		fmt.Printf("1000 - 100: %v: %v\n", res, practice9_1.Balance())
		finish <- struct{}{}
	}()

	<-finish
	<-finish
	fmt.Println("finished")
}
