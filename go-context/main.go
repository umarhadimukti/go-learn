package main

import (
	"fmt"
	"context"
)

func contextWithValueTest() {
	contextA := context.Background() // by default is empty
	
	contextB := context.WithValue(contextA, "b", "Bajigur")
	contextC := context.WithValue(contextA, "c", "Cuanki")

	contextD := context.WithValue(contextB, "d", "Dadar")
	contextE := context.WithValue(contextC, "e", "Endok")
	contextF := context.WithValue(contextD, "f", "Fufu")

	fmt.Println("context A:", contextA)
	fmt.Println("context B:", contextB)
	fmt.Println("context C:", contextC)
	fmt.Println("context D:", contextD)
	fmt.Println("context E:", contextE)
	fmt.Println("context F:", contextF)
	
	fmt.Println("====== Get value from parent ======")

	fmt.Println("Context F Hierarki:", contextF.Value("f"))
	fmt.Println("Context F Hierarki:", contextF.Value("d"))
	fmt.Println("Context F Hierarki:", contextF.Value("b"))
}

func main() {
	contextWithValueTest()
}