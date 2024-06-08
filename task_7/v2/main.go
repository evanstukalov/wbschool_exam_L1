package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[int]int)
	ch := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for kv := range ch {
			myMap[kv]++
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i
		}(i)
	}

	wg.Wait()
	close(ch)
	fmt.Println(myMap)
}
