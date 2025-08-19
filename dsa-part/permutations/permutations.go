package main

import "fmt"

	func permutations(arr []int) [][]int {
	var result [][]int

	var backtrack func(int)

	backtrack = func(start int) {
		if start == len(arr) {
			prim := make([]int, len(arr))

			copy(prim, arr)
			result = append(result, prim)
			return
		}

		for i := start; i < len(arr); i++ {
			arr[start], arr[i] = arr[i], arr[start]
			backtrack(start + 1)
			arr[start], arr[i] = arr[i], arr[start]
		}

	}
	backtrack(0)
	return result

}

func main() {
	arr := []int{1, 2, 3}

	res := permutations(arr)
	fmt.Println(res)

}
