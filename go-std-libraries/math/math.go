package math

import (
	"math"
	"fmt"
)

func DemoMathPackage() {
	const (
		numberA float64 = 3.14
		numberB int = 16
	)

	fmt.Println("Square root of", numberA, "is", math.Sqrt(numberA)) // square root
	fmt.Println("Square root of", numberB, "is", math.Sqrt(float64(numberB))) // square root

	fmt.Println("Power of", fmt.Sprint(numberA) + "^2 is", math.Pow(numberA, 2)) // power

	fmt.Println("Round of", numberA, "is", math.Round(numberA)) // round to nearest integer

	fmt.Println(math.IsNaN(4.2)) // check if value is NaN
}