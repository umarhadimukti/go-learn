package main;

import "fmt";

func multipleCalculations(x int, y int) (any, any) {
	sum := x + y;
	difference := x - y;

	return sum, "Difference is " + fmt.Sprint(difference);
}

func main() {
	sumResult, diffResult := multipleCalculations(10, 5);
	sumResultInt := sumResult.(int);
	diffResultString := diffResult.(string);

	fmt.Println("Sum Result:", sumResultInt, fmt.Sprintf("(Type: %T)", sumResultInt));
	fmt.Println("Difference Result:", diffResultString, fmt.Sprintf("(Type: %T)", diffResultString));
}