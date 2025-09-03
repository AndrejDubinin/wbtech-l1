package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x int

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("parameter: %v\n", err)
	}

	fmt.Printf("number: %b\n", x)

	x = x &^ 1

	fmt.Printf("after clear bit: %b\n", x)
}
