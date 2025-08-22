package main

import "fmt"

/* Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human. */

type Human struct {
	name       string
	occupation string
}

type Action struct {
	Human
}

func (h Human) GetOccupation() {
	fmt.Printf("%s is a %s!\n", h.name, h.occupation)
}

func (h *Human) ChangeOccupation(newJob string) {
	h.occupation = newJob
	fmt.Printf("%s is now a %s!\n", h.name, h.occupation)
}

func (h Human) Greet() {
	fmt.Printf("%s says \"hi!\"\n", h.name)
}

func main() {
	human := Human{
		name:       "Anton",
		occupation: "Data Engineer",
	}

	action := Action{Human: human}

	human.Greet()
	action.Greet()

	human.GetOccupation()
	action.GetOccupation()

	action.ChangeOccupation("Barista")
}
