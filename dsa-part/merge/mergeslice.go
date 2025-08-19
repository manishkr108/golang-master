package main

import "fmt"

func mergeTwoSlice(arr1 []int, m int, arr2 []int, n int) {

	i, j, k := m-1, n-1, m+n-1

	fmt.Println(i, j, k)

	for j >= 0 {
		if i >= 0 && arr1[i] > arr2[j] {
			arr1[k] = arr1[i]
			i--
		} else {
			arr1[k] = arr2[j]
			j--
		}
		k--
	}
}

func main() {

	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 4, 6}
	mergeTwoSlice(arr1, 3, arr2, 3)
	fmt.Println(arr1)

}
