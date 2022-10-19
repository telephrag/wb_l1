package main

import "fmt"

type Human struct{}

func (h *Human) Breath() { fmt.Println("breathes...") }

type Action struct {
	Human
}

func (a *Action) Breath() { a.Human.Breath() }

func main() {
	a := &Action{Human: Human{}}
	a.Breath()
	a.Human.Breath()
}
