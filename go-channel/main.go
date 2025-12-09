package channel

import (
	"fmt"
	"time"
)

func ProcessChannel(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Sended from Process Channel."
	fmt.Println("Successfully sended!")
}

func ProcessInChannel(channel chan<- int) {
	channel <- 500
	fmt.Println("Successfully sended integer value")
}

func ProcessOutChannel(channel <-chan int) {
	data := <- channel
	fmt.Println(data)
}

func ProcessBufferChannel(channel chan<- []int) {
	channel <- []int{2,4,1,6,1,9,8}
	channel <- []int{3,5,6,2,5,4,2}
	channel <- []int{9,4,5,2,6,1,5}
	channel <- []int{5,1,5,6,2,3,6}
	fmt.Println("Successfully sended numbers")
}