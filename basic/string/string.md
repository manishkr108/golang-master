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


How do you efficiently concatenate multiple strings in a loop? 
Why is strings.Builder better than +?

Problem with + (string concatenation)

Strings in Go are immutable.
When you do this:

result := ""
for i := 0; i < 5; i++ {
    result += "x"
}


What actually happens:

Each += creates a new string (copies old + new).

Time complexity becomes O(n²) for big concatenations (lots of re-copying).

Inefficient in memory usage.

Solution: strings.Builder

Go provides strings.Builder
 for efficient concatenation.

    var sb strings.Builder

    for i := 0; i < 5; i++ {
        sb.WriteString("x")
    }

    result := sb.String()
    fmt.Println(result) // "xxxxx"

Why strings.Builder is better:

No repeated copying: It grows an internal buffer (amortized O(n)).

Efficient memory use: You can Grow(n) upfront to avoid resizing.

Safe to use: Once you call String(), you shouldn’t reuse it, ensuring correctness.

Real-World Use Cases
1. Building Dynamic SQL Queries
var b strings.Builder
b.WriteString("SELECT * FROM users WHERE 1=1")

if active := true; active {
	b.WriteString(" AND status = 'active'")
}
if country := "IN"; country != "" {
	b.WriteString(" AND country = '")
	b.WriteString(country)
	b.WriteString("'")
}

query := b.String()
fmt.Println(query)




Why len(str) is O(1)

Recall the internal structure of a string in Go:

type stringStruct struct {
    str unsafe.Pointer  // pointer to the byte array
    len int             // number of bytes
}


The length is already stored in the len field.

When you call len(str), Go simply returns that integer field — it doesn’t need to count bytes or scan memory.

That’s why len(str) is constant time O(1).


Example (what happens internally)
s := "Hello 😀"
fmt.Println(len(s))   // 8 bytes


At runtime:

s = (pointer, len=8)

len(s) just returns the len field → 8.

No traversal of the string happens.

Rule of thumb:

If you’re working with bytes, it’s usually O(1).

If you’re working with runes/characters, it’s usually O(n) (because of UTF-8 decoding).

What happens if you modify a []byte obtained from a string?
s := "hello" b := []byte(s) b[0] = 'H' fmt.Println(string(b)) // ???

    s := "hello"
    b := []byte(s)  // copy string bytes into new slice
    b[0] = 'H'      // modify the first byte
    fmt.Println(string(b)) 

    Step 1: s := "hello"

In Go, a string is immutable. You cannot modify its bytes directly.

Internally: "hello" is stored as a read-only byte array ([104 101 108 108 111] in ASCII).

Step 2: b := []byte(s)

This creates a copy of the string's bytes into a new []byte slice.

So b = [104 101 108 108 111] (the same bytes, but mutable).

Step 3: b[0] = 'H'

'H' has ASCII value 72.

Now b = [72 101 108 108 111] → which corresponds to "Hello".

Step 4: fmt.Println(string(b))

Converts the modified []byte back into a string.

Key point:

The original string s is untouched (still "hello").

The slice b is a new copy, so modifying it doesn’t break immutability.

