package main

import (
	"fmt"
)

func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

func main() {
	num := 8
	printType(num)

	sentence := "Hello world!"
	printType(sentence)

	isBool := true
	printType(isBool)

	ch := make(chan int)
	printType(ch)
}
