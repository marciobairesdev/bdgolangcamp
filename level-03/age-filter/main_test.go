/*
 * - Statement:
 * 		Create many test cases as needed to validate the right behavior of the Age Filter exercise from Level 1
 */

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterByAgeRange(t *testing.T) {
	ages := []int{0, 1, 2, 3, 7, 9, 11, 25, 27, 38, 44, 59, 77, 105}
	fromAge := 11
	toAge := 44
	filteredAges := []int{11, 25, 27, 38, 44}

	assert.Equal(t, filteredAges, filterByAgeRange(fromAge, toAge, ages))

	ageTests := []struct {
		ages    []int
		fromAge int
		toAge   int
		wanted  []int
	}{
		{ages: []int{0, 1, 2, 3, 4, 5, 6}, fromAge: 1, toAge: 4, wanted: []int{1, 2, 3, 4}},
		{ages: []int{10, 15, 20, 21, 22, 50, 60}, fromAge: 15, toAge: 21, wanted: []int{15, 20, 21}},
		{ages: []int{4, 5, 6, 17, 35, 80, 105, 110}, fromAge: 6, toAge: 85, wanted: []int{6, 17, 35, 80}},
	}

	for _, test := range ageTests {
		assert.Equal(t, test.wanted, filterByAgeRange(test.fromAge, test.toAge, test.ages))
	}
}
