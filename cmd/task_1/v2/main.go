package main

import "fmt"

// The Speaker interface defines the Speak() method
type Speaker interface {
	Speak()
}

type Human struct {
	Name string
}

type Action struct {
	Speaker
	typeOf string
}

func (h *Human) Speak() {
	fmt.Printf("My name is %s\n", h.Name)
}

func (a *Action) DoActivity() {
	// Accessing Speak() through the embedded Speaker interface
	a.Speaker.Speak()

	fmt.Printf("%s is doing action of type %s\n", a.Speaker.(*Human).Name, a.typeOf)
}

func main() {
	human := &Human{Name: "Alex"}

	// Creating an instance of the Action struct, initializing the Speaker and typeOf fields
	action := Action{
		Speaker: human,
		typeOf:  "eating"}

	action.DoActivity()
}
