package main

import "fmt"

//Write a Go function to reverse an array in place.
func reverceInplace(arr [5]int) [5]int {

	lan := len(arr)

	for i := 0; i < lan/2; i++ {
		temp := arr[i]
		arr[i] = arr[lan-i-1]
		arr[lan-i-1] = temp
	}
	return arr
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	result := reverceInplace(arr)
	fmt.Println("org", arr)
	fmt.Println(result)
}
