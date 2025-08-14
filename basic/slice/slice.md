What is a Slice?
A slice is like a dynamic array — it can grow or shrink in size.

Slices reference the underlying array, so changes affect the same data.

Creating a Slice with make

nums := male([]int,3)
make([]type, length, capacity)

 Slicing an Array

 array to slice is called slicing

 arr := [5]int{1,2,3,4,5}

 myarr := arr[2:]//elements at index 2, to last

 Slice → Flexible size, can grow with append().


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