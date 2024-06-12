package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}

	for i := range arr {
		if i == len(arr)/2 {
			continue
		}
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)

}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println(quicksort(arr))
}
