package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func goroutine(quitCh <-chan struct{}) {
	for {
		select {
		case <-quitCh:
			fmt.Println("Leaving program")
			return
		default:
			fmt.Println("I`ma still working")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	quitCh := make(chan struct{})

	go goroutine(quitCh)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		close(quitCh)

	}()

	time.Sleep(5 * time.Second)

}
