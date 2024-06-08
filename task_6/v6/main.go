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
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	go goroutine(ctx)

	time.Sleep(5 * time.Second)

}
