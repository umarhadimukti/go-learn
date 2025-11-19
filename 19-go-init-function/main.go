package main

import (
	"fmt"
	"practice-go/19-go-init-function/utils"
)

func main() {
	utils.GetDatabaseName()
	isValid := true
	switch isValid {
	case false:
		fmt.Println("Database is not valid")
	case true:
		fmt.Println("Database is valid")
	default:
		fmt.Println("Unknown database status")
	}
}
