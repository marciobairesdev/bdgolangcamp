/*
 * Simple Calculator Operations
 * 		- Statement:
 * 			Create a calculator struct with a result attribute and with the most common methods:
 * 				- Add;
 * 				- Subtract;
 * 				- Multiply;
 * 				- Divide.
 *
 *  		Plus: Handle possible error, Divide function won’t be able to handle second parameter
 * 			with value zero, instead function should panic and print an alert.
 */

package main

import (
	"fmt"
)

type Calculator struct {
	Result float64
}

func (c *Calculator) add(a float64, b float64) {
	c.Result = a + b
}

func (c *Calculator) subtract(a float64, b float64) {
	c.Result = return a - b
}

func (c *Calculator) multiply(a float64, b float64) {
	c.Result = a * b
}

func (c *Calculator) divide(a float64, b float64) {
	if b == 0 {
		panic("* Error: division by zero not allowed")
	}

	c.Result = a / b
}

func main() {
	c := Calculator{}

	c.add(5, 7)
	fmt.Printf("* 5 + 7 = %f\n", c.Result)

	c.add(3.2, 1.5)
	fmt.Printf("* 3.2 + 1.5 = %f\n", c.Result)

	c.subtract(55, 12)
	fmt.Printf("* 55 - 12 = %f\n", c.Result)

	c.subtract(8.8, 13.9)
	fmt.Printf("* 8.8 - 13.9 = %f\n", c.Result)

	c.multiply(9, 10)
	fmt.Printf("* 9 * 10 = %f\n", c.Result)

	c.multiply(100.45, 3.6)
	fmt.Printf("* 100.45 * 3.6 = %f\n", c.Result)

	c.divide(4, 3)
	fmt.Printf("* 4 / 2 = %f\n", c.Result)

	c.divide(2.2, 1.3)
	fmt.Printf("* 2.2 / 1.3 = %f\n", c.Result)

	c.divide(0, 8)
	fmt.Printf("* 0 / 8 = %f\n", c.Result)

	fmt.Printf("* 5 / 0 = ERR\n")
	c.divide(5, 0)
}
