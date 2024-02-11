package main

import (
	"fmt"
	"gobook/ch9/bank"
)

func main() {
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
	}()

	go bank.Deposit(100)
}
