package main;

import "fmt";

func main() {
	var employees = [...]string{"Umar", "John", "Jane"};

	fmt.Println("Employee List:");
	for i := 0; i < len(employees); i++ {
		fmt.Println("- " + employees[i]);
	}
	fmt.Println("Total Employee: " + fmt.Sprint(len(employees)));
	fmt.Println("======================");

	var pantsSizes = [...]int{28, 30, 32, 34, 36};
	fmt.Println("Available Pants Sizes:")
	for _, size := range pantsSizes {
		fmt.Println("- Size " + fmt.Sprint(size));
	}
	fmt.Println("Total Sizes: " + fmt.Sprint(len(pantsSizes)));
	fmt.Println("======================");

	var messages = [...]string{"Halo ini siapa", "Lagi dimana", "Sama siapa"};
	fmt.Println("Chat list:");
	for i, message := range messages {
		fmt.Println(fmt.Sprint(i+1, "."), message);
	}
	fmt.Println("Total length of chats:", len(messages));
	fmt.Println("=====================");
}