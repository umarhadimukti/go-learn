package main;

import "fmt";

func printFirstTriangle(rows int) string {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ");
		}
		fmt.Println();
	}
	return "";
}

func printSecondTriangle(rows int) string {
	for i := rows; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ");
		}
		fmt.Println();
	}
	return "";
}

func printThirdTriangle(rows int) string {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows - i; j++ {
			fmt.Print(" "); // print spaces
		}
		for k := 1; k <= i; k++ {
			fmt.Print("* "); // print stars
		}
		fmt.Println();
	}
	return "";
}

func main() {
	numberOfRows := 10;
	printFirstTriangle(numberOfRows);
	fmt.Println("======================");
	printSecondTriangle(numberOfRows);
	fmt.Println("======================");
	printThirdTriangle(numberOfRows);
}