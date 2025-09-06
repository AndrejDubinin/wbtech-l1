package main

import (
	"bufio"
	"fmt"
	"os"
)

const minStringLength = 2

func reverseString(s string) string {
	if len(s) < minStringLength {
		return s
	}

	data := []rune(s)
	l, r := 0, len(data)-1

	for l < r {
		data[l], data[r] = data[r], data[l]
		l++
		r--
	}

	return string(data)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Println("Print string:")
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("reading string: %v\n", err)
	}

	fmt.Println(reverseString(str))
}
