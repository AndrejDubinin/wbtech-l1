package main

import (
	"math/rand"
)

var justString string

func createHugeString(size int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, size)
	for i := range size {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
}
