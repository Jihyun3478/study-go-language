package main

import (
	"../study-go-language/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("jihyun")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
