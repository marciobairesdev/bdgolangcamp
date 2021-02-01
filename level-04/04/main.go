/*
 * 		- Statement:
 * 			Create a function that will increase a number and that function will be executed
 * 			by a goroutine inside a for loop (x1000 times). To avoid race conditioning,
 * 			implement the sync.Mutex and Lock and Unlock inside the increase() function.
 * 			Note: Add a time.Sleep() to be able to see the final n.
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	Mutex sync.Mutex
	Value int
}

func (sn *SafeNumber) Increase() {
	sn.Mutex.Lock()
	sn.Value++
	sn.Mutex.Unlock()
}

func main() {
	sn := SafeNumber{Value: 0}
	for i := 0; i < 1000; i++ {
		go sn.Increase()
	}

	time.Sleep(time.Second)
	fmt.Println(sn.Value)
}
