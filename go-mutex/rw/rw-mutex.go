package rw

import (
	"sync"
)

type BankAccount struct {
	Balance int
	mu      sync.RWMutex
}

func (account *BankAccount) DepositWithRWMutex(amount int) {
	account.mu.Lock()
	defer account.mu.Unlock()
	account.Balance += amount
}

func (account *BankAccount) GetFinalBalance() int {
	account.mu.RLock()
	defer account.mu.RUnlock()
	balance := account.Balance
	return balance
}

// calculate deposit with rw mutex
func CalcDepositWithRWMutex(account *BankAccount) {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			account.DepositWithRWMutex(1)
		}()
	}
	wg.Wait()
}

// get final balance with rw mutex
func GetBalanceWithRWMutex() int {
	account := &BankAccount{Balance: 0}

	CalcDepositWithRWMutex(account)
	finalBalance := account.GetFinalBalance()

	return finalBalance
}
