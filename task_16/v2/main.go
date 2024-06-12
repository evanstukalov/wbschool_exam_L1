package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Println(arr)
}
