package main

import (
	"fmt"
)

type Product struct {
	Name, Category string
	Price string
}

func main() {
	var product1 *Product = new(Product) // create empty pointer using 'new' keyword
	var product2 *Product = product1
	product2 = &Product{"Acer Predator", "Laptop", "Rp 20.000.000"}
	fmt.Println(*product1)
	fmt.Println(*product2)

	x := new(int)
	fmt.Println("Value of x:", *x)
}