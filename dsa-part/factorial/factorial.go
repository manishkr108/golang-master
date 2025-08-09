package main

import "fmt"

//factorial cllled recursion
func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}

	return num * factorial(num-1)//5*4*3*2*1
}

func main() {
	res := factorial(5)
	fmt.Print(res)
}
