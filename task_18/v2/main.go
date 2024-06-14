package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter uint64

func (c *Counter) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddUint64((*uint64)(c), 1)
}

func main() {
	var counter Counter
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go counter.inc(&wg)
	}

	wg.Wait()

	fmt.Println(counter)

}
