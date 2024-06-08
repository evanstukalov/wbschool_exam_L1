package main

import (
	"fmt"
)

func main() {
	numbers := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstCh := make(chan int)
	secondCh := make(chan int)

	go func(ch chan<- int, numbers [9]int) {
		for num := range numbers {
			firstCh <- num
		}
		close(ch)
	}(firstCh, numbers)

	go func(firstCh <-chan int, secondCh chan<- int) {
		for num := range firstCh {
			secondCh <- num * num
		}
		close(secondCh)
	}(firstCh, secondCh)

	for num := range secondCh {
		fmt.Println(num)
	}

}
