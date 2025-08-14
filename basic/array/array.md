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
