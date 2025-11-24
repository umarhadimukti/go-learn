package main

import (
	"fmt"
	"os"
)

func createNewFile(msg string) {
	file, err := os.OpenFile("23-go-file-manipulation/data.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(msg)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File created and data written successfully.")
}

func main() {
	dataString := "Hello, welcome to the Hotel California.\nSuch a lovely place"
	createNewFile(dataString)
}