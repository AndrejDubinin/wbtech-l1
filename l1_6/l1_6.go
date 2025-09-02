package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	stopByConditionMaxNum = 5
	producerTimeout       = 5
)

func producer(ctx context.Context, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		num := 0

		for {
			select {
			case <-ctx.Done():
				fmt.Println("producer: quit")
				return
			case ch <- num:
				num++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	return ch
}

func stopByCondition(wg *sync.WaitGroup, in <-chan int) <-chan int {
	out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)

		for num := range in {
			if num > stopByConditionMaxNum {
				fmt.Println("stopByCondition: quit")
				return
			}

			fmt.Printf("stopByCondition: %d\n", num)
			out <- num
		}
	}()

	return out
}

func stopByQuitChannel(wg *sync.WaitGroup, quit <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)

		var n int

		for {
			select {
			case <-quit:
				fmt.Println("stopByQuitChannel: quit")
				return
			case n = <-in:
				fmt.Printf("stopByQuitChannel: %d\n", n)
				out <- n
			}
		}
	}()

	return out
}

func stopByContextCancel(ctx context.Context, wg *sync.WaitGroup, in <-chan int) <-chan int {
	out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)

		var n int

		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopByContextCancel: quit")
				return
			case n = <-in:
				fmt.Printf("stopByContextCancel: %d\n", n)
				out <- n
			}
		}
	}()

	return out
}

func stopByGoexit(wg *sync.WaitGroup, in <-chan int) <-chan int {
	out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(out)

		for num := range in {
			fmt.Printf("stopByGoexit: %d\n", num)
			out <- num

			if num == stopByConditionMaxNum {
				runtime.Goexit()
			}
		}
	}()

	return out
}

func main() {
	wg := &sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), producerTimeout*time.Second)
	defer cancel()
	ch := producer(ctx, wg)
	in := stopByCondition(wg, ch)

	quit := make(chan struct{})
	in = stopByQuitChannel(wg, quit, in)

	ctxCtxCancel, cancelCtxCancel := context.WithCancel(ctx)
	in = stopByContextCancel(ctxCtxCancel, wg, in)

	in = stopByGoexit(wg, in)

	for n := range in {
		if n == stopByConditionMaxNum {
			close(quit)
			cancelCtxCancel()
		}
	}

	wg.Wait()
}
