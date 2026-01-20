package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(3 * time.Second)
		}()
	}
	
	fmt.Println("===============")
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)
	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total Goroutines:", totalGoroutines)

	wg.Wait()
}