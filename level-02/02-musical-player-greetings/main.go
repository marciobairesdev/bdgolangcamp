/*
 * Musical Player Greetings
 * 		- Statement:
 * 			Create a Trumpeter struct and a Violinist struct both struts should have
 * 			a Name attribute and a Greetings() function to present themself.
 * 			Then create a MusicalPlayer interface. In the main function, create a slice
 * 			with two or more musical players and for each call the Greetings() function.
 */

package main

import (
	"fmt"
	"strings"
)

type Trumpeter struct {
	Name string
}

func (t *Trumpeter) Greetings() string {
	return fmt.Sprintf("Greetings from Trumpeter %s!", strings.ToUpper(t.Name))
}

type Violinist struct {
	Name string
}

func (v *Violinist) Greetings() string {
	return fmt.Sprintf("Greetings from Violinist %s!", strings.ToUpper(v.Name))
}

type MusicalPlayer interface {
	Greetings() string
}

func main() {
	musicalPlayers := []MusicalPlayer{
		&Trumpeter{Name: "Louis Armstrong"},
		&Violinist{Name: "Niccolò Paganini"},
	}

	for _, mp := range musicalPlayers {
		fmt.Printf("* %s\n", mp.Greetings())
	}
}
