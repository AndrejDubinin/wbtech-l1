package main

import (
	"fmt"
	"sort"
)

func main() {
	setOne := []int{1, 2, 3, 7, 4, 9}
	setTwo := []int{2, 3, 4, 1, 8, 5, 7}

	sort.Ints(setOne)
	sort.Ints(setTwo)

	var i, j int
	result := make([]int, 0)

	for i < len(setOne) && j < len(setTwo) {
		if setOne[i] == setTwo[j] {
			result = append(result, setOne[i])
			i++
			j++
		} else if setOne[i] < setTwo[j] {
			i++
		} else {
			j++
		}
	}

	fmt.Printf("%v\n", result)
}
