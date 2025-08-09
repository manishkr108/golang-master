package main

import (
	"fmt"
)

func main() {
	var a []int
	fmt.Println(a)
	// Slicing an Array
	arr := []int{1, 23, 4, 5, 6}

	myar := arr[1:]

	fmt.Print(myar)

	mya := arr[:2]
	mya = append(mya, 44)
	fmt.Print(mya)

	//
	// Creating a Slice with make

	nums := make([]int, 3)
	nums = append(nums, 1, 2, 3, 4)
	fmt.Println(len(nums), cap(nums))

	//Copy slice

	aa := []int{1, 2, 3, 4, 5}

	bb := make([]int, len(aa))

	copy(bb,aa)

	fmt.Println(bb);


}
