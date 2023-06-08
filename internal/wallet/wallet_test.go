package wallet_test

import (
	"github.com/theghostmac/wallet-engine/internal/wallet"
	"testing"
)

func TestNewWallet(t *testing.T) {
	// Arrange.
	initialBalance := wallet.Money(100)
	// Act.
	testWallet := wallet.NewWallet(initialBalance)
	// Assert.
	if testWallet.Balance() != initialBalance {
		t.Errorf("New wallet has different initial balance...")
	}
}

func TestDeposit(t *testing.T) {
	// Arrange.
	testWallet := wallet.NewWallet(0)
	testAmount := wallet.Money(50)
	// Act.
	testWallet.Deposit(testAmount)
	// Assert.
	if testWallet.Balance() != testAmount {
		t.Errorf("Deposit function doesn't increase the wallet balance...")
	}
}

func TestWithdraw(t *testing.T) {
	// Arrange.
	initialBalance := wallet.Money(100)
	testWallet := wallet.NewWallet(initialBalance)
	// Act.
	testAmount := wallet.Money(50)
	// Assert.
	if
}