/*
 * Simple Queue
 * 		- Statement:
 * 			Create a Queue implementation. To build it we will need a slice of int,
 * 			and an enqueue function to add an int into the slice and dequeue function
 * 			to remove the first element of the slice.
 */

package main

import "fmt"

type Queue struct {
	Items []int
}

func (q *Queue) enqueue(item int) {
	q.Items = append(q.Items, item)
}

func (q *Queue) dequeue() {
	q.Items = q.Items[1:]
}

func main() {
	q := Queue{Items: make([]int, 0)}
	fmt.Printf("* Initial Queue: %v\n", q)

	fmt.Printf("* Adding 1 to queue...\n")
	q.enqueue(1)
	fmt.Printf("* Current Queue: %v\n", q)
	fmt.Printf("* Adding 2 to queue...\n")
	q.enqueue(2)
	fmt.Printf("* Current Queue: %v\n", q)
	fmt.Printf("* Adding 3 to queue...\n")
	q.enqueue(3)
	fmt.Printf("* Current Queue: %v\n", q)
	fmt.Printf("* Removing the 1st element of the queue...\n")
	q.dequeue()
	fmt.Printf("* Current Queue: %v\n", q)
	fmt.Printf("* Removing the 1st element of the queue...\n")
	q.dequeue()
	fmt.Printf("* Current Queue: %v\n", q)
	fmt.Printf("* Removing the 1st element of the queue...\n")
	q.dequeue()
	fmt.Printf("* Current Queue: %v\n", q)
}
