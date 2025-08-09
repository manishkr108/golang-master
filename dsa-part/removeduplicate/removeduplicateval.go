package main

import "fmt"

func removeduplicate(arr []int) []int {
	x := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[x] {
			x++
			arr[x] = arr[i]
		}
	}
	fmt.Println("total unique val is ", x+1)
	return arr
}

func removeEle(arr []int) []int {
	n := 3
	x := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != n {
			arr[x] = arr[i]
			x++
		}
	}

	return arr
}
func main() {
	arr := []int{0, 0, 1, 2, 2, 2, 3, 3, 3, 4, 5, 6, 6, 7}
	arr1 := []int{1, 2, 3, 4, 5, 3, 3, 3, 6, 7, 3, 9}

	inplaceres := removeduplicate(arr)
	fmt.Println(inplaceres)
	res := removeEle(arr1)
	fmt.Print(res)
}
