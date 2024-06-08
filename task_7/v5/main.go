package main

import (
	"fmt"
	"sync"
)

type SharedMap struct {
	mu sync.Mutex
	m  map[int]int
}

func (m *SharedMap) Inc(key int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.m[key]++

}

func main() {
	var sharedMap = SharedMap{m: make(map[int]int)}
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(mu *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				mu.Lock()
				sharedMap.Inc(i)
				mu.Unlock()
			}
		}(&mu, &wg)
	}

	wg.Wait()

	for k, v := range sharedMap.m {
		fmt.Println(k, v)
	}

}
