package main

import (
	"fmt"
	"sync"
)

const goroutineAmount = 3

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int, len(numbers))
	res := make(chan int, len(numbers))
	wg := &sync.WaitGroup{}

	for _, num := range numbers {
		ch <- num
	}
	close(ch)

	for range goroutineAmount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range ch {
				res <- num * num
			}
		}()
	}

	wg.Wait()
	close(res)

	for n := range res {
		fmt.Println(n)
	}
}
