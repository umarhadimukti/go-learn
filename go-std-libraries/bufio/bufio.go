package bufio

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func DemoBufioPackage() {
	input := strings.NewReader("Hello, Umar!\nWelcome to Go programming.")
	reader := bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Writting...\n")
	writer.Flush()
}