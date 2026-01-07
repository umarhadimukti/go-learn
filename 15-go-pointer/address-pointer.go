/**
* To run this file, follow this command:
* cd 15-go-pointer
* go run .
*/

package main

import (
	"fmt"
)

func PracticeAddress (user *User) {
	fmt.Println("Before update user:")
	fmt.Println("User ID:", user.ID, "\nName:", user.Name, "\nEmail:", user.Email)

	*user = User{"111", "Nabila", "nabila@gmail.com", "Bandung"}

	fmt.Println("After update user:")
	fmt.Println("User ID:", user.ID, "\nName:", user.Name, "\nEmail:", user.Email)
}