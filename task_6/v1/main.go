package main

import (
	"context"
	"fmt"
	"time"
)

func goroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Leaving program")
			return
		default:
			fmt.Println("I`ma still working")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go goroutine(ctx)

	time.Sleep(2 * time.Second)

	cancel()

	time.Sleep(2 * time.Second)

}
