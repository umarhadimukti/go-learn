package main

import (
	"fmt"
)

type Product struct {
	Name, Category string
	Price float32
}

func (prod *Product) changeName(newName string) {
	prod.Name = newName
}

func main() {
	var product1 Product = Product{"Acer Predator", "Laptop", 20000000}
	product1.changeName("Macbook Air M4")
	fmt.Println("--- Product 1 ---")
	fmt.Println("Product Name:", product1.Name)
	fmt.Println("Category:", product1.Category)
	fmt.Println("Price:", "Rp" + fmt.Sprint(int(product1.Price)))
}