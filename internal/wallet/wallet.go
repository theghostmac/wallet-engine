package wallet

import "sync"

// Money represents monetary values.
type Money int64

// Wallet represents a user's wallet.
type Wallet struct {
	WalletBalance Money
	mutex         sync.Mutex
}

// NewWallet creates a new Wallet instance with the given initial balance.
func NewWallet(initialBalance Money) *Wallet {
	return &Wallet{
		WalletBalance: initialBalance,
	}
}

// Deposit adds funds to the wallet.
func (w *Wallet) Deposit(amount Money) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.WalletBalance += amount
}

// Withdraw deducts funds from the wallet if sufficient balance is available.
func (w *Wallet) Withdraw(amount Money) bool {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.WalletBalance >= amount {
		w.WalletBalance -= amount
		return true
	}
	return false
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Money {
	return w.WalletBalance
}
