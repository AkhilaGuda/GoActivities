package main

type Account struct {
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}
