package main

import (
	"fmt"

	"sort"
)

// Worst case
func missingNum(arr []int) (int, bool) {
	sort.Ints(arr)
	axpected := 1
	for _, v := range arr {
		if v != axpected {
			return axpected, true
		}
		axpected++
	}

	return -1, false
}

//best solution

func findMissing(arr []int) int {
	n := len(arr)
	total := n * (n + 1) / 2
	fmt.Println(total)
	sum := 0

	for i := 1; i < n; i++ {
		sum += arr[i]
	}
fmt.Println(sum)
	final :=  sum- total
	return final
}

func main() {
	arr := []int{1, 3,4,5}
	res, ok := missingNum(arr)
	if ok {
		fmt.Println("the missing number is :", res)
	} else {
		fmt.Println("missing num not found:")
	}

	resul := findMissing(arr)
	fmt.Println(resul)
}
