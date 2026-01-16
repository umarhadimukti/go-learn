package main

import (
	"fmt"
)

func addTwoNumbers(a, b int, result chan int) {
	sum := a + b
	result <- sum // send first data
	result <- sum + 5 // send second data
	result <- sum + 10 // send third data
	result <- sum + 20 // send fourth data
}

func main() {
	// first, make empty channel (3 capacities)
	emptyChannel1 := make(chan int, 3) // buffer: [] capacity: 3
	defer close(emptyChannel1)

	// second, run goroutines with function addTwoNumbers and send empty channel
	go addTwoNumbers(5, 6, emptyChannel1)

	// receive first data
	result1 := <- emptyChannel1 // buffer: [11] capacity: 3
	fmt.Println("Result is", result1)

	// receive second data
	result2 := <- emptyChannel1 // buffer: [11,16] capacity: 3
	fmt.Println("Result is", result2)

	// receive third data
	result3 := <- emptyChannel1 // buffer: [11,16,21] (full) capacity: 3
	fmt.Println("Result is", result3)

	// receive fourth data
	// blocked because buffer is full
	result4 := <- emptyChannel1 // then, buffer: [16,21,31] capacity: 3
	fmt.Println("Result is", result4)
}