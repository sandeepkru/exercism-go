package account

import "sync"

// Account implements simple representation of bank account.
type Account struct {
	balance int64
	active  bool
	mutex   sync.Mutex
}

// Close implements account close functionality.
func (a *Account) Close() (payout int64, ok bool) {
	//defer a.mutex.Unlock()
	a.mutex.Lock()
	if a.active {
		payout = a.balance
		ok = true
		a.active = false
	} else { // already closed
		ok = false
	}

	a.mutex.Unlock()
	return
}

// Balance return account balance if it is active.
func (a *Account) Balance() (balance int64, ok bool) {
	//defer a.mutex.Unlock()
	a.mutex.Lock()
	if a.active {
		ok = true
		balance = a.balance
	} else {
		ok = false
	}
	a.mutex.Unlock()
	return
}

// Deposit implements deposit API.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	//defer a.mutex.Unlock()
	a.mutex.Lock()
	if a.active && a.balance+amount > 0 {
		ok = true
		a.balance += amount
		newBalance = a.balance
	} else {
		ok = false
	}
	a.mutex.Unlock()
	return
}

// Open returns new account if initialDeposit is valid.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, active: true}
}
