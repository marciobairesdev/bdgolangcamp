/*
 * 		- Statement:
 * 			Create a channel of int and then create a goroutine to add a value to
 * 			the channel and then print the channel value in the main function.
 */

package main

import "fmt"

func addValueToChannel(value int, ch chan int) {
	ch <- value
}

func main() {
	ch := make(chan int)
	go addValueToChannel(555, ch)
	fmt.Println("Channel", <-ch)
}
