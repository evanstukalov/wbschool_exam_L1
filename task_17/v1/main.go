package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
