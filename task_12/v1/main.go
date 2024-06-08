package main

import "fmt"

type Set struct {
	elements map[string]struct{}
}

func (s *Set) Add(elements []string) {
	for _, element := range elements {
		s.elements[element] = struct{}{}
	}
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := Set{elements: make(map[string]struct{})}

	set.Add(strings)

	for k := range set.elements {
		fmt.Println(k)
	}
}
