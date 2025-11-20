package flag

import (
	"flag"
	"fmt"
)

func GetDatabaseConfig() {
	host := flag.String("h", "localhost", "Database host")
	dbname := flag.String("d", "db-example", "Database name")
	port := flag.Int("port", 5432, "Database port")
	user := flag.String("u", "root", "Database user")
	password := flag.String("p", "root", "Database password")

	flag.Parse()

	fmt.Println("Database Configuration:")
	fmt.Println("Host:", *host)
	fmt.Println("Database Name:", *dbname)
	fmt.Println("Port:", *port)
	fmt.Println("User:", *user)
	fmt.Println("Password:", *password)
}