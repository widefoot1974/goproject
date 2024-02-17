package main

import "fmt"

type IBankAccount interface {
	GetAccountName() string
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {

	// wf := NewWellsFargo()
	// wf.Deposit(200)
	// wf.Deposit(100)
	// wf.Deposit(300)
	// if err := wf.Withdraw(100); err != nil {
	// 	panic(err)
	// }
	// currentBalance := wf.GetBalance()
	// fmt.Printf("WF balance: %d\n", currentBalance)

	// btc := NewBitcoinAccount()
	// btc.Deposit(200)
	// btc.Deposit(100)
	// btc.Deposit(300)
	// if err := btc.Withdraw(100); err != nil {
	// 	panic(err)
	// }
	// currentBalance = btc.GetBalance()
	// fmt.Printf("BTC balance: %d\n", currentBalance)

	myAccounts := []IBankAccount{
		NewWellsFargo(),
		NewBitcoinAccount(),
	}

	for _, account := range myAccounts {
		balance := account.GetBalance()
		fmt.Printf("[%s] default balance = %d\n", account.GetAccountName(), balance)

		account.Deposit(1000)
		if err := account.Withdraw(400); err != nil {
			fmt.Printf("account.Withdraw(400) fail: %v", err)
		}

		balance = account.GetBalance()
		fmt.Printf("[%s] balance = %d\n", account.GetAccountName(), balance)
	}

}
