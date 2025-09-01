package main

import (
	"fmt"
	"time"
)

const secondsBeforeQuit = 5

func producer(timeCh <-chan time.Time) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		var num int

		for {
			select {
			case <-timeCh:
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

func printer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("printer: %d\n", num)
	}

	fmt.Println("printer: quit")
}

func main() {
	timeCh := time.After(secondsBeforeQuit * time.Second)
	ch := producer(timeCh)
	printer(ch)
}
