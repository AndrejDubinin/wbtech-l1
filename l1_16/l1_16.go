package main

import (
	"fmt"
	"math/rand"
)

func getPivot(nums []int, l, r int) (int, int) {
	pivotIdx := l
	maxIdx := r - l - 1

	if maxIdx > 0 {
		pivotIdx = rand.Intn(maxIdx) + l
	}

	return nums[pivotIdx], pivotIdx
}

func partition(nums []int, lIdx, rIdx int) int {
	pivot, pivotIdx := getPivot(nums, lIdx, rIdx)

	nums[pivotIdx], nums[lIdx] = nums[lIdx], nums[pivotIdx]
	l, r := lIdx+1, rIdx-1

	for l <= r {
		if nums[l] < pivot {
			l++
		} else if nums[r] > pivot {
			r--
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	nums[lIdx], nums[r] = nums[r], nums[lIdx]
	return r
}

func qsort(nums []int, l, r int) []int {
	if l >= r {
		return nums
	}

	pivotIdx := partition(nums, l, r)

	qsort(nums, l, pivotIdx)
	qsort(nums, pivotIdx+1, r)
	return nums
}

func quickSort(nums []int) []int {
	return qsort(nums, 0, len(nums))
}

func main() {
	arr := []int{1, 13, 16, 5, 12, 3, 17, 18, 10, 19, 7, 2, 9, 8, 11, 14, 15, 4, 6, 20}
	fmt.Printf("array before: %v\n", arr)

	arr = quickSort(arr)
	fmt.Printf("array after: %v\n", arr)
}
