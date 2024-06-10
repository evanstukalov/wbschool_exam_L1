package main

import (
	"fmt"
)

func createHugeString(size int) string {
	hugeString := make([]byte, size)
	for i := range hugeString {
		hugeString[i] = 'a'
	}

	return string(hugeString)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return string(v[:100])
}

func main() {
	result := someFunc()
	fmt.Println(result)
}
