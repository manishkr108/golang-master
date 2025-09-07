package main

import (
	"fmt"
	"math"
)

func consecutive(arr []int) float64 {
	var maxCount float64 = 0.0
	var count float64 = 0.0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			count++
		} else {
			maxCount = math.Max(count, maxCount)
			count = 0

		}
	}

	return maxCount
}

func main() {
	arr := []int{1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1,0, 1}
	res := consecutive(arr)
	fmt.Println(res)
}
