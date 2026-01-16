package main

import (
	"fmt"
	"sync"
)

var counter = 0

func IncrementOnlyOnce() {
	counter++
}

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(IncrementOnlyOnce)
		}()
	}
	wg.Wait()
	fmt.Println("Final counter is", counter)
}