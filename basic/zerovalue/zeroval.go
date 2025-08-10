package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	//Type Conversion
	a := 42
	bb := float64(a)
	c := uint(bb)

	fmt.Printf("%T %T", bb, c)
}
