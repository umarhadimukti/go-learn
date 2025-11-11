package main;

import "fmt";

func main() {
	var months = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"};

	var firstSliceOfMonths = months[1:4];
	fmt.Println("Months from index 1 to 3: " + fmt.Sprint(firstSliceOfMonths));
	fmt.Println("=========================");

	var secondSliceOfMonths = months[:7];
	fmt.Println("Months from index 0 to 6: " + fmt.Sprint(secondSliceOfMonths));
	fmt.Println("=========================");

	var thirdSliceOfMonths = months[:];
	fmt.Println("All Months: " + fmt.Sprint(thirdSliceOfMonths));
	fmt.Println("=========================");

	fmt.Printf("Type of months variable: %T\n", months);
	fmt.Printf("Type of firstSliceOfMonths variable: %T\n", firstSliceOfMonths)
	fmt.Println("=========================");

	days := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"};
	var firstSliceOfDays = days[5:];
	firstSliceOfDays[0] = "Funday"; // modify the slice
	fmt.Println("Modified Slice of Days: " + fmt.Sprint(firstSliceOfDays));
	fmt.Println("Original Days Array after modifying the slice: " + fmt.Sprint(days));
	fmt.Println("=========================");

	secondSliceOfDays := append(firstSliceOfDays, "Newday");
	secondSliceOfDays[1] = "NewdayModified";
	fmt.Println("Second Slice of Days after appending: " + fmt.Sprint(secondSliceOfDays));
	fmt.Println("Original Days Array after modifying the second slice: " + fmt.Sprint(days));
	fmt.Println("=========================");

	fromSlice := days[:];
	toSlice := make([]string, len(fromSlice), cap(fromSlice));
	copy(toSlice, fromSlice);
	toSlice[0] = "ChangedDay";
	fmt.Println("From Slice: " + fmt.Sprint(fromSlice));
	fmt.Println("To slice after copy and modification: " + fmt.Sprint(toSlice));
	fmt.Println("fromSlice after toSlice modification: " + fmt.Sprint(fromSlice));
}