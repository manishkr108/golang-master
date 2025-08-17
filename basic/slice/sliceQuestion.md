1. Basic Level – Fundamentals

What is a slice in Go? How is it different from an array?

How do you declare and initialize a slice? Show multiple ways.

What is the zero value of a slice?

How do you get the length and capacity of a slice?

What happens if you access an index out of range in a slice?

Explain the difference between nil slice and an empty slice.

How do you iterate over a slice?

Can slices be compared using == operator? Why/why not?

How do you check if a slice is nil?

What’s the difference between make([]int, 5) and make([]int, 0, 5)?


2. Intermediate Level – Practical Use

How do slices grow when using append?

What happens if you append an element beyond slice capacity?

How do you copy one slice into another? Difference between = and copy()?

How can you remove an element from a slice? (at index i)

How do you implement stack/queue using slices?

Explain how slices share the same underlying array.

Write a function to reverse a slice in place.

What’s the difference between slice[:], slice[0:len(slice)], and slice[0:cap(slice)]?

How do you clone a slice (deep copy)?

What happens when you slice a slice (s[low:high])?


3. Advanced Level – Deep Understanding

Explain the internal structure of a slice (header: pointer, length, capacity).

What happens internally when we do append? (capacity doubling, memory allocation).

What’s the time complexity of append in Go?

What’s the difference between copy() and slicing when duplicating data?

Why can modifying a subslice affect the original slice?

What’s the difference between passing a slice to a function vs an array?

How do you prevent accidental data modification when passing a slice to another function?

Can a slice be used as a map key? Why/why not?

Explain memory leaks possible with slices when handling large arrays.

How does append behave when appending a slice to itself (s = append(s, s...))?


4. Expert Level – Problem Solving with Slices

Implement a function to remove duplicates from a slice (in-place).

Implement a function to rotate a slice left/right by k positions.

Write a function to find the intersection of two slices.

Write a function to merge two sorted slices into one sorted slice.

Implement dynamic resizing using slices (simulate an ArrayList).

Write a function to find all unique pairs in a slice that sum to a target.

Implement an algorithm to move all zeroes to the end of a slice.

Implement a function to check if a slice is a palindrome.

Write a function to chunk a slice into smaller slices of size k.

Write a function to generate permutations of a slice.
