package main

import "fmt"

type (
	Human struct {
		Name string
		Age  int
	}

	Action struct {
		Human
	}
)

func (h *Human) GetName() string {
	return h.Name
}

func (h *Human) GetAge() int {
	return h.Age
}

func main() {
	action := &Action{
		Human: Human{
			Name: "Andrey",
			Age:  20,
		},
	}

	fmt.Println("Name:", action.GetName())
	fmt.Println("Age:", action.GetAge())
}
