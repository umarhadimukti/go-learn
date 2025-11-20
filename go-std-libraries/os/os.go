package os

import (
	"log"
	"os"
	"fmt"
)

func ReadFile(filepath string) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}

func WriteFile(filepath string, data string) {
	err := os.WriteFile(filepath, []byte(data), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadArgs() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func GetHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hostname:", hostname)
}