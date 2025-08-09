package main

import "fmt"

func main() {

	age := 'm'
fmt.Print(age)
	switch {
	case age < 18:
		fmt.Println("not eligibal for vote")
	case age > 18:
		fmt.Println("Now you are eligibal for vote")
	default:
		fmt.Print("Please enter valid age")
	}
}
