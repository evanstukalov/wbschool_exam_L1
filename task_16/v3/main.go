package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	arr := IntSlice{3, 6, 8, 10, 1, 2, 1}

	sort.Sort(arr)

	fmt.Println(arr)
}
