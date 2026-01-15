package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	Balance int
	mu sync.RWMutex
}

func (account *BankAccount) Deposit(amount int) {
	account.mu.Lock()
	defer account.mu.Unlock()
	account.Balance += amount
}

func (account *BankAccount) GetBalance() int {
	account.mu.RLock()
	defer account.mu.RUnlock()
	return account.Balance
}

func main() {
	account := BankAccount{Balance:0}
	var wg sync.WaitGroup

	// test with 10 writers
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				account.Deposit(1)
				time.Sleep(time.Millisecond * 1)
			}
		}(i)
	}

	// test with 1000 readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				balance := account.GetBalance()
				fmt.Printf("Reader %d sees balance: %d\n", id, balance)
				time.Sleep(time.Millisecond * 5)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final balance: %d\n", account.Balance)
}