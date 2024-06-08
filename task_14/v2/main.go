package main

import (
	"fmt"
)

func main() {
	var a interface{} = 10
	var b interface{} = "text"
	var c interface{} = true
	var d interface{} = make(chan int)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
}
