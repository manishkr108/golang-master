package main

import "fmt"

func findPairs(nums []int, target int) [][2]int {
	seen := make(map[int]bool)
	used := make(map[[2]int]bool)
	result := make([][2]int, 0)

	for _, num := range nums {
		complement := target - num
		if seen[complement] {
			pair := [2]int{num, complement}

			if num > complement {
				pair[0], pair[1] = complement, num
			}
			if !used[pair] {
				result = append(result, pair)
				used[pair] = true
			}
		}
		seen[num] = true
	}
	return result
}

func main() {
	arr := []int{1, 2, 3, 2, 4, 5, 2}
	res := findPairs(arr, 5)
	fmt.Println(res)
}
