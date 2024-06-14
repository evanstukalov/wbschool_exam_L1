package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("snow dog sun"))
}

func ReverseWords(str string) string {
	var builder strings.Builder
	builder.Grow(len(str))
	stringSlice := strings.Split(str, " ")

	for i := range stringSlice {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(stringSlice[len(stringSlice)-1-i])
	}

	return builder.String()
}
