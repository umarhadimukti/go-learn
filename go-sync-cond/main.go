package main

import (
	"fmt"
	"sync"
	"time"
)

type Cashier struct {
	queue []string
	mu sync.Mutex // protect queue
	cond *sync.Cond // tools for waiting and waking up
}

func NewCashier() *Cashier {
	var c *Cashier = &Cashier{} // declare new cashier
	c.cond = sync.NewCond(&c.mu) // c.cond must be reference with c.mu
	return c
}

func (c *Cashier) CustomerArrive(name string) {
	defer c.mu.Unlock()
	c.mu.Lock()
	// append customer to the last queue
	c.queue = append(c.queue, name)
	fmt.Println("Customer arrive:", name)
	// send signal to the cashier that there is a customer arrive
	c.cond.Signal()
}

/**
 * serve customer function (pov cashier)
*/
func (c *Cashier) ServeCustomer() {
	defer c.mu.Unlock()
	c.mu.Lock()
	// loop while length of queue is empty
	for len(c.queue) == 0 {
		fmt.Println("Waiting...")
		c.cond.Wait() // cond.Wait() makes go-routines sleep and unlock mutex until cond.Signal() called
	}
	// get the first customer, then remove them from queue
	customer := c.queue[0]
	c.queue = c.queue[1:]

	fmt.Println("Serving customer:", customer)
}

func main() {
	cashier := NewCashier()
	
	// cashier working
	go func() {
		for {
			cashier.ServeCustomer()
			time.Sleep(time.Second)
		}
	}()

	// customer arrives
	customers := []string{"Umar", "Yami", "Asta", "Luck"}
	for _, name := range customers {
		time.Sleep(2 * time.Second)
		cashier.CustomerArrive(name)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}