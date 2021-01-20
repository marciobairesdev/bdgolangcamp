/*
Statement
Create a Stack structure that means a LIFO (Last In First Out).
Note: The func main should be provided and Just need to create the Stack struct and Methods to Push, Pop and Peek. IsEmpty is optional.
To test it, push two or three strings and then pop the last one and finally, push another one.
*/

package main

import (
	"fmt"
)

type Stack struct {
	Items []string
}

func (s *Stack) Push(value string) {
	s.Items = append(s.Items, value)
}

func (s *Stack) Peek() string {
	return s.Items[len(s.Items)-1]
}

func (s *Stack) Pop() string {
	valuePoped := s.Items[len(s.Items)-1]
	s.Items = s.Items[0 : len(s.Items)-1]
	return valuePoped
}

func NewStack() Stack {
	return Stack{
		Items: make([]string, 0),
	}
}

func main() {
	s := NewStack()
	s.Push("Apple")
	fmt.Printf("Current Stack: %+v \n", s) // Current Stack: {nodes:[Apple] count:1}
	s.Push("Melon")
	fmt.Printf("Current Stack: %+v \n", s) // Current Stack: {nodes:[Apple Melon] count:2}
	s.Push("Mangoes")
	fmt.Printf("Current Stack: %+v \n", s)             // Current Stack: {nodes:[Apple Melon Mangoes] count:3}
	fmt.Printf("Show top element: (%s)\n", s.Peek())   // Show top element: (Mangoes)
	fmt.Printf("Current Stack: %+v \n", s)             // Current Stack: {nodes:[Apple Melon Mangoes] count:3}
	fmt.Printf("Extract top element: (%s)\n", s.Pop()) // Extract top element: (Mangoes)
	fmt.Printf("Current Stack: %+v \n", s)             // Current Stack: {nodes:[Apple Melon] count:2}
	s.Push("Pineapple")
	fmt.Printf("Current Stack: %+v \n", s) // Current Stack: {nodes:[Apple Melon Pineapple] count:3}
}
