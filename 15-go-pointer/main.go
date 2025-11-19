package main

import (
	"fmt"
)

type User struct {
	ID, Name, Email, Address string
}

func main() {
	user1 := User{"A1", "Umar", "umar@example.com", "123 Street"}
	user2 := user1 // pass by value (default) || copy user1 value to user2
	user2.Address = "456 Avenue"
	fmt.Println("--- Pass by value ---")
	fmt.Println(user1)
	fmt.Println(user2)

	var user3 *User = &user1 // pass by reference || user3 is pointer to user1
	user3.Address = "789 Boulevard"
	fmt.Println("\n--- Pass by reference ---")
	fmt.Println(user1)
	fmt.Println(user3)

	*user3 = User{"B3", "Aisha", "aisha@example.com", "321` Road"} // update value by asterisk
	fmt.Println("\n--- Update value (by asterisk) ---")
	fmt.Println(user1)
	fmt.Println(user3)
}