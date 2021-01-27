package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Animal struct {
	Name string
}

/*
 * Different Asserts
 * 		- Statement:
 * 			Using Testify/assert pkg, execute different functions and validate the expected result.
 * 			You can create as many cases as you want: Equal | NotEqual | Nil | NotNil
 */
func TestDifferentAsserts(t *testing.T) {
	assert.Equal(t, 9, 9)
	assert.NotEqual(t, 5, 0)
	assert.Nil(t, nil)
	assert.True(t, true)
	assert.False(t, false)

	dog := Animal{Name: "Dog"}
	assert.NotNil(t, dog)
	assert.Equal(t, "Dog", dog.Name)
	assert.NotEqual(t, "DOG", dog.Name)
	assert.NotEqual(t, "Cat", dog.Name)
}
