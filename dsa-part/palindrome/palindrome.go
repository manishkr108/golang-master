package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func palindrom(num int) bool {
	// The line `//org := num` is a commented-out line in the code. It is not being used for any
	// functionality in the program. It appears to be a commented-out variable declaration that was not
	// removed from the code. This line is ignored by the Go compiler and has no impact on the program's
	// execution.
	org := num
	rem := 0
	reverce := 0
	for num > 0 {
		rem = num % 10
		reverce = reverce*10 + rem
		num = num / 10
	}

	return reverce == org

}

func strpalindrom(str string) bool {
	str = strings.ToLower(str)
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}

	}
	return true
}

func main() {
	res := strpalindrom("mam")
	fmt.Println(res)
	var num int
	fmt.Println("Please enter a number")
	fmt.Scan(&num)
	if palindrom(num) {
		fmt.Println(num, "is a palindrom")
	} else {
		fmt.Println(num, "not a palindrom")
	}

	//scanner is a new variable , using short hand syntex :=
	//bufio.NewScanner=> this create new Scanner object from Go bufio package
	//os.Stdin =>This repersent the standard input (keybord terminal input)

	//Create a scannerthat will read input from user keyboard
	scanner := bufio.NewScanner(os.Stdin)
	//fmt.Printf("%+v\n", scanner)
	//spew.Dump(scanner)
	fmt.Println("Please enter number to check palindrome")
	scanner.Scan()                // wait for user input and hit enter
	input := scanner.Text()       // get the input as a String
	word := strings.Fields(input) // split a string into separate words based on any whitespace,It returns a slice of strings ([]string), where each element is one word.
	fmt.Print(word)
	result := ""
	for _, ch := range input {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			result += string(unicode.ToLower(ch))
		}
	}
	reverce := ""
	for i := len(result) - 1; i >= 0; i-- {
		reverce += string(result[i])
	}

	if reverce == input {
		fmt.Print("Input is palindrom ", reverce)
	} else {
		fmt.Print("Input Not a palindrom ", input)
	}
}
