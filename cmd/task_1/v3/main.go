package main

import "fmt"

type Human struct {
	Name string
}

type Action struct {
	Human
	typeOf string
}

func (h Human) Speak() {
	fmt.Printf("My name is %s\n", h.Name)
}

// Speak method for the Action struct, calls the Speak method of the Human struct
func (a Action) Speak() {
	a.Human.Speak()
}

func (a Action) DoActivity() {
	fmt.Printf("%s is doing action of type %s\n", a.Name, a.typeOf)
}

func main() {
	action := Action{
		Human:  Human{Name: "Alex"},
		typeOf: "eating"}

	action.Speak()
	action.DoActivity()
}
