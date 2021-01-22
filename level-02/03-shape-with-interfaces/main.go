/*
 * Shape With Interfaces
 * 		- Statement:
 * 			Create a Rectangle struct and a Circle struct, both struct
 * 			should have two methods: Area() and Perimeter().
 * 			And then create a Shape interface for those struct.
 * 			After that create a function that will receive a Shape interface
 * 			as parameter and will execute the Area() and the Perimeter()
 * 			from each struct.
 */

package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float32
	Height float32
}

func (r *Rectangle) Area() float32 {
	return (r.Width * r.Height)
}

func (r *Rectangle) Perimeter() float32 {
	return ((2 * r.Width) + (2 * r.Height))
}

type Circle struct {
	Radius float32
}

func (c *Circle) Area() float32 {
	return (math.Pi * c.Radius * c.Radius)
}

func (c *Circle) Perimeter() float32 {
	return (2 * math.Pi * c.Radius)
}

type Shape interface {
	Area() float32
	Perimeter() float32
}

func Calculate(s Shape) {
	fmt.Printf("* Shape Type: %T\n", s)
	fmt.Printf("* Area: %.2f\n", s.Area())
	fmt.Printf("* Perimeter: %.2f\n\n", s.Perimeter())
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 3},
		&Rectangle{Width: 3, Height: 4},
	}

	for _, shape := range shapes {
		Calculate(shape)
	}
}
