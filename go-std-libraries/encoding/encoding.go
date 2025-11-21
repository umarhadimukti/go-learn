package encoding

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"os"
)

func DemoEncodingPackage() {
	fmt.Println("(Base64 Encoding Demo)")
	var data []string = []string{
		"Hello, World!",
		"Go is awesome.",
		"Base64 encoding in Go.",
	}
	for _, str := range data {
		encoded := base64.StdEncoding.EncodeToString([]byte(str))
		fmt.Println("Original:", str)
		fmt.Println("Encoded:", encoded)
	}

	fmt.Println("(CSV Encoding Demo)")
	csvString := "name,age,city\nAlice,30,New York\nBob,25,Los Angeles"
	reader := csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println("Record:", record)
	}
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{"name", "age", "address"})
	writer.Write([]string{"John", "29", "321 Main St"})
	writer.Flush()
}