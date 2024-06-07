/*
Отдельный канал для сигналов может быть хорошим выбором для более простых сценариев,
где требуется только передача сигнала завершения, и нет необходимости в дополнительных возможностях, предоставляемых контекстом.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func worker(id int, dataCh <-chan int, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data := <-dataCh:
			fmt.Printf("Worker #%d received data %d \n", id, data)

		case <-doneCh:
			return
		}
	}

}

func main() {
	var workerNum int
	flag.IntVar(&workerNum, "workers", 10, "number of workers")
	flag.Parse()

	dataCh := make(chan int)
	doneCh := make(chan struct{})

	var wg sync.WaitGroup

	for i := 0; i < workerNum; i++ {
		wg.Add(1)
		go worker(i, dataCh, doneCh, &wg)
	}

	go func() {
		for i := 0; ; i++ {
			dataCh <- i
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	close(doneCh)

	wg.Wait()

	time.Sleep(2 * time.Second)

	fmt.Println("Active goroutines:", runtime.NumGoroutine())
}
