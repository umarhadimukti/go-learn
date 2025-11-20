package main

import (
	"fmt"
	"errors"
	errors_pkg "practice-go/go-std-libraries/errors"
	os_pkg "practice-go/go-std-libraries/os"
	flag_pkg "practice-go/go-std-libraries/flag"
	strings_pkg "practice-go/go-std-libraries/strings"
	strconv_pkg "practice-go/go-std-libraries/strconv"
	math_pkg "practice-go/go-std-libraries/math"
	container_list_pkg "practice-go/go-std-libraries/container-list"
	container_ring_pkg "practice-go/go-std-libraries/container-ring"
	time_pkg "practice-go/go-std-libraries/time"
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

	fmt.Println("\n---Strings Library---")
	strings_pkg.DemoStringsPackage()

	fmt.Println("\n---Strconv Library---")
	strconv_pkg.DemoStrconvPackage()

	fmt.Println("\n---Math Library---")
	math_pkg.DemoMathPackage()

	fmt.Println("\n---Container lists Library---")
	container_list_pkg.DemoLinkedList()

	fmt.Println("\n---Container ring library---")
	container_ring_pkg.DemoCircularList()

	fmt.Println("\n---Time Library---")
	time_pkg.DemoTimePackage()
}