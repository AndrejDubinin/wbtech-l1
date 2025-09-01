package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func worker(wg *sync.WaitGroup, i int, in <-chan int) {
	defer wg.Done()

	for num := range in {
		fmt.Printf("%d worker: %d\n", i, num)
	}

	fmt.Printf("%d worker: quit\n", i)
}

func producer(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		num := 1

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

func main() {
	var numWorkers int
	var err error
	wg := &sync.WaitGroup{}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	numWorkers, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("parameter: %v\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := producer(ctx)

	for i := range numWorkers {
		wg.Add(1)
		go worker(wg, i, ch)
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Println("Shutdown ...")
		cancel()
	}()

	wg.Wait()

	fmt.Println("exiting")
}
