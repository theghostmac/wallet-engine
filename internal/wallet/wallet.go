package wallet

import "sync"

// Wallet represents a user's wallet.
type money int64

// Wallet represents a user's wallet.
type Wallet struct {
	balance money
	mutex   sync.Mutex
}

// NewWallet creates a new Wallet instance with the given initial balance.
func NewWallet(initialBalance money) *Wallet {
	return &Wallet{
		balance: initialBalance,
	}
}

// Deposit adds funds to the wallet.
func (w *Wallet) Deposit(amount money) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.balance += amount
}

// Withdraw deducts funds from the wallet if sufficient balance is available.
func (w *Wallet) Withdraw(amount money) bool {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	if w.balance >= amount {
		w.balance -= amount
		return true
	}
	return false
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() money {
	return w.balance
}
