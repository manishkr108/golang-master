package main

import (
	"reflect"
	"testing"
)

func TestMergeSlice(t *testing.T) {
	testmergeslice := []struct {
		name     string
		inputone []int
		m        int
		inputtwo []int
		n        int
		result   []int
	}{
		{
			name:     "Merge Two sorted slice",
			inputone: []int{1, 2, 3, 0, 0, 0},
			m:        3,
			inputtwo: []int{2, 4, 5},
			n:        3,
			result:   []int{1, 2, 2, 3, 4, 5},
		},
	}

	for _, tc := range testmergeslice {
		t.Run(tc.name, func(t *testing.T) {
			mergeTwoSlice(tc.inputone, tc.m, tc.inputtwo, tc.n)
			if !reflect.DeepEqual(tc.inputone, tc.result) {
				t.Errorf("expected %v, got %v", tc.result, tc.inputone)
			}

		})
	}

}
