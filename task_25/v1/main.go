package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Sleeping for 2 seconds...")
	sleep(2 * time.Second)
	fmt.Println("Awake!")
}
