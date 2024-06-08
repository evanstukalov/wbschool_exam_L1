package main

import (
	"fmt"
	"sync"
)

func main() {
	var myMap sync.Map
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			myMap.Store(i, i*i)
		}(i)
	}

	wg.Wait()

	myMap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
