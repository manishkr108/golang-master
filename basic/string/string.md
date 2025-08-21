📌 What is a String in Go?

👉 “In Go, a string is an immutable sequence of bytes, stored as UTF-8 encoded text. It is not a character array like in C; instead, it is a read-only slice of bytes with a pointer to underlying data and a length field.”


🔎 Internal Representation

A Go string is defined like this under the hood:

type string struct {
    Data uintptr // pointer to the underlying byte array
    Len  int     // length in bytes
}

    s := "Hello, Go!"
    fmt.Println(s)          // Hello, Go!
    fmt.Println(len(s))     // 10 (bytes count)
    fmt.Println(s[0])       // 72 (ASCII of 'H')
    fmt.Println(string(s[0])) // "H"

⚠️ Important Points

Strings are immutable

    s := "Hello"
    // s[0] = 'h'  // ❌ Compile error

UTF-8 encoded
Go strings support Unicode (multi-byte chars).

    str := "Go भाषा"
    fmt.Println(len(str)) // 11 (bytes, not chars!)
Runes vs Bytes

    len(str) → counts bytes

    utf8.RuneCountInString(str) → counts characters (runes)

    import "unicode/utf8"

    str := "Go भाषा"
    fmt.Println(len(str))                      // 11
    fmt.Println(utf8.RuneCountInString(str))   // 6

📌 If String is Not an Array, How Does len(str) Work?

In Go, a string is not an array of characters, it is:

    type string struct {
    Data uintptr  // pointer to the underlying byte array
    Len  int      // length in bytes
    }

🔎 How len(str) Works

len(str) does not count characters one by one.

It simply returns the stored Len field of the string structure.

That’s why it runs in O(1) time complexity (constant time).

    s := "Go भाषा"

    fmt.Println(len(s))                    // 11 (bytes)
    fmt.Println(utf8.RuneCountInString(s)) // 6 (characters/runes)

    // Internally
    // "Go भाषा" in UTF-8 bytes:
    // [71 111 32 224 164 174 224 164 184 224 164 184]

Here:

"Go भाषा" looks like 6 characters.

But in memory it’s stored in 11 bytes.

len(s) = 11, because Go counts bytes not characters.

Key Difference from Array

Array → fixed-size collection, len(arr) actually counts elements.

String → stored with a Len field, so len(str) just returns that field.

***Bytes vs Runes in Go Strings***

type string struct {
    Data uintptr  // pointer to the bytes
    Len  int      // length in bytes
}

len(str) → returns bytes count (Len field).

But text is UTF-8 encoded, meaning one character may take multiple bytes.

Example: ASCII String
    s := "Hello"
    fmt.Println(len(s)) // 5

**Memory (bytes):**

[72 101 108 108 111]  
 H   e   l   l   o  


Each char = 1 byte.

len(s) = 5 → matches number of characters.

**Example: Unicode String**
s := "भा"
fmt.Println(len(s))                    // 6 (bytes)
fmt.Println(utf8.RuneCountInString(s)) // 2 (runes/characters)

Memory (UTF-8 bytes):

[224 164 174   224 164 184]  
   भ (3 bytes)     ा (3 bytes)

"भा" looks like 2 letters.

Stored in 6 bytes.

len(s) = 6 (bytes).

utf8.RuneCountInString(s) = 2 (characters).


Iterating with Bytes vs Runes

s := "Go भाषा"

// By bytes
for i := 0; i < len(s); i++ {
    fmt.Printf("%d: %x\n", i, s[i])
}

// By runes
for i, r := range s {
    fmt.Printf("%d: %c\n", i, r)
}