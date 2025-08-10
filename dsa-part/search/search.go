package main

import "fmt"

func findEle(arr []int, n int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return i
		}
	}
	return -1
}
func main() {

	arr := []int{1, 2, 3, 47, 9}

	res := findEle(arr,3)
	fmt.Println(res)
}