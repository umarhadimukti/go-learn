package channel

import (
	"fmt"
	"strconv"
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

func TestBufferChannel(t *testing.T) {
	channel := make(chan []int, 3)
	defer close(channel)

	go ProcessBufferChannel(channel)

	data1 := <- channel
	data2 := <- channel
	data3 := <- channel
	data4 := <- channel
	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)
	fmt.Println(data4)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- "Data ke-" + strconv.Itoa(i)
		}
	}()

	fmt.Println("Selesai mengirim data")

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan int)

	defer close(channel1)
	defer close(channel2)

	go func() {
		channel1 <- "Tattakae!!"
		channel2 <- 800
	}()

	var counter int
	for {
		select {
		case data1 := <- channel1:
			fmt.Println("this from channel 1:", data1)
			counter++
		case data2 := <- channel2:
			fmt.Println("this from channel 2:", data2)
			counter++
		default:
			fmt.Println("waiting for data..")
		}

		if counter == 2 {
			break
		}
	}
}