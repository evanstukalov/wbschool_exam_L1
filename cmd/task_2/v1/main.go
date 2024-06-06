package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num * num)
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg)
	}

	wg.Wait()
}
