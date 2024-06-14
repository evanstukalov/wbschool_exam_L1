package main

import "fmt"

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
	a.Speaker.Speak()

	fmt.Printf("%s is doing action of type %s\n", a.Speaker.(*Human).Name, a.typeOf)
}

func main() {
	human := &Human{Name: "Alex"}

	action := Action{
		Speaker: human,
		typeOf:  "eating"}

	action.DoActivity()
}
