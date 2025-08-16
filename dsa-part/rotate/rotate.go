package main

import "fmt"

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
func rotateArray(arr []int, k int) {
	n := len(arr)
	if n == 0 || k <= 0 {
		return
	}

	k = k % n // handle case when k > n
	fmt.Println(k)

	// Step 1: reverse first k elements
	reverse(arr, 0, k-1)
	// Step 2: reverse the remaining n-k elements
	reverse(arr, k, n-1)
	// Step 3: reverse the entire array
	reverse(arr, 0, n-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotateArray(arr, 3)
	fmt.Println(arr)
}
