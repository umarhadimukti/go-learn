package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool sync.Pool = sync.Pool{
		New: func() interface{} {
			return "kosong" // default value if pool is empty
		},
	}
	var wg sync.WaitGroup

	// inject or put value for pool
	pool.Put("umar")
	pool.Put("john")
	pool.Put("doe")

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := pool.Get() // get data
			fmt.Println(data) // then print data
			pool.Put(data) // then put back data (preventing empty pool)
		}()
	}
	wg.Wait()
	fmt.Println("=== done ===")
}