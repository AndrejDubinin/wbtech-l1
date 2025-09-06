package main

import (
	"fmt"
	"sync"
)

const goroutineAmount = 10

type counter struct {
	cnt int
	mx  sync.RWMutex
}

func (c *counter) Increment() {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.cnt++
}

func (c *counter) Get() int {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return c.cnt
}

func NewCounter() *counter {
	return &counter{
		cnt: 0,
	}
}

func main() {
	wg := &sync.WaitGroup{}
	counter := NewCounter()

	for range goroutineAmount {
		wg.Add(1)
		go func() {
			defer wg.Done()

			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("counter: %d\n", counter.Get())
}
