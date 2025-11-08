package main

import (
	"fmt"
	"../study-go-language/accounts"
)

func main() {
	account := accounts.NewAccount("jihyun")
	account.Deposit(10)
	fmt.Println(account)
}
