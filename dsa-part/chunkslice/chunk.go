package main

import "fmt"

func chunkSlice(arr []int, size int) [][]int {
	var result [][]int

	for i := 0; i < len(arr); i += size {
		end := i + size

		if end > len(arr) {
			end = len(arr)
		}

		result = append(result, arr[i:end])
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	res := chunkSlice(arr, 2)

	fmt.Println(res)
}
