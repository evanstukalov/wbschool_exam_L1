package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	mu      sync.Mutex
}

func (c *Counter) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counter++
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go counter.inc(&wg)
	}

	wg.Wait()

	fmt.Println(counter.counter)

}
