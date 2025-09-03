package main

import (
	"fmt"
	"strconv"
	"sync"
)

const workersAmount = 10

type CustomMap struct {
	data map[string]string
	mx   sync.RWMutex
}

func (m *CustomMap) Put(key, value string) {
	m.mx.Lock()
	defer m.mx.Unlock()

	m.data[key] = value
}

func (m *CustomMap) Get(key string) (string, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()

	if value, inMap := m.data[key]; inMap {
		return value, true
	}
	return "", false
}

func NewCustomMap() *CustomMap {
	return &CustomMap{
		data: make(map[string]string),
	}
}

func main() {
	customMap := NewCustomMap()
	wg := &sync.WaitGroup{}

	for i := range workersAmount {
		wg.Add(1)
		go func() {
			defer wg.Done()

			val := strconv.Itoa(i)
			customMap.Put(fmt.Sprintf("key-%s", val), val)
		}()
	}

	for i := range workersAmount {
		wg.Add(1)
		go func() {
			defer wg.Done()

			key := fmt.Sprintf("key-%d", i)
			val, inMap := customMap.Get(key)
			if !inMap {
				fmt.Printf("no value for key: %s\n", key)
				return
			}
			fmt.Println("value:", val)
		}()
	}

	wg.Wait()
}
