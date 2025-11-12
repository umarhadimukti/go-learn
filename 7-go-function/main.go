package main;

import (
	"fmt"
	"strings"
);

type FuncType func(string) string;
type AnonymousFuncType func(string) bool;

// Function that returns multiple values
func multipleCalculations(a int, b int) (resultA int, resultB int, resultC int) {
	resultA = a + b;
	resultB = a - b;
	resultC = a * b;
	return;
}

// Variadic function to calculate the sum of all numbers
func sumNumbers(nums ...int) int {
	total := 0;
	for _, num := range nums {
		total += num;
	}
	return total;
}

func greetingsWithFilter(name string, filterFunc FuncType) {
	fmt.Println("Hello, " + fmt.Sprint(filterFunc(name)) + "!");
}

func spamFilter(name string) string {
	switch name {
		case "anjing", "babi", "bangsat":
			return "*****";
	}
	return name;
}

func registerUser(name string, validateFunc AnonymousFuncType) {
	if !validateFunc(name) {
		fmt.Println("Invalid user name: " + name);
	} else {
		fmt.Println("User registered: " + name);
	}
}

func factorial(number int) int {
	if number == 1 {
		return 1;
	}
	return number * factorial(number - 1);
}

func main() {
	sum, difference, multiply := multipleCalculations(20, 10);
	fmt.Println("20 + 10 = " + fmt.Sprint(sum));
	fmt.Println("20 - 10 = " + fmt.Sprint(difference));
	fmt.Println("20 * 10 = " + fmt.Sprint(multiply));
	fmt.Println("======================");

	fmt.Println("Sum of numbers: " + fmt.Sprint(sumNumbers(1, 2, 3, 4, 5)));
	numbersSlice := []int{10, 20, 30, 40, 50};
	fmt.Println("Sum of slice numbers: " + fmt.Sprint(sumNumbers(numbersSlice...)));
	fmt.Println("======================");

	fmt.Println("Function as variable");
	numberA := sumNumbers;
	result := numberA(5, 10, 15);
	fmt.Println("Sum of 5, 10, 15 = " + fmt.Sprint(result));
	fmt.Println("======================");

	fmt.Println("Function as parameter");
	greetingsWithFilter(strings.ToLower("Umar"), spamFilter);
	greetingsWithFilter(strings.ToLower("AnJIng"), spamFilter);
	fmt.Println("======================");

	fmt.Println("Anonymous Function");
	validated := func (name string) bool {
		if len(name) < 3 {
			return false;
		}
		return true;
	}
	registerUser("Umar", validated);
	registerUser("Jo", func(name string) bool {
		if len(name) < 3 {
			return false;
		}
		return true;
	});
	fmt.Println("======================");

	fmt.Println("Recursive Function");
	factNumber := 5;
	fmt.Println("Factorial of " + fmt.Sprint(factNumber) + " is " + fmt.Sprint(factorial(factNumber)));
}