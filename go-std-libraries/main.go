package main

import (
	"fmt"
	"errors"
	errors_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/errors"
	os_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/os"
	flag_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/flag"
	strings_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/strings"
	strconv_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/strconv"
	math_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/math"
	container_list_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/container-list"
	container_ring_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/container-ring"
	time_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/time"
	reflect_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/reflect"
	regexp_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/regexp"
	encoding_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/encoding"
	slices_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/slices"
	path_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/path"
	bufio_pkg "github.com/umarhadimukti/go-learn/go-std-libraries/bufio"
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

	fmt.Println("\n---Reflect Library---")
	reflect_pkg.DemoReflectPackage()

	fmt.Println("\n---Regexp Library---")
	regexp_pkg.DemoRegexpPackage()

	fmt.Println("\n---Encoding Library---")
	encoding_pkg.DemoEncodingPackage()

	fmt.Println("\n---Slices Library---")
	slices_pkg.DemoSlicesPackage()

	fmt.Println("\n---Path Library---")
	path_pkg.DemoPathPackage()

	fmt.Println("\n---Bufio Library---")
	bufio_pkg.DemoBufioPackage()
}