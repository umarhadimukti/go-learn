package slices

import (
	"slices"
	"fmt"
)

func slicesContains(number int) {
	a := []int{5, 3, 1, 6, 8, 4}
	existsNumber := slices.Contains(a, number)
	status := "doesn't exists"
	if existsNumber {
		status = "exists"
	}
	fmt.Println("Number", number, "is", status, "in slice", a)
}

func DemoSlicesPackage() {
	slicesContains(10)

	var a = []int{2, 11, 20, 10, 8, 19, 31}
	fmt.Println("Minimum value in", a, "is", slices.Min(a))
	fmt.Println("Maximum value in", a, "is", slices.Max(a))
	fmt.Println("Index of 10 in", a, "is", slices.Index(a, 10))

	var students = []string{"Alice", "Bob", "Vera", "Yuni", "Dela"}
	slices.Sort(students)
	n, found := slices.BinarySearch(students, "Dela")
	fmt.Println(students, "found Dela at index", n, "found:", found)
}