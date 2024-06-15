package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	end := time.Now().Add(d)

	for time.Now().Before(end) {
	}
}

func main() {
	fmt.Println("Sleeping for 2 seconds...")
	sleep(2 * time.Second)
	fmt.Println("Awake!")
}
