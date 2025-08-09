package main

import "fmt"

func findTheLarge(arr []int) int {
	large := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > large {
			large = arr[i]
		}
	}

	return large
}

func findminival(arr []int) int {
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}
func main() {
	arr := []int{11, 44, 66, 3, 23, 56, 8, 3}
	res := findTheLarge(arr)
	fmt.Println("the maximim value is ", res)

	min := findminival(arr)
	fmt.Print("the minimum value is ", min)
}
