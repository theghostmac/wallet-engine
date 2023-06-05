package main

import (
	"fmt"
	"github.com/theghostmac/wallet-engine/internal/wallet"
)

func main() {
	// Create a new wallet with an initial balance of 100.
	w := wallet.NewWallet(100)

	// Perform some operations on the wallet.
	w.Deposit(50)
	fmt.Println("Current balance:\t ", w.Balance())

	success := w.Withdraw(75)
	if success {
		fmt.Println("Withdraw successful")
	} else {
		fmt.Println("Insufficient balance")
	}

	fmt.Println("Current balance:\t ", w.Balance())
}
