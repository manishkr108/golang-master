package main

import "fmt"

func main() {

	// map is unordered collection ofkey valuepairs
	// key must be unique
	// value can be any type
	//Empty Map with make
	mymap := make(map[string]int)
	fmt.Println(mymap)

	m := map[string]int{"banana": 1, "apple": 5}
	m["banana"] = 5
	fmt.Println(m)

	// if i want check key is there or not

	value, key := m["apple"]
	if key {
		fmt.Println("key exist or value", value)
	} else {
		fmt.Println("key not exist")
	}

	//delete key
	delete(m, "apple")
	fmt.Println(m)

	maps := map[string]string{"name": "manish", "lname": "kumar", "pin": "851212"}

	for key, val := range maps {
		fmt.Printf("the keys is %s and value is %s\n", key, val)
	}

	/*Key takeaway:

		count[w] returns the current count for that word (or 0 if not found).

		count[w]++ increases that count by 1 every time the word appears.
	** Summary:**
	words is a slice of string because it’s declared without a fixed size. That’s why you can easily loop over it and even append to it.
	*/

	words := []string{"Go", "Java", "PHP", "GO", "PHP", "GO"}

	count := make(map[string]int)

	for _, val := range words {
		count[val]++
	}

	fmt.Println(count)
	fmt.Printf("%T\n", words)//[]string
	fmt.Printf("%T\n", count)//map[string]int

}
