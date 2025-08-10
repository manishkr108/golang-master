package main

import "fmt"

// 👇 Global variable
//Use global variables only when necessary (e.g., config, constants, shared state).
var globalMessage string = "Hello from Global!"

func main() {
	//Variables declared with var can be used outside functions (package level).
	//The := short syntax only works inside functions.

	//Using var keyword with type

	var name string = "Manish"
	age := 25
	var isAdmin = true
	var foo, bar string = "Hello", "World"
	fmt.Println(foo, bar)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Admin:", isAdmin)

	// Grouped declaration
	var (
		product string = "Manish"
		price   int    = 25
		user    bool   = true
	)

	fmt.Println(product, price, user)
	fmt.Println(globalMessage)
}
