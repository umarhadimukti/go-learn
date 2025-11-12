package main;

import "fmt";

type HasPosition interface {
	GetPosition() string;
	SetPosition(position string);
}

type Employee struct {
	Name, Position, Country string;
	Age int;
}

func (employee Employee) GetPosition() string {
	return fmt.Sprint(employee.Name) + " is a " + fmt.Sprint(employee.Position);
}

func (employee Employee) SetPosition(position string) {
	employee.Position = position;
	Greeting(employee);
}

func Greeting(employee HasPosition) {
	fmt.Println("Greetings! " + employee.GetPosition());
}

func main() {
	firstEmployee := Employee{"Umar", "Software Engineer", "Indonesia", 14};
	Greeting(firstEmployee);
	fmt.Println("=======================================");
	firstEmployee.SetPosition("Senior Software Engineer");
}