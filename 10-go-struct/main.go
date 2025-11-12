package main;

import "fmt";

type Employee struct {
	Name, Position, Country string;
	age int;
}

func (employee Employee) getInfo() string {
	return "Name: " + employee.Name + ", Position: " + employee.Position + ", Country: " + employee.Country + ", Age: " + fmt.Sprint(employee.age);
}

func main() {
	firstEmployee := Employee{"Umar", "Software Engineer", "Indonesia", 14};
	fmt.Println("First Employee");
	fmt.Println(firstEmployee.getInfo());
	fmt.Println("===========================");

	var secondEmployee Employee;
	secondEmployee.Name = "John McClane";
	secondEmployee.Position = "Backend Engineer";
	secondEmployee.Country = "USA";
	secondEmployee.age = 35;
	fmt.Println("Second Employee");
	fmt.Println(secondEmployee.getInfo());
	fmt.Println("===========================");

	thirdEmployee := Employee{
		Name:    "Jane Doe",
		Position: "Frontend Engineer",
		Country:  "UK",
		age:      28,
	};
	fmt.Println("Third Employee");
	fmt.Println(thirdEmployee.getInfo());
}
