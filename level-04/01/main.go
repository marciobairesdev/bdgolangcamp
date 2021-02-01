/*
 *		- Statement:
 * 			Create a goroutine that will execute an anonymous function to print
 * 			“Hello World” and in the main routine print “main function”.
 */

package main

import (
	"fmt"
	"time"
)

func helloWorld() {
	fmt.Println("Hello World")
}

func main() {
	go helloWorld()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")
}
