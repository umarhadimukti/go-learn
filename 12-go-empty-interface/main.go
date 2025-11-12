package main;

import "fmt";

// Empty interface can hold values of any types
func emptyInterfaceDemo(value interface{}) {
	fmt.Println("Value:", value, ", Type:", fmt.Sprintf("%T", value));
}

func secondEmptyInterfaceDemo(arg1 any) any {
	return arg1;
}

func main() {
	emptyInterfaceDemo(50);
	emptyInterfaceDemo("Hello, Go!");
	fmt.Println("=================================");

	result := secondEmptyInterfaceDemo(3.14);
	fmt.Println("Returned value:", result, ", Type:", fmt.Sprintf("%T", result));
}