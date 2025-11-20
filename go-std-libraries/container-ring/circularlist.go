package container_ring

import (
	"container/ring"
	"fmt"
	"strconv"
)

func DemoCircularList() {
	var data *ring.Ring = ring.New(5)
	for i := 1; i <= data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i) // assign values to each element
		data = data.Next() // move to the next element
	}
	data.Do(func(value interface{}) { // iterate through the ring
		fmt.Println(value)
	})
}