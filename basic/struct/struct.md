What is a struct?
A struct (short for structure) is a custom data type in Go that lets you group different pieces of related data together under one name.

Think of it like a container that can hold variables of different types.

Why use a struct?
Without structs, you might have to keep track of many separate variables:

name := "Alice"
age := 25
salary := 50000.0

 Syntax of a struct

 type StructName struct {
    fieldName fieldType
    fieldName fieldType
}

for key, val := range something works only on iterable types in Go:

Arrays

Slices

Maps

Strings

Channels

A struct is not iterable — it’s just a collection of named fields in memory.
Go doesn’t give you the field names automatically in a loop.