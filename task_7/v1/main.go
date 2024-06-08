package main

import (
	"fmt"
	"sync"
)

func main() {
	sharedMap := map[int]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(mu *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				mu.Lock()
				sharedMap[i]++
				mu.Unlock()
			}
		}(&mu, &wg)
	}

	wg.Wait()

	for k, v := range sharedMap {
		fmt.Println(k, v)
	}

}
