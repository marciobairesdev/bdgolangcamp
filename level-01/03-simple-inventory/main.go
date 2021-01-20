/*
 * Simple Inventory
 * 		- Statement:
 * 			Create a Product struct with an attribute ID and Name, and an Inventory struct that
 * 			will contain a map[string]Products. Create a method with a pointer of Inventory as a
 * 			receiver and It will expect a product to be added into the map of products.
 *
 * 			Required: Add validations into the Add method, to check that the ID can’t be empty
 * 			and also not duplicated ID.
 */

package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID   int
	Name string
}

type Inventory struct {
	Products map[int]Product
}

func (i *Inventory) add(p Product) error {
	if p.ID == 0 {
		return errors.New("* Please provid a product ID")
	} else if _, ok := i.Products[p.ID]; ok {
		return errors.New("* The provided ID already exists")
	}

	i.Products[p.ID] = p
	return nil
}

func main() {
	i := Inventory{
		Products: make(map[int]Product),
	}

	fmt.Printf("* Initial Inventory: %v\n", i)

	p := Product{
		ID:   1,
		Name: "iPhone SE",
	}

	fmt.Printf("* Trying to add %v...\n", p)
	err := i.add(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("* Current Inventory: %v\n", i)

	p = Product{
		ID:   2,
		Name: "Moto G",
	}

	fmt.Printf("* Trying to add %v...\n", p)
	err = i.add(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("* Current Inventory: %v\n", i)

	p = Product{
		ID:   2,
		Name: "Redmi 7",
	}

	fmt.Printf("* Trying to add %v...\n", p)
	err = i.add(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("* Current Inventory: %v\n", i)
}
