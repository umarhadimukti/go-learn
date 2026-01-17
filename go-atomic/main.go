package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func incrementPrimitive() {
	atomic.AddInt64(&counter, 1)
	// counter++
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementPrimitive()
		}()
	}
	wg.Wait()
	fmt.Println("Counter with atomic:", counter)
}