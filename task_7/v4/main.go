package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan struct{}, 1)
	sharedMap := map[int]int{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				ch <- struct{}{}
				sharedMap[i]++
				<-ch
			}
		}(&wg)
	}

	wg.Wait()

	for k, v := range sharedMap {
		fmt.Println(k, v)
	}

}
