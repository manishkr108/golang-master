package main

import "fmt"

func findSingleNum(arr []int) int {
	singleNum := make(map[int]int)

	for _, v := range arr {
		singleNum[v]++
	}

	for k, v := range singleNum {
		if v == 1 {
			return k
		}
	}

	return -1
}

// This is Optimize solution

func singleNumber(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		result ^= arr[i]
	}
	return result
}

func main() {
	arr := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}
	res := findSingleNum(arr)
	fmt.Println(res)

	result := singleNumber(arr)
	fmt.Println("optimize", result)
}
