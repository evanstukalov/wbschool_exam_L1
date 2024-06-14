package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {
	runes := []rune(str)

	var builder strings.Builder
	builder.Grow(len(runes))

	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}

	return builder.String()
}

func main() {
	result := Reverse("главрыба")
	fmt.Println(result)
}
