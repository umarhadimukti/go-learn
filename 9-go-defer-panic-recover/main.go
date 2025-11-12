package main;

import "fmt";

func logging() {
	message := recover();
	if message != nil {
		fmt.Println("Error message: " + fmt.Sprint(message));
		fmt.Println("Application error handled");
	}
	fmt.Println("log: running function");
}

func runApplication(isError bool) {
	defer logging();
	if isError {
		panic("Application Error");
	}
}

func main() {
	fmt.Println("Run application without error");
	runApplication(false);
	fmt.Println("======================");
	
	fmt.Println("Run application with error");
	runApplication(true);
}