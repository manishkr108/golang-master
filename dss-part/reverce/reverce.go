package main

import "fmt"

func reverce(num int) int {
	rem := 0
	rev := 0
	for num > 0 {
		rem = num % 10
		rev = (rev * 10) + rem
		num = num / 10
	}

	return rev
}
func main() {
	res := reverce(234)
	fmt.Println("the reverce numer is ", res)
}