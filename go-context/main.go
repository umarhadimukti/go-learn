package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func contextWithValueTest() {
	contextA := context.Background() // by default is empty
	
	contextB := context.WithValue(contextA, "b", "Bajigur")
	contextC := context.WithValue(contextA, "c", "Cuanki")

	contextD := context.WithValue(contextB, "d", "Dadar")
	contextE := context.WithValue(contextC, "e", "Endok")
	contextF := context.WithValue(contextD, "f", "Fufu")

	fmt.Println("context A:", contextA)
	fmt.Println("context B:", contextB)
	fmt.Println("context C:", contextC)
	fmt.Println("context D:", contextD)
	fmt.Println("context E:", contextE)
	fmt.Println("context F:", contextF)
	
	fmt.Println("====== Get value from parent ======")

	fmt.Println("Context F Hierarki:", contextF.Value("f"))
	fmt.Println("Context F Hierarki:", contextF.Value("d"))
	fmt.Println("Context F Hierarki:", contextF.Value("b"))
	fmt.Println("==================================")
}

func contextWithCancelTest(ctx context.Context) chan int  {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <- ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()
	return destination
}

func main() {
	// run contextWithValueTest function
	contextWithValueTest()

	// run contextWithCancel function
	fmt.Println("Goroutines before WithCancel:", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := contextWithCancelTest(ctx)
	fmt.Println("Goroutines while processing:", runtime.NumGoroutine())
	for dest := range destination {
		fmt.Println("Destination:", dest)
		if dest == 10 {
			break
		}
	}
	cancel() // cancel() method will send trigger to ctx.Done()

	time.Sleep(2 * time.Second)

	fmt.Println("Goroutines after WithCancel:", runtime.NumGoroutine())
}