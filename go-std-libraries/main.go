package main

import (
	"fmt"
	"errors"
	errors_pkg "practice-go/go-std-libraries/errors"
)

func main() {
	fmt.Println("---Errors Libraries---")

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
}