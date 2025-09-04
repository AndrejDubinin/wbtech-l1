package main

import "fmt"

func main() {
	a := 10
	b := 25

	fmt.Printf("before a: %d, b: %d\n", a, b)

	// a, b = b, a
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("after a: %d, b: %d\n", a, b)
}
