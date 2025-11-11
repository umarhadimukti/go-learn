package main;

import "fmt";

func sumTwoNumbers(x int, y int) int {
	return x + y;
}

func multiplyNumber(num int, multiplier int) int {
	return num * multiplier;
}

func main() {
	sum := sumTwoNumbers(10, 5);
	fmt.Println("10 + 5\nResult: " + fmt.Sprint(sum));
	fmt.Println("======================");

	var numberA int;
	numberA = multiplyNumber(10, 5);
	fmt.Println("10 * 5\nResult: " + fmt.Sprint(numberA));
	fmt.Println("======================");

	const (
		firstName = "Umar"
		middleName = "Hadi"
		lastName = "Mukti"
	);
	fullName := firstName + " " + middleName + " " + lastName;
	fmt.Println("Full Name: " + fullName);
	fmt.Println("======================");

	const pi float32 = 3.14;
	var radius float32 = 7;
	area := pi * radius * radius;
	fmt.Println("Area of Circle with radius 7: " + fmt.Sprint(area));
	fmt.Println("======================");
}