package main

import (
	"math"
	"testing"
)

/*
✅ Key takeaway:

	t *testing.T is how the Go testing framework communicates

	with your test code — it’s your only handle for marking pass/fail,
	 logging, and controlling the test run.

	 TestXxx → Must start with Test (capital T).

(t *testing.T) → Must take exactly one argument, a pointer to testing.T from the testing package.
*/
func TestMaxAndSecMax(t *testing.T) {
	tests := []struct {
		name    string
		input   [5]int
		wantMax int
		wantSec int
	}{{
		name:    "Normal Case",
		input:   [5]int{10, 3, 45, 7, 20},
		wantMax: 45,
		wantSec: 20,
	}, {
		name:    "Duplicate Max",
		input:   [5]int{10, 3, 45, 45, 7},
		wantMax: 45,
		wantSec: 10,
	},
		{
			name:    "All equal",
			input:   [5]int{5, 5, 5, 5, 5},
			wantMax: 5,
			wantSec: math.MinInt, // no second max found
		},
		{
			name:    "Negative numbers",
			input:   [5]int{-1, -2, -3, -4, -5},
			wantMax: -1,
			wantSec: -2,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			gotMax, gotSec := maxAndSecMax(tc.input)
			if gotMax != tc.wantMax || gotSec != tc.wantSec {
				t.Errorf("maxAndSecMax(%v) = (%d, %d); want (%d, %d)",
					tc.input, gotMax, gotSec, tc.wantMax, tc.wantSec)
			}
		})

	}
}
