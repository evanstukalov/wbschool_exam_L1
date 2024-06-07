/*
context.Context рекомендуется использовать для более сложных и масштабируемых систем,
где требуется управление временем выполнения, дедлайнами, отменой операций и передача значений между горутинами.
*/

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case data := <-dataCh:
			fmt.Printf("Worker #%d received data %d \n", id, data)
		case <-ctx.Done():
			fmt.Printf("Worker #%d stopped \n", id)
			return
		}
	}
}

func producer(ctx context.Context, dataCh chan<- int) {
	for i := 0; ; i++ {
		select {
		case dataCh <- i:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			fmt.Println("Producer stopped")
			return
		}
	}
}

func main() {
	var workerNum int
	flag.IntVar(&workerNum, "workers", 10, "number of workers")
	flag.Parse()

	dataCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go worker(ctx, i, dataCh, &wg)
	}

	go producer(ctx, dataCh)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	cancel()

	wg.Wait()

	time.Sleep(2 * time.Second)

	fmt.Println("Active goroutines:", runtime.NumGoroutine())
}
