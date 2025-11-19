package main

import (
	"fmt"
	"errors"
	errors_pkg "practice-go/go-std-libraries/errors"
	os_pkg "practice-go/go-std-libraries/os"
)

func main() {
	fmt.Println("---Errors Library---")

	student, err := errors_pkg.GetStudentById(2)
	if err != nil {
		if errors.Is(err, errors_pkg.ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, errors_pkg.NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("Unknown Error:", err)
		}
	} else {
		fmt.Println("Student found:", student)
	}

	fmt.Println("\n---Os Library---") 
	os_pkg.ReadFile("go-std-libraries/os/data.txt")
	os_pkg.WriteFile("go-std-libraries/os/data.txt", "Let's go, common!")
}