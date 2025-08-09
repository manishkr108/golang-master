package main

import "fmt"

func main() {
	// var arr [5]int
	arr := [5]int{10, 20, 30}
	arr[2] = 100
	fmt.Println(arr)

	ar := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(ar))

	fmt.Println(ar)

	for index, value := range ar {
		fmt.Printf("Index : %d => Value %d\n", index, value)
	}
	// multi dimational array
	matrix := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{11, 12, 13, 14},
	}
	fmt.Println(matrix)
}
