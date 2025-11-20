package strconv

import (
	"strconv"
	"fmt"
)

func DemoStrconvPackage() {
	var port string = "8000"
	portInt, err := strconv.Atoi(port) // convert string to int
	if err != nil {
		fmt.Println("Error converting port:", err)
	} else {
		fmt.Println("Port as integer:", portInt + 80)
	}

	var emptySpace int = 500
	emptySpaceInt := strconv.Itoa(emptySpace) // convert int to string
	fmt.Println("Empty space as string:", emptySpaceInt + "MB")


	var price float64 = 155000
	priceStr := strconv.FormatFloat(price, 'f', 2, 64) // convert float to string
	fmt.Println("Price as string:", priceStr)

	var isAvailable string = "false"
	isAvailableBool, err := strconv.ParseBool(isAvailable) // convert string to bool
	if err != nil {
		fmt.Println("Error converting availability:", err)
	} else {
		if isAvailableBool {
			fmt.Println("The item is available")
		} else {
			fmt.Println("The item is not available")
		}
	}

	var negativeNumber string = "-30"
	negativeNumberInt, err := strconv.ParseInt(negativeNumber, 10, 16) // convert negative string to int16
	if err != nil {
		fmt.Println("Error converting to int16:", err.Error())
	} else {
		fmt.Println("Converted int16 value:", negativeNumberInt, "Type:", fmt.Sprintf("%T", negativeNumberInt))
	}
	
}