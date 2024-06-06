package main

import "fmt"

func square(num int, ch chan<- int) {
	ch <- num * num
}

func main() {
	intCh := make(chan int)
	numbers := [5]int{2, 4, 6, 8, 10}

	for _, num := range numbers {
		go square(num, intCh)
	}

	for range numbers {
		fmt.Println(<-intCh)
	}
}
