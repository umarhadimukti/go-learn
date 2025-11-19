package main

import (
	"fmt"
)

type ValidationError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (validationError *ValidationError) Error() string {
	return validationError.Message
}

func (notfoundError *NotFoundError) Error() string {
	return notfoundError.Message
}

func SaveData(id int, data any) (string, error) {
	if id != 3 {
		return "", &NotFoundError{Message: "record not found"}
	}
	if data == nil {
		return "", &ValidationError{Message: "data cannot be empty"}
	}
	return "data saved successfully", nil
}

func main() {
	result, err := SaveData(3, nil)
	if err != nil {
		// data validation error
		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *NotFoundError:
			fmt.Println("Not Found Error:", finalError.Error())
		default:
			fmt.Println("Unknown Error:", finalError.Error())
		}
	} else {
		// data saved successfully
		fmt.Println(result)
	}

}