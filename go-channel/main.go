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