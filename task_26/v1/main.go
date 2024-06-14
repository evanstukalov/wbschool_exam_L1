package main

import (
	"fmt"
	"strings"
)

func uniqueCheck(str string) bool {
	uniqueChars := make(map[rune]bool)

	s := strings.ToLower(str)

	for _, c := range s {
		if uniqueChars[c] {
			return false
		}

		uniqueChars[c] = true

	}
	return true
}

func main() {
	fmt.Println(uniqueCheck("abcd"))      // true
	fmt.Println(uniqueCheck("Aabcd"))     // false
	fmt.Println(uniqueCheck("abCdefAaf")) // false
	fmt.Println(uniqueCheck("aabcd"))     // false
}
