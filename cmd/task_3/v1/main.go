package main

import (
	"fmt"
	"sync"
)

func square(num int, ch chan<- int, wg *sync.WaitGroup) {
	ch <- num * num
	wg.Done()
}

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	intCh := make(chan int, len(numbers))
	var wg sync.WaitGroup
	var sum int

	for _, num := range numbers {
		wg.Add(1)
		go square(num, intCh, &wg)
	}

	wg.Wait()
	close(intCh)

	for result := range intCh {
		sum += result
	}

	fmt.Println("Sum:", sum)
}
