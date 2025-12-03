package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func DisplayNumber(num int) {
	fmt.Println("Display Number:", num)
}

func TestCreateGoRoutines(t *testing.T) {
	counter := 10 // can run through 1.000.000 go routines
	for count := range counter {
		go DisplayNumber(count)
	}
	fmt.Println("AKU")
	time.Sleep(5 * time.Second)
}