package main

import "fmt"

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float32)

	for _, val := range temperatures {
		key := (int(val) / 10) * 10
		if _, inMap := result[key]; !inMap {
			result[key] = make([]float32, 0, 1)
		}

		result[key] = append(result[key], val)
	}

	fmt.Printf("%v\n", result)
}
