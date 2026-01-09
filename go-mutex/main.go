package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	Balance int
	mu sync.Mutex
}

func (account *BankAccount) DepositWithMutex(amount int) {
	account.mu.Lock()
	defer account.mu.Unlock()
	account.Balance += amount
}

func (account *BankAccount) DepositWithoutMutex(amount int) {
	account.Balance += amount
}

func BalanceWithMutex() {
	account := BankAccount{Balance: 0}
	for i := 0; i < 10000; i++ {
		go account.DepositWithMutex(1)
	}
	time.Sleep(time.Second)
	fmt.Println("Final balance with:", account.Balance)
}

func BalanceWithoutMutex() {
	account := BankAccount{Balance: 0}
	for i := 0; i < 10000; i++ {
		go account.DepositWithoutMutex(1)
	}
	time.Sleep(time.Second)
	fmt.Println("Final balance without mutex:", account.Balance)
}

func main() {
	BalanceWithMutex()
	BalanceWithoutMutex()
}