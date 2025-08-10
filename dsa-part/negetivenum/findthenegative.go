package main

import "fmt"

func findTheNegative(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			count++
		}
	}

	return count
}

func main() {
	arr := []int{1, 2, 4, 6, 8, 5, 7}
	res := findTheNegative(arr)
	fmt.Print(res)
}
