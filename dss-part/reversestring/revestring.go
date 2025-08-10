package main

import (
	"fmt"
)

func reverseString(str []string) []string {
	leng := len(str)

	for i := 0; i < leng/2; i++ {
		temp := str[i]
		str[i] = str[leng-1-i]
		str[leng-1-i] = temp
	}

	return str
}

func main() {
	str := []string{"m", "a", "n","i"}

	res := reverseString(str)
	fmt.Print(res)
}
