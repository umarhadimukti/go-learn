package utils

import (
	"fmt"
)

var databaseName string

func init() {
	databaseName = "db-hotel-california"
}

func GetDatabaseName() {
	fmt.Println("Database Name:", databaseName)
}