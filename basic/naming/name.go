package main

import "fmt"

func main() {
	//Capitalization Determines Visibility
	//Capitalized names are exported (public – accessible outside the package).

	//Lowercase names are unexported (private – only accessible within the same package).
	//Short, Clear, and Descriptive Names
	//Avoid over-abbreviation (c is bad, config is better).
	//Don't write unnecessary prefixes (getUserName() → just UserName()).
	firstName := "Manish"
	userAge := 25
	fmt.Println("Name:", firstName)
	fmt.Println("Age:", userAge)
}
