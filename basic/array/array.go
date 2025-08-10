package main

import (
	"fmt"
)

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

	// Declaring an Array
	var intArray [5]int

	// Assigning values
	intArray[0] = 10
	intArray[1] = 2

	// Accessing the elements
	fmt.Println(intArray[0])
	fmt.Println(intArray[1])

	// Initialize Array using Array literals
	x := [5]int{0, 5, 10, 15, 20}
	var y [5]int = [5]int{0, 5, 10, 15, 20}

	fmt.Println(x)
	fmt.Println(y)

	// Initializing an Array with ellipses
	k := [...]int{10, 20, 30}

	fmt.Println(len(k))

	// Initialize values for specific array elements
	a := [5]int{1: 1, 3: 25}

	fmt.Println(a)
}
