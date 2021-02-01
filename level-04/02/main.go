/*
 * 		- Statement:
 * 			Create a function goroutine that will execute an anonymous function to just print
 * 			the number “1”, in the main function print the number “0” and also add a
 * 			time.Sleep() to wait 2 seconds.
 */

package main

import (
	"fmt"
	"time"
)

func printNumber() {
	fmt.Println("1")
}

func main() {
	go printNumber()
	fmt.Println("0")
	time.Sleep(2 * time.Second)
}
