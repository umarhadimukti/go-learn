package main;

import "fmt";

func main() {
	counter := 10;
 
	decrement := func() {
		counter -= 2;
		fmt.Println("Counter inside closure: " + fmt.Sprint(counter));
	};

	fmt.Println("Counter before calling closure: " + fmt.Sprint(counter));
	decrement();
	fmt.Println("Counter after calling closure: " + fmt.Sprint(counter));
}