package main

import "fmt"

func secMax(arr []int) int {
	large := arr[0]
	seclarge := arr[1]

	for i := 0; i < len(arr); i++ {
		if arr[i] > large {
			seclarge = large
			large = arr[i]
		} else if arr[i] > seclarge && large < arr[i] {
			seclarge = arr[i]
		}
	}
	return seclarge
}
func main() {

	arr := []int{1, 5, 3, 7, 12, 45, 34, 6, 89, 4}

	res := secMax(arr)
	fmt.Println("the secound max value is ", res)
}
