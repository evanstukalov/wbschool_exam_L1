package main

import (
	"fmt"
	"time"
)

func goroutine(done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Leaving program")
			return
		default:
			fmt.Println("I`ma still working")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	done := make(chan struct{})

	go goroutine(done)

	time.Sleep(2 * time.Second)

	close(done)

	time.Sleep(2 * time.Second)

}
