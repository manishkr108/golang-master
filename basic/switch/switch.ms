 age := 'm'
 'm' actually has the integer value 109 (ASCII code for m).
 age := 109


 Normally Go stops after the first matching case.
If you want to continue to the next case, use fallthrough.

package main

import "fmt"

func main() {
    num := 2

    switch num {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two")
        fallthrough
    case 3:
        fmt.Println("Three")
    }
}
output
Two
Three