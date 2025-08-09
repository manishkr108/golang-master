package main

import "fmt"

func fibbonacci(num int) int {

	if num == 0 || num == 1 {
		return num
	}

	return fibbonacci(num-1) + fibbonacci(num-2)

}

func main() {

	n := 10
	sum := 0
	for i := 0; i <= n; i++ {

		fmt.Print(fibbonacci(i), " ")
		sum += fibbonacci(i)
	}

	fmt.Print("total sum of fibbonacci ", sum)

}
