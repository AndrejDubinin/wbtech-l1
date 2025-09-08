package main

import (
	"fmt"
	"regexp"
	"strings"
)

func reverseBytes(arr []byte, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func reverseWords(s string) string {
	reCombined := regexp.MustCompile(`\s{2,}`)
	s = reCombined.ReplaceAllString(s, " ")
	arr := []byte(strings.TrimSpace(s))
	l, r := 0, len(arr)-1

	reverseBytes(arr, l, r)

	start := 0
	for i := range len(arr) {
		if arr[i] == ' ' || i == (len(arr)-1) {
			end := i - 1
			if i == (len(arr) - 1) {
				end = i
			}

			reverseBytes(arr, start, end)
			start = i + 1
		}
	}

	return string(arr)
}

func main() {
	s := "snow dog sun"
	fmt.Printf("origin: %s\n", s)
	fmt.Printf("result: %s\n", reverseWords(s))
}
