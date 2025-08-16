array is a fixed size collaction of same elements.
Once you set the size, you can’t change it.

b:=2

if you not using b it will throw error like b not used 
avoid this error using _=b

1. Array in Memory
arr := [5]int{10, 20, 30, 40, 50}

Memory layout:

Index:    0     1     2     3     4
Value:   10    20    30    40    50
Address: 100  104   108   112   116   (each int = 4 bytes, contiguous)


Fixed length → exactly 5 integers.

When assigned to another array, full copy is made.

2. Slice in Memory
nums := []int{10, 20, 30}

Under the hood, a slice is:
+------------+-----------+-----------+
| Pointer -> | Length=3  | Capacity=3|
+------------+-----------+-----------+

Pointer points to underlying array:

Underlying Array:
Index:    0     1     2
Value:   10    20    30
Address: 200  204   208


Pointer → refers to an underlying array.

Length → number of elements in the slice.

Capacity → space available in underlying array from start index.

When appended beyond capacity → new underlying array is created.

Key Differences:
Feature 	    Array	                                        Slice
Length	        Fixed at creation	                            Can grow or shrink
Syntax	        [n]Type	                                        []Type
Memory          Allocation	Contiguous,                         fixed size	References an underlying array
Value Type	    Yes (copy on assignment)	                    Reference type (shares data)
Usage	        Rare in high-level code	Very common in Go programs
Performance	    Slightly more predictable	                     More flexible, minimal overhead

How do you declare an array in Go? Give examples of fixed-size and inferred-size arrays.

var name [length]Type
length → must be a constant (cannot change at runtime).

Type → data type of elements (all elements must be the same type).

var nums [3]int // fixed size

2. Inferred-size array
You let Go count the elements using ...


func main() {
	arr := [...]int{1, 2, 3, 4, 5} // Go counts → length is 5
	fmt.Println(arr)               // [1 2 3 4 5]
	fmt.Println(len(arr))          // 5
}

[...]int tells Go: “figure out the size based on the number of elements.”

var a [5]int                   // empty array of size 5 → [0 0 0 0 0]
b := [3]float64{1.1, 2.2, 3.3} // fixed-size with initialization
c := [...]bool{true, false}    // inferred-size → length 2

What is the default value of an uninitialized array in Go?

    var numbers [5]int      // array of int, uninitialized
    var names [3]string     // array of string, uninitialized
    var flags [4]bool       // array of bool, uninitialized

    fmt.Println(numbers) // [0 0 0 0 0]
    fmt.Println(names)   // [  "" "" "" ]
    fmt.Println(flags)   // [false false false false]


     numbers := [5]int{0: 100, 3: 200}
    fmt.Println(numbers) // [100 0 0 200 0]

    💡 Key points for beginners:

Arrays have fixed length — you can’t add/remove elements.

Missing values are automatically filled with the zero value.

[...] lets Go figure out the length from the initializer.

How do you get the length of an array in Go?

How do you get the length of an array in Go?

💡 Key points:
len(array) returns the number of elements in the array.

Works for arrays, slices, strings, maps, and channels.

For strings, len() returns the number of bytes, not runes (characters).

✅ Real-world use case:
You might use len() to loop over an array:

go
Copy
Edit
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}

How do you compare two arrays for equality in Go?
In Go, you can compare two arrays directly using the == operator, but there are a few important rules.
    a := [3]int{1, 2, 3}
    b := [3]int{1, 2, 3}
    c := [3]int{3, 2, 1}
    d := [4]int{1, 2, 3, 4} // different length

    fmt.Println(a == b) // true  (same length and same elements in order)
    fmt.Println(a == c) // false (elements differ)
    // fmt.Println(a == d) // ❌ compile-time error: mismatched array sizes

Why this works for arrays but not slices
Arrays in Go are value types — comparisons work like comparing structs.

Slices are reference types — you can’t compare them with == except checking if both are nil.
If you need to compare slices, you can use:
import "reflect"
reflect.DeepEqual(slice1, slice2)
or manually loop through elements.

When to Use Which
Manual function → Faster, more control, and no extra imports.

reflect.DeepEqual → Easy, works for any type (but a bit slower).


1. What Does “Contiguous Memory” Mean?
All the elements of an array are stored next to each other in memory, with no gaps in between.

If you know the memory address of the first element, you can calculate the address of any element with:

address_of_element = base_address + (index × size_of_element)