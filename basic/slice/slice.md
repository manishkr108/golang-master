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