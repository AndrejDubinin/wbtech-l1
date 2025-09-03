package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	x, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("parameter: %v\n", err)
	}

	fmt.Printf("number: %d, byte: %b\n", x, x)

	x = x &^ 1

	fmt.Printf("after clear bit: %d, byte: %b\n", x, x)
}
