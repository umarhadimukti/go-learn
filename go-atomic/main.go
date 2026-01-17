package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/* atomic can process only single variable without mutex
	* if there is many variable, its better to use mutex than atomic
*/

var wg sync.WaitGroup
var counter int64

func incrementPrimitive() {
	atomic.AddInt64(&counter, 1)
	// counter++
}

func getCounter() {
	result := atomic.LoadInt64(&counter)
	fmt.Println("Counter with atomic:", result)
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
	getCounter()
}