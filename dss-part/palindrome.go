package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter number to check palindrome")
	scanner.Scan()
	input := scanner.Text()
	result := ""
	for _, ch := range input {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			result+= string(unicode.LowerCase)
		}
	}
}
