package main

import "fmt"

type Set struct {
	elements map[int]struct{}
}

func (s *Set) Add(elements []int) {
	for _, element := range elements {
		s.elements[element] = struct{}{}
	}
}

func Intersection(s1 Set, s2 Set) Set {
	set := Set{elements: make(map[int]struct{})}

	for k1 := range s1.elements {
		for k2 := range s2.elements {
			if k1 == k2 {
				set.Add([]int{k1})
			}
		}
	}
	return set
}

func main() {
	set1 := Set{elements: make(map[int]struct{})}
	set1.Add([]int{1, 2, 3, 4})

	set2 := Set{elements: make(map[int]struct{})}
	set2.Add([]int{2, 3})

	set3 := Intersection(set1, set2)

	for k := range set3.elements {
		fmt.Print(k, " ")
	}

}
