package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	timer := time.NewTimer(d).C
	<-timer
}

func main() {
	fmt.Println("Sleeping for 2 seconds...")
	sleep(2 * time.Second)
	fmt.Println("Awake!")
}
