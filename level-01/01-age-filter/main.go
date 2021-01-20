/*
 * Age Filter
 * 		- Statement:
 * 			Create a function that will filter a Slice of ages that are between the range.
 * 			The function will receive two numbers and a slice of ages as parameters.
 * 			It should return the ages between the range.
 */

package main

import (
	"errors"
	"fmt"
)

func filterAges(startAge int, endAge int, ages []int) (filteredAges []int, err error) {
	if startAge > endAge {
		return nil, errors.New("* Error: start age must be less than or equal to end age")
	}

	for _, i := range ages {
		if i >= startAge && i <= endAge {
			filteredAges = append(filteredAges, i)
		}
	}

	return filteredAges, nil
}

func main() {
	ages := []int{0, 1, 2, 4, 5, 10, 15, 27, 39, 48, 50}
	startAge := 3
	endAge := 48
	filteredAges, err := filterAges(startAge, endAge, ages)

	fmt.Printf("* Initial Ages: %v\n", ages)
	fmt.Printf("* Start Age Filter: %d\n", startAge)
	fmt.Printf("* End Age Filter: %d\n", endAge)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("* Filtered Ages: %v\n", filteredAges)
	}
}
