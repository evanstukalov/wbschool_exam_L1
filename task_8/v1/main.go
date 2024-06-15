package main

import (
	"fmt"
)

func setBit(n int64, i uint, value int) int64 {
	if value == 1 {
		n |= (1 << i)
	} else {
		n &= ^(1 << i)
	}
	return n
}

func main() {
	var num int64 = 42
	var index uint = 3
	var value int = 1

	fmt.Printf("Before: %064b\n", num)
	num = setBit(num, index, value)
	fmt.Printf("After:  %064b\n", num)
}
