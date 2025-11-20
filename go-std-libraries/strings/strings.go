package strings

import (
	"strings"
	"fmt"
)

func DemoStringsPackage() {
	sample := "   Hello, Ladies and Gentlemen, Welcome to the Hotel California! "

	fmt.Println("Original String:", sample)

	trimmed := strings.TrimSpace(sample)
	fmt.Println("Trimmed String:", trimmed)

	upper := strings.ToUpper(trimmed)
	fmt.Println("Uppercase String:", upper)

	lower := strings.ToLower(trimmed)
	fmt.Println("Lowercase String:", lower)

	contains := strings.Contains(trimmed, "California")
	fmt.Println("Contains 'California':", contains)

	split := strings.Split(trimmed, ",")
	fmt.Println("Split String:", split)

	joined := strings.Join(split, " ")
	fmt.Println("Joined String:", joined)

	replaced := strings.ReplaceAll(trimmed, "Hotel", "Motel")
	fmt.Println("Replaced String:", replaced)

	index := strings.Index(trimmed, "Ladies")
	fmt.Println("Index of 'Ladies':", index)

}