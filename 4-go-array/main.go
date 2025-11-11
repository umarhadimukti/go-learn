package main;

import "fmt";

func main() {
	var employees []string;
	employees = append(employees, "Umar Hadi Mukti");
	employees = append(employees, "John Doe");
	employees = append(employees, "Jane Smith");

	fmt.Println("Employee List:");
	for i := 0; i < len(employees); i++ {
		fmt.Println("- " + employees[i]);
	}
	fmt.Println("Total Employee: " + fmt.Sprint(len(employees)));
	fmt.Println("======================");

	var pantsSizes []int = []int{28, 30, 32, 34, 36};
	fmt.Println("Available Pants Sizes:")
	for _, size := range pantsSizes {
		fmt.Println("- Size " + fmt.Sprint(size));
	}
	fmt.Println("Total Sizes: " + fmt.Sprint(len(pantsSizes)));
	fmt.Println("======================");
}