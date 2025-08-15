package main

import (
	"fmt"
	"math"
)

func maxAndSecMax(arr [5]int) (int, int) {
	max := math.MinInt
	min := math.MinInt

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			min = max
			max = arr[i]
		} else if arr[i] > min && arr[i] < max {
			min = arr[i]
		}
	}

	return max, min
}

func runTests() {
	type test struct {
		input   [5]int
		wantMax int
		wantSec int
	}
	tests := []test{
		{[5]int{5, 4, 3, 2, 1}, 5, 4},
		{[5]int{1, 1, 1, 1, 1}, 1, math.MinInt},
		{[5]int{-1, -2, -3, -4, -5}, -1, -2},
	}

	passCount := 0
	for _, tc := range tests {
		gotMax, gotSec := maxAndSecMax(tc.input)
		if gotMax == tc.wantMax && gotSec == tc.wantSec {
			passCount++
		} else {
			fmt.Printf("❌ FAIL: input=%v got=(%d,%d) want=(%d,%d)\n",
				tc.input, gotMax, gotSec, tc.wantMax, tc.wantSec)
		}
	}
	fmt.Printf("✅ %d/%d test cases passed\n", passCount, len(tests))
}

func main() {
	// Run tests first
	runTests()

	// Your actual program logic
	fmt.Println("Running main program...")
	arr := [5]int{10, 5, 8, 12, 7}
	arr2 := arr
	arr2[2] = 100
	fmt.Println("this is array 2", arr2)
	max, sec := maxAndSecMax(arr)
	fmt.Printf("Max: %d, Second Max: %d\n", max, sec)
}
