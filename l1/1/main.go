package main

import "fmt"

/*
	Дана структура Human (с произвольным набором полей и методов).

	Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

	Подсказка: используйте композицию (embedded struct), чтобы Action имел все методы Human.
*/

type Human struct {
	name string
	age  int
}

type Action struct {
	Human
	anotherField string
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) SetName(s string) {
	h.name = s
}

func (h *Human) GetAge() int {
	return h.age
}

func (h *Human) SetAge(n int) {
	h.age = n
}

func (a *Action) GetField() string {
	return a.anotherField
}

func (a *Action) SetField(s string) {
	a.anotherField = s
}

func main() {
	a := Action{}

	a.SetName("Petya")
	a.SetAge(10)
	a.SetField("another")

	fmt.Println(a.GetName())
	fmt.Println(a.GetAge())
	fmt.Println(a.GetField())
}
