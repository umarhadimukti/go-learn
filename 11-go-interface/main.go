package main

import (
	"fmt"
)

type HasPosition interface {
	GetPosition() string
	SetPosition(position string)
}

type Employee struct {
	Name, Position, Country string
	Age                     int
}

func (employee *Employee) GetPosition() string {
	return fmt.Sprint(employee.Name) + " is a " + fmt.Sprint(employee.Position)
}

func (employee *Employee) SetPosition(position string) {
	employee.Position = position
	Greeting(employee)
}

func Greeting(employee HasPosition) {
	fmt.Println("Greetings! " + employee.GetPosition())
}

type Message struct {
	ID                        int
	SenderName, Type, Content string
}

func ShowMessageType(message Message) {
	fmt.Println("Message ID:", message.ID)
	fmt.Println("Sender:", message.SenderName)
	fmt.Println("Type:", message.Type)
	fmt.Println("Content:", message.Content)
}

func main() {
	firstEmployee := Employee{"Umar", "Software Engineer", "Indonesia", 14}
	Greeting(&firstEmployee)
	firstEmployee.SetPosition("Senior Software Engineer")
	fmt.Println("=======================================")

	firstMessage := Message{1, "Umar Hadi", "private", "Hello, sir"}
	ShowMessageType(firstMessage)
	fmt.Println("=======================================")
}
