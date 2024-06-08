package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine(flag *bool, mu *sync.Mutex) {
	for {
		mu.Lock()
		if *flag {
			fmt.Println("Leaving program")
			return
		}
		mu.Unlock()

		fmt.Println("I`ma still working")
		time.Sleep(1 * time.Second)

	}
}

func main() {
	var stopFlag bool
	var mu sync.Mutex
	go goroutine(&stopFlag, &mu)

	time.Sleep(2 * time.Second)

	mu.Lock()
	stopFlag = true
	mu.Unlock()

	time.Sleep(2 * time.Second)

}
