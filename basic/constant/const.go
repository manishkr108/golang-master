package main

import "fmt"

const pi = 3.14
const appName string = "GoMaster"

func main() {
	//What is a Constant?
	/*
						   A constant is a named value that:

						   Is set at compile time

						   Never changes

						   Is typically used for values like limits, labels, fixed settings, etc.

						   Single constant
					const pi = 3.14
					const appName string = "GoMaster"

					2. Multiple constants
			const (
			    maxUsers = 100
			    minAge   = 18
			    version  = "v1.0"
			)

			iota for Enum-like Constants

			const (
	    Admin = iota  // 0
	    Manager       // 1
	    User          // 2
	)
	*/
	fmt.Println("App:", appName)
	fmt.Println("Value of Pi:", pi)

	// Rules for Constants
	//You cannot assign a new value to a constant.
	//You cannot use variables or runtime functions to initialize a constant:
	//const x = math.Sqrt(4)  // ❌
}
