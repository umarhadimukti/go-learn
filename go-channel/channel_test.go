package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Umar Jago ProgRamminG"
		fmt.Println("Successful send channel")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(3 * time.Second)
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	
	go ProcessChannel(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(6 * time.Second)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan int)
	defer close(channel)

	go ProcessInChannel(channel)
	go ProcessOutChannel(channel)

	time.Sleep(4 * time.Second)
}