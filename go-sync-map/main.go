package main

import (
	"fmt"
	"sync"
)

var cache sync.Map

func Set(key string, val interface{}) {
	cache.Store(key, val)
}

func Get(key string) (interface{}, bool) {
	val, ok := cache.Load(key)
	if !ok {
		return nil, false
	}
	return val, true
}

func GetAll() {
	cache.Range(func(key, val interface{}) bool {
		fmt.Println(fmt.Sprint(key) + ":", val)
		return true
	})
}

func main() {
	var wg sync.WaitGroup

	Set("apple", "red")
	Set("pear", "yellow")
	Set("random_access_memory", 16)
	ram, ram_ok := Get("random_access_memory")
	if ram_ok {
		fmt.Println("Total RAM:", fmt.Sprint(ram) + "GB")
	}
	fmt.Println("================")
	GetAll()
	fmt.Println("======================")
	
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			Set("sayur_" + fmt.Sprint(num), "Mangga " + fmt.Sprint(num))
		}(i)
	}
	wg.Wait()

	
	fmt.Println("After go-routines:")
	GetAll()
}