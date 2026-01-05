package main;

import (
	"fmt"
	"os"
)

func logging() {
	message := recover();
	if message != nil {
		fmt.Println("Error message: " + fmt.Sprint(message));
		fmt.Println("Application error handled");
	}
	fmt.Println("log: running function");
}

func runApplication(isError bool) {
	defer logging();
	if isError {
		panic("Application Error");
	}
}

func runProcess(path string) (err error) {
	if path == "" {
		panic("Empty path");
	}
	file, err := os.Open(path);
	if err != nil {
		return;
	}
	defer file.Close();
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic:%v", r);
		}
	}()
	return nil;
}

func main() {
	fmt.Println("Run application without error");
	runApplication(false);
	fmt.Println("======================");
	
	fmt.Println("Run application with error");
	runApplication(true);
	fmt.Println("======================");

	fmt.Println("Run process");
	err := runProcess("9-go-defer-panic-recover/test.txt");
	if err != nil {
		fmt.Println("Error here:", err);
	}
	fmt.Println("File processed successfully");
	fmt.Println("=======================");
}