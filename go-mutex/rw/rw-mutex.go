package rw

import (
	_"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	Balance int
	mu sync.RWMutex
}

func (account *BankAccount) DepositWithRWMutex(amount int) {
	account.mu.Lock()
	defer account.mu.Unlock()
	account.Balance += amount
}

// calculate deposit with rw mutex
func CalcDepositWithRWMutex(account *BankAccount) {
	for i := 0; i < 10000; i++ {
		go account.DepositWithRWMutex(1)
	}
	time.Sleep(time.Second)
}

// get final balance with rw mutex
func GetBalanceWithRWMutex() int {
	account := &BankAccount{Balance: 0}
	CalcDepositWithRWMutex(account)
	
	// read lock
	account.mu.RLock()
	defer account.mu.RUnlock()

	return account.Balance
}