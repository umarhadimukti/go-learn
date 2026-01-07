package utils

import (
	"fmt"
)

var databaseName string

// init function's executed every time a file is run
// init function cannot return a value
func init() {
	databaseName = "db-hotel-california"
}

func GetDatabaseName() {
	fmt.Println("Database Name:", databaseName)
}