package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func producer(ctx context.Context, ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting producer...")
			wg.Done()
			return
		case ch <- i:
			time.Sleep(3 * time.Second)
		}
	}
}

func receiver(ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	wg.Add(2)

	go producer(ctx, ch, &wg)
	go receiver(ctx, ch, &wg)

	wg.Wait()

}

// через time.After
