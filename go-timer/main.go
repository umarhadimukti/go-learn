package main

import (
	"time"
	"fmt"
)

func timerTest() {
	timer := time.NewTimer(5 * time.Second) // set delay 5 second
	fmt.Println(time.Now()) // print current time
	timeAfterDelay := <- timer.C // receive timer
	fmt.Println(timeAfterDelay) // print timer after delay
}

func afterTimerTest() {
	timer := <- time.After(3 * time.Second)
	fmt.Println("Second timer:", timer)
}

func tickerTimerTest() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()
	done := make(chan struct{})
	go func() {
		for {
			select {
			case <- ticker.C:
				fmt.Println("tick..")
			case <- done:
				fmt.Println("stop")
				return
			}
		}
	}()
	time.Sleep(10 * time.Second)
	close(done)
}

func main() {
	// timerTest()
	// afterTimerTest()
	tickerTimerTest()
}