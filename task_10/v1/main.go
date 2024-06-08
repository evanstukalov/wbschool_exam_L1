package main

import (
	"fmt"
)

func main() {
	tempCollection := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	tempMap := map[int][]float64{}

	for _, temp := range tempCollection {
		key := int(temp/10) * 10
		tempMap[key] = append(tempMap[key], temp)
	}

	for key, temps := range tempMap {
		fmt.Printf("%d: %v\n", key, temps)
	}

}
