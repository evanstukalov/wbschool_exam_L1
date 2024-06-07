package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stopper <-chan time.Time, ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-stopper:
			fmt.Println("Exiting producer...")
			wg.Done()
			return
		case ch <- i:
			time.Sleep(3 * time.Second)
		}
	}
}

func receiver(stopper <-chan time.Time, ch <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-stopper:
			fmt.Println("Exiting receiver...")
			wg.Done()
			return
		case data := <-ch:
			fmt.Println(data)
		}

	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	stopper := time.After(3 * time.Second)

	wg.Add(2)

	go producer(stopper, ch, &wg)
	go receiver(stopper, ch, &wg)

	wg.Wait()

}
