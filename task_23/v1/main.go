package main

import "fmt"

func removeByIndex(index int, slice []string) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []string{"one", "two", "three", "four"}

	fmt.Println(removeByIndex(2, slice))

}
