package main;

import "fmt";

func main() {
	product := map[string]string{
		"name": "Laptop Asus TUF A15",
		"brand": "Asus",
		"price": "Rp. 15.000.000",
		"color": "Black",
	};

	fmt.Println("Product Details:");
	for key, value := range product {
		fmt.Println(key + ": " + value);
	}
	fmt.Println("======================");
	fmt.Println("Length of product map: " + fmt.Sprint(len(product)));
	fmt.Println("======================");
	delete(product, "color");
	fmt.Println("Product Details after deleting 'color':");
	for key, value := range product {
		fmt.Println(key + ": " + value);
	}
	fmt.Println("======================");

	ages := map[string]int{
		"Alice": 30,
		"Bob": 25,
		"Charlie": 35,
	};

	newAges := make(map[string]int);
	for name, age := range ages {
		newAges[name] = age;
	}
	fmt.Println("New Ages: ", fmt.Sprint(newAges));
	fmt.Println("======================");

	val1, isOk1 := ages["Alice"];
	if isOk1 {
		fmt.Println("Alice's Age: " + fmt.Sprint(val1));
	} else {
		fmt.Println("Alice's age not found");
	}
	fmt.Println("=======================" );

	var grades [9]int = [9]int{90, 85, 88, 92, 95, 45, 69, 30, 55};
	var badGrades []int;
	for i := 0; i < len(grades); i++ {
		if grade := grades[i]; grade < 75 {
			badGrades = append(badGrades, grade);
		}
	}
	fmt.Println("Bad Grades: " + fmt.Sprint(badGrades));
	fmt.Println("=======================");
}