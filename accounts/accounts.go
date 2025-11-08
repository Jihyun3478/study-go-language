package accounts

import "errors"

// Account struct
type Account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("Can't withdraw more than the current balance")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount to the account
func (a *Account) Deposit(amount int) { // receiver method: (a Account)
	a.balance += amount
}

// Withdraw x amount from the account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Balance returns the account balance
func (a Account) Balance() int {
	return a.balance
}
