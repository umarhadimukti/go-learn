package main

import (
	"fmt"
	"github.com/umarhadimukti/go-learn/19-go-init-function/utils"
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
