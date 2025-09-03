package main

import "fmt"

func moveZero(arr []int) []int {
	x := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			temp := arr[x]
			arr[x] = arr[i]
			arr[i] = temp
			x++
		}
	}

	return arr
}

func main() {
	arr := []int{0, 1, 3, 0, 0, 6, 7, 4, 0, 0}
	res := moveZero(arr)
	fmt.Println(res)
}
