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