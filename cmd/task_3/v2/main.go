package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, mu *sync.Mutex, sum *int) {
	pow := num * num
	mu.Lock()
	*sum += pow
	mu.Unlock()
	wg.Done()
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var mu sync.Mutex
	var sum int

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, &mu, &sum)
	}

	wg.Wait()

	fmt.Println("Sum:", sum)
}
