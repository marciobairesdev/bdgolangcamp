/*
 * Inteface Switch
 * 		- Statement:
 * 			Create a function that will receive an interface and it will print
 * 			the interface value with a specific message based on the type.
 */

package main

import (
	"fmt"
	"strings"
)

type Puppy struct {
	Name string
	Roar string
}

func (p *Puppy) Cry() string {
	return fmt.Sprintf("* The %s is a PUPPY and here's its roar: %s!!!\n", strings.ToUpper(p.Name), strings.ToUpper(p.Roar))
}

type WildBeast struct {
	Name    string
	Country string
	Roar    string
}

func (wb *WildBeast) Cry() string {
	return fmt.Sprintf("* The %s is a WILD BEAST from %s and here's its roar: %s!!!\n", strings.ToUpper(wb.Name), strings.ToUpper(wb.Country), strings.ToUpper(wb.Roar))
}

type Animal interface {
	Cry() string
}

func GetSound(a Animal) string {
	switch a.(type) {
	case *Puppy:
		return a.Cry()
	case *WildBeast:
		return a.Cry()
	}
	return ""
}

func main() {
	dog := &Puppy{Name: "Dog", Roar: "Wooof"}
	lion := &WildBeast{Name: "Lion", Country: "South Africa", Roar: "Grrrr"}

	fmt.Print(GetSound(dog))
	fmt.Print(GetSound(lion))
}
