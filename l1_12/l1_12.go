package main

import "fmt"

func main() {
	items := []string{"cat", "cat", "dog", "cat", "tree"}
	result := []string{}
	seen := map[string]struct{}{}

	for _, val := range items {
		if _, inMap := seen[val]; inMap {
			continue
		}

		seen[val] = struct{}{}
		result = append(result, val)
	}

	fmt.Printf("%v\n", result)
}
