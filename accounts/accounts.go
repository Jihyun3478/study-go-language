package accounts

// Account struct
type Account struct {
	owner string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount to the account
func (a Account) Deposit(amount int) { // receiver method: (a Account)
	a.balance += amount
}

// Balance returns the account balance
func (a Account) Balance() int {
	return a.balance
}