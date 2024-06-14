package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	ch      chan struct{}
}

func (c *Counter) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	c.ch <- struct{}{}
	c.counter++
	<-c.ch
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{ch: make(chan struct{}, 1)}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go counter.inc(&wg)
	}

	wg.Wait()

	fmt.Println(counter.counter)

}
