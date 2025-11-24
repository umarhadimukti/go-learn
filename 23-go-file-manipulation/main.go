package main

import (
	"bufio"
	"fmt"
	"io"
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

func readFile(filePath string) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var content string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		content += string(line) + "\n"
	}
	return content, nil
}

func appendContentToFile(filePath, msg string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(msg)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	dataString := "Hello, welcome to the Hotel California.\nSuch a lovely place"

	// ==== Creating and writing to a file ====
	createNewFile(dataString)

	// ==== Reading the file ====
	content, err := readFile("23-go-file-manipulation/data_manipulation.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("File Content:\n" + content)
	}

	// ==== Appending content to the file ====
	appendMsg := "\nYou can check out any time you like, but you can never leave!"
	err = appendContentToFile("23-go-file-manipulation/data_manipulation.txt", appendMsg)
	if err != nil {
		fmt.Println("Error appending to file:", err)
	} else {
		fmt.Println("Content appended successfully.")
	}
}