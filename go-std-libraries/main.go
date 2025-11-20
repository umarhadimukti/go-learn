package main

import (
	"fmt"
	"errors"
	errors_pkg "practice-go/go-std-libraries/errors"
	os_pkg "practice-go/go-std-libraries/os"
	flag_pkg "practice-go/go-std-libraries/flag"
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
	fmt.Println("Read file:")
	os_pkg.ReadFile("go-std-libraries/os/data.txt")
	fmt.Println("Write file:")
	os_pkg.WriteFile("go-std-libraries/os/data.txt", "Let's go, common!")
	fmt.Println("Read args:")
	os_pkg.ReadArgs()
	fmt.Println("Get hostname:")
	os_pkg.GetHostname()

	fmt.Println("\n---Flag Library---")
	flag_pkg.GetDatabaseConfig()
}