package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

type Account struct {
	balance int
	lock    sync.Mutex
	// mutable mutual exclusion,
	//which means that we will only allow one customer to access this function at one time
}

func (a *Account) GetBalance() int {
	a.lock.Lock() // lock the account down
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock() // lock the account so that nobody can access the account when withdrawing
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough money in account!")
	} else {
		a.balance -= v
		fmt.Printf("%d withdrawn; balance: %d\n", v, a.balance)

	}
}

func main() {
	var acct Account
	acct.balance = 100
	pl("Account Balance: ", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(10) // use go routine here
	}

	time.Sleep(2 * time.Second) // to pause the main function from running ahead
}
