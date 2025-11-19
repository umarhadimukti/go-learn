package main

import (
	"fmt"
	"errors"
)

type Student struct {
	Name string
	Age int
}

func (student *Student) validateAge() (int, error) {
	if student.Age < 15 {
		return 0, errors.New("age must be at least 15 years old")
	}
	return student.Age, nil
}

func main() {
	malik := Student{
		Name: "Malik",
		Age: 14,
	}

	age, err := malik.validateAge()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(fmt.Sprint(malik.Name) + " age is:", age)
	}
}