package main

import (
	"fmt"
)

func main() {
	//1. Basic for loop (like C-style)

	for i := 0; i < 5; i++ {

		fmt.Println(i)
	}

	//2. for as a while loop
	i := 0
	for i < 5 {
		fmt.Println("for like while loop", i)
		i++

	}

	//Looping over arrays/slices
	arr := []string{"GO", "PHP", "JAVA"}

	for i, val := range arr {
		fmt.Println("I is index and val is value", i, val)
	}

	//. Using continue and break
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		if i == 7 {
			break
		}

		fmt.Println(i)
	}

	//Reverse the number

	var num int = 1234
	var rev int = 0
	for num > 0 {
		digit := num % 10
		rev = rev*10 + digit
		num = num / 10
	}
	fmt.Println(rev)

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j + 1)
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print(j + 1)
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= 5-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= 5-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for m := 5; m > 5-i; m-- {
			fmt.Print(" ")
		}
		for j := 5; j >= 2*i-1; j-- {
			fmt.Print("*")
		}

		for k := 0; k < 5-i; k++ {
			fmt.Print(" ")
		}

		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= i; k++ {
			fmt.Print(k)
		}

		for m := i - 1; m >= 1; m-- {
			fmt.Print(m)
		}

		fmt.Println()
	}

	count := 1
	for i := 1; i < 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		ch := 'A' + i
		for j := 0; j <= i; j++ {
			fmt.Print(string(ch), " ")
		}
		fmt.Println()
	}
	for i := 0; i < 5; i++ {
		cha := 'A'
		for j := 0; j <= i; j++ {
			fmt.Print(string(cha))
			cha++
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 || i == 5-1 || j == 0 || j == 5-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i-1; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < 2*i+1; j++ {
			if i == 0 || i == 5-1 || j == 0 || j == 2*i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	for i := 1; i <= 10; i++ {
		fmt.Print(5 * i)
		fmt.Println()
	}

	nums := 12345678
	counts := 0
	sum := 0
	for nums > 0 {
		digit := nums % 10
		sum = sum + digit
		nums = nums / 10
		counts++

	}

	fmt.Print("the number of digit count ", counts)
	fmt.Print("the sum of digit ", sum)

}
