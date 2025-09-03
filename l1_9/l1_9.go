package main

import (
	"fmt"
	"sync"
)

func producer(wg *sync.WaitGroup, source []int) <-chan int {
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)

		for _, val := range source {
			ch <- val
		}
	}()

	return ch
}

func multiplier(wg *sync.WaitGroup, in <-chan int) <-chan int {
	out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)

		for i := range in {
			out <- i * 2
		}
	}()

	return out
}

func main() {
	wg := &sync.WaitGroup{}
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	ch := producer(wg, source)
	ch = multiplier(wg, ch)

	for i := range ch {
		fmt.Println(i)
	}

	wg.Wait()
}
