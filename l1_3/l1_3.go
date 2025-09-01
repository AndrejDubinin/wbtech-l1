package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(i int, in <-chan int) {
	for num := range in {
		fmt.Printf("%d worker: %d\n", i, num)
	}
}

func main() {
	var numWorkers int
	var err error
	ch := make(chan int)

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <num_workers>")
		return
	}

	numWorkers, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("parameter: %v\n", err)
	}

	for i := range numWorkers {
		go worker(i, ch)
	}

	num := 1
	for {
		ch <- num
		num++
		time.Sleep(500 * time.Millisecond)
	}
}
