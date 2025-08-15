package main

import (
	"fmt"
)

func sumArray(arr []int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}

	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	res := sumArray(arr)
	fmt.Println(res)
}
