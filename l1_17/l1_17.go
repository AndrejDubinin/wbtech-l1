package main

import "fmt"

const notFound = -1

func search(arr []int, target int) int {
	l, r := 0, len(arr)

	for (r - l) > 1 {
		m := (l + r) / 2

		if arr[m] <= target {
			l = m
		} else {
			r = m
		}
	}

	if arr[l] == target {
		return l
	}

	return notFound
}

func main() {
	arr := []int{-5, -4, -3, -1, 0, 1, 2, 3, 4, 5, 8, 9, 11, 13, 14, 15, 17, 18, 19, 20}

	fmt.Println("targer exists:", search(arr, 5))
	fmt.Println("targer not exists:", search(arr, 10))
	fmt.Println("targer is first element:", search(arr, -5))
	fmt.Println("targer is last element:", search(arr, 20))
}
