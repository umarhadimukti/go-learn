package container_list

import (
	"container/list"
	"fmt"
)

func DemoLinkedList() {
	var data *list.List = list.New()

	data.PushBack("Apple")
	data.PushBack("Banana")
	data.PushBack("Cherry")

	data.PushFront("Elderberry")

	head := data.Front()
	for element := head; element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}