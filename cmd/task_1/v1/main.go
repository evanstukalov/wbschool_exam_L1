package main

import "fmt"

type Human struct {
	Name string
}

// Action struct embeds the Human struct and adds an additional field typeOf
type Action struct {
	Human
	typeOf string
}

// Speak is a method for the Human struct that prints the Human's name
func (h Human) Speak() {
	fmt.Printf("My name is %s\n", h.Name)
}

// DoActivity is a method for the Action struct that prints the type of action
func (a Action) DoActivity() {
	fmt.Printf("%s is doing action of type %s\n", a.Name, a.typeOf)
}

func main() {
	// Creating an instance of Action struct, initializing the embedded Human struct and the typeOf field
	action := Action{
		Human:  Human{Name: "Alex"},
		typeOf: "eating"}

	// Calling the Speak method of the embedded Human struct
	action.Speak()

	action.DoActivity()
}
