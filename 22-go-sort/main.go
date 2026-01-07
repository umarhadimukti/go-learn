package main

import (
	"fmt"
	"sort"
)

type Student struct {
	No int
	Name string
	Age int
}

type StudentSlice []Student

func (student StudentSlice) Len() int {
	return len(student)
}

func (student StudentSlice) Less(i, j int) bool {
	return student[i].Age < student[j].Age
}

func (student StudentSlice) Swap(i, j int) {
	student[i], student[j] = student[j], student[i]
}

func main() {
	students := StudentSlice{
		{No: 1, Name: "Alice", Age: 23},
		{No: 2, Name: "Bob", Age: 21},
		{No: 3, Name: "Charlie", Age: 22},
	}
	fmt.Println("Before sorting:", students)
	sort.Sort(StudentSlice(students))
	fmt.Println("After sorting by Age:", students)
	
	fmt.Println()
	SortLecture()
}