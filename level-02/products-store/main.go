/*
A store sells two types of products: books and games. Each product has the fields name (string) and price (float). Define the following functionalities:
Product has the following methods:
A method to print the information of each product (type, name, and price)
A method to apply a discount ratio to the price
The store should be able to apply custom discounts based on the type of product: 10% discount for books and 20% discount for games.
*/

package main

import "fmt"

type Book struct {
	Name  string
	Price float32
}

func (b *Book) PrintInfo() string {
	return fmt.Sprintf("* Book Name: %s - Price: $%.2f", b.Name, b.Price)
}

func (b *Book) ApplyDiscount(discountRatio float32) {
	b.Price = (b.Price * (1 - discountRatio/100))
}

type Game struct {
	Name  string
	Price float32
}

func (g *Game) PrintInfo() string {
	return fmt.Sprintf("* Game Name: %s - Price: $%.2f", g.Name, g.Price)
}

func (g *Game) ApplyDiscount(discountRatio float32) {
	g.Price = (g.Price * (1 - discountRatio/100))
}

type Product interface {
	PrintInfo() string
	ApplyDiscount(discountRatio float32)
}

func CalculateDiscount(p Product) {
	switch p.(type) {
	case *Book:
		p.ApplyDiscount(10)
	case *Game:
		p.ApplyDiscount(20)
	}
}

func main() {
	book := &Book{Name: "The Hobbit", Price: 30}
	CalculateDiscount(book)
	fmt.Printf("%s\n", book.PrintInfo())
	game := &Game{Name: "FIFA 2021", Price: 100}
	CalculateDiscount(game)
	fmt.Printf("%s\n", game.PrintInfo())
}
