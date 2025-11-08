package accounts

import (
	"errors"
	"fmt"
)

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

// ChangeOwner changes the account owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner returns the account owner
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}