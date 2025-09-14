package main

import (
	"fmt"
	"strings"
)

func unique(s string) bool {
	m := make(map[rune]bool)
	s = strings.ToLower(s)

	for _, r := range s {
		if m[r] {
			return false
		}
		m[r] = true
	}
	return true
}

func main() {
	s := "abcd"
	fmt.Printf("%s, result: %t\n", s, unique(s))

	s = "abCdefAaf"
	fmt.Printf("%s, result: %t\n", s, unique(s))

	s = "aabcd"
	fmt.Printf("%s, result: %t\n", s, unique(s))
}
