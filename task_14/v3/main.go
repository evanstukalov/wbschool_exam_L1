package main

import (
	"fmt"
)

func typeChecker(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println(t, " is int")
	case string:
		fmt.Println(t, " is string")
	case bool:
		fmt.Println(t, " is bool")
	case chan int:
		fmt.Println(t, " is chan")

	}
}

func main() {
	var a interface{} = 10
	var b interface{} = "text"
	var c interface{} = true
	var d interface{} = make(chan int)

	typeChecker(a)
	typeChecker(b)
	typeChecker(c)
	typeChecker(d)

}
