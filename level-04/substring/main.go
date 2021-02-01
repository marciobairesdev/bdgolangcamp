/*
 * Given a list of strings find all strings that contain a given substring.
 * You should split the list into chunks and process them in parallel.
 * Each goroutine should write the response to an output channel which
 * will be consumed in the end to consolidate the response.
 */

package main

import (
	"fmt"
	"strings"
	"sync"
)

func findString(ch chan string, w *sync.WaitGroup, searchString string, stringList []string) {
	defer w.Done()
	for _, s := range stringList {
		if strings.Contains(s, searchString) {
			ch <- s
		}
	}
}

func print(ch chan string, w *sync.WaitGroup) {
	defer w.Done()
	for c := range ch {
		fmt.Println(c)
	}
}

func main() {
	stringList := []string{"Marcio", "Cesar", "Juan", "Carlos", "Maria", "Alejandra"}
	ch := make(chan string, 6)
	var wg sync.WaitGroup

	wg.Add(2)

	go findString(ch, &wg, "ar", stringList[:3])
	go findString(ch, &wg, "ar", stringList[3:])

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go print(ch, &wg2)
	wg.Wait()
	close(ch)
	wg2.Wait()
}
