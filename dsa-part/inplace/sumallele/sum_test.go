package main

import (
	"testing"
)

// sumArray is the function we want to test.
// It sums all integers in a slice and returns the result
func TestSumAllVal(t *testing.T) {
	test := []struct {
		name     string // test case name for readability
		input    []int  // data to pass into sumArray
		valiwant int    //expected result
	}{
		{
			name:     "Simple Test",
			input:    []int{1, 2, 3, 4, 5},
			valiwant: 15,
		},
		{
			name:     "Nigative Val",
			input:    []int{-1, -2, -3, -4, -5},
			valiwant: -15,
		},
		{
			name:     "Empty Slice",
			input:    []int{},
			valiwant: 0,
		},
	}
	// 2️⃣ Loop through each test case
	for _, tc := range test {
		// 3️⃣ Create a subtest for each case
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			sum := sumArray(tc.input)
			// Compare result with expected value
			if sum != tc.valiwant {
				// ❌ Fail if result is wrong
				t.Errorf("sumArray(%v) = %d; want %d", tc.input, sum, tc.valiwant)
				return
			} else {
				// ✅ Pass message (only shown with `go test -v`)
				t.Logf("✅ %s passed: input=%v sum=%d", tc.name, tc.input, sum)
			}
		})
	}
}
