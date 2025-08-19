package main

import (
	"reflect"
	"testing"
)

func TestChunkSlice(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		size   int
		result [][]int
	}{
		{
			name:   "Normal Test Case",
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			size:   2,
			result: [][]int{{1, 2}, {3, 4}, {5, 6}, {7}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := chunkSlice(tc.input, tc.size)
			if !reflect.DeepEqual(res, tc.result) {
				t.Errorf("expected %v, got %v", tc.result, res)
			}
		})
	}

}
