package main

import "testing"

/*
 * Single Test case
 * 		- Statement:
 * 			Create a Testing function to check the behavior of the following function.
 * 			If the function returns a different value from the expected one,
 * 			return an error specifying the test case.
 */
func TestSingleIntMin(t *testing.T) {
	if min := IntMin(0, 55); min != 0 {
		t.Errorf("IntMin(2, 1) = 1 want 1")
	}
}

/*
 * Multiple Test cases
 * 		- Statement:
 * 			Create a Testing function to check the behavior of the following function.
 * 			Create a struct with the function parameter and the ‘want’ (expected) value,
 * 			then iterate all the test cases and validate the behavior of the function
 * 			against each case.
 */
func TestMultipleIntMin(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		result int
	}{
		{a: 10, b: 25, result: 10},
		{a: 32, b: 111, result: 32},
		{a: 883, b: 1070, result: 883},
		{a: 4, b: 62, result: 4},
	}

	for _, tt := range tests {
		if min := IntMin(tt.a, tt.b); min != tt.result {
			t.Errorf("IntMin(%v, %v) = %v, want %v", tt.a, tt.b, min, tt.result)
		}
	}
}
