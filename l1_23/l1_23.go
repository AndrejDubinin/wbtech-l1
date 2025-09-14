package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	i := 2

	copy(arr[i:], arr[i+1:])
	arr = arr[:len(arr)-1]
	fmt.Printf("%v\n", arr)
}
