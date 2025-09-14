package main

import (
	"fmt"
	"sync"
	"time"
)

const sleepDuration = 5 * time.Second

func sleep(duration time.Duration) {
	ch := make(chan struct{})

	time.AfterFunc(duration, func() {
		defer close(ch)
		ch <- struct{}{}
	})

	<-ch
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("time:", time.Now())
		sleep(sleepDuration)
		fmt.Println("time after sleep:", time.Now())
	}()

	wg.Wait()
}
