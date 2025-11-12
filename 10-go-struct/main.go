package main;

import "fmt";

type Employee struct {
	Name, Position, Country string;
	age int;
}

func main() {
	firstEmployee := Employee{"Umar", "Software Engineer", "Indonesia", 14};
	fmt.Println("First Employee");
	fmt.Println("Name: " + firstEmployee.Name);
	fmt.Println("Position: " + firstEmployee.Position);
	fmt.Println("Country: " + firstEmployee.Country);
	fmt.Println("Age: " + fmt.Sprint(firstEmployee.age));
	fmt.Println("===========================");

	var secondEmployee Employee;
	secondEmployee.Name = "John McClane";
	secondEmployee.Position = "Backend Engineer";
	secondEmployee.Country = "USA";
	secondEmployee.age = 35;
	fmt.Println("Second Employee");
	fmt.Println("Name: " + secondEmployee.Name);
	fmt.Println("Position: " + secondEmployee.Position);
	fmt.Println("Country: " + secondEmployee.Country);
	fmt.Println("Age: " + fmt.Sprint(secondEmployee.age));
	fmt.Println("===========================");

	thirdEmployee := Employee{
		Name:    "Jane Doe",
		Position: "Frontend Engineer",
		Country:  "UK",
		age:      28,
	};
	fmt.Println("Third Employee");
	fmt.Println("Name: " + thirdEmployee.Name);
	fmt.Println("Position: " + thirdEmployee.Position);
	fmt.Println("Country: " + thirdEmployee.Country);
	fmt.Println("Age: " + fmt.Sprint(thirdEmployee.age));
	fmt.Println("===========================");
}
