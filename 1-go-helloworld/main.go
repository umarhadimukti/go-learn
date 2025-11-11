package main;

import "fmt";

func greetings(name string) string {
	return "Hello, " + name + "!";
}

func main() {
	greetUser1 := greetings("Alice");
	greetUser2 := greetings("John Doe");
	fmt.Println("User 1: " + greetUser1 + "\nLength: " + fmt.Sprint(len(greetUser1)));
	fmt.Println(greetUser2);
}