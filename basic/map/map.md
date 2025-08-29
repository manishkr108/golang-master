What is a Map in Go?

A map is Go’s built-in hash table implementation.

 map is unordered collection of key value pairs
 key must be unique
 value can be any type



syntex 
Empty Map with make
mymap := make(map[string]int)


 Why Maps Are Important
Quick access → You can get a value instantly by its key, no need to loop through everything.

Unique keys → Each key is stored only once.

Dynamic size → You can add/remove pairs without worrying about length.

Flexible → Keys can be string, int, etc., and values can be any type.


✅ Create Maps
m := make(map[string]int) // Empty map
m2 := map[string]int{"a": 1, "b": 2} // With values

✅ Insert / Update

m["go"] = 100
m["go"] = 200 // Updates value

✅ Access
value := m["go"]        // 200
value2 := m["missing"]  // 0 (zero-value if key not found)

✅ Check Key Existence

v, ok := m["missing"]
if ok {
    fmt.Println("Found:", v)
} else {
    fmt.Println("Not found")
}

🔑 3. Iterating a Map

for k, v := range m {
    fmt.Println(k, v)
}

🔑 4. Delete Keys

delete(m, "go")


🔹 Go Map Internals (runtime/map.go)

In Go, a map is not just a simple array → it’s a hash table with buckets.
The compiler makes your map[...]... type, but under the hood the runtime uses a private struct called hmap.
***Map Internals***

        type hmap struct {
            count     int    // number of elements in the map
            flags     uint8
            B         uint8  // log2(number of buckets)
            noverflow uint16 // number of overflow buckets
            hash0     uint32 // hash seed (randomized)

            buckets    unsafe.Pointer // array of buckets (size 2^B)
            oldbuckets unsafe.Pointer // previous bucket array (during grow)
            nevacuate  uintptr        // progress counter for growing

            extra *mapextra // optional fields (for overflow handling)
        }

🔹 Each Bucket

A bucket in Go holds up to 8 key-value pairs.
Struct (simplified):

    type bmap struct {
        tophash [8]uint8       // top 8 bits of each key's hash
        keys    [8]keytype     // actual keys
        values  [8]valuetype   // values
        overflow *bmap         // pointer to another bucket if full
    }
So the map is really:

hmap → pointer to buckets (bmap array)

Each bmap → up to 8 slots (keys + values + hash top bits)

 m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

    // reflect to get runtime pointer
    header := *(*uintptr)(unsafe.Pointer(&m))

    fmt.Printf("Map header pointer: %#x\n", header)

    // See full hmap structure
    type hmap struct {
        count     int
        flags     uint8
        B         uint8
        noverflow uint16
        hash0     uint32
        buckets   uintptr
        oldbuckets uintptr
        nevacuate uintptr
        extra     uintptr
    }

    // Cast pointer to hmap
    h := (*hmap)(unsafe.Pointer(header))

    fmt.Printf("hmap: %+v\n", *h)

Example Output (differs by run):
    Map header pointer: 0xc0000a6000
    hmap: {count:3 flags:0 B:1 noverflow:0 hash0:279384682 buckets:824634848288 oldbuckets:0 nevacuate:0 extra:0}

Recap: Map internals

A Go map is backed by an hmap struct.

hmap.buckets points to an array of bmap structs.

Each bmap can store up to 8 key/value pairs.

If more than 8 go into the same bucket → it chains into overflow buckets.

Example 1: Very Simple Custom Map (array of key-value pairs)    

package main

import "fmt"

// KeyValue represents one entry
type KeyValue struct {
	Key   string
	Value int
}

// MyMap is our custom map (backed by a slice)
type MyMap struct {
	data []KeyValue
}

// Set adds or updates a key-value pair
func (m *MyMap) Set(key string, value int) {
	for i, kv := range m.data {
		if kv.Key == key {
			m.data[i].Value = value // update existing
			return
		}
	}
	m.data = append(m.data, KeyValue{Key: key, Value: value})
}

// Get retrieves a value by key
func (m *MyMap) Get(key string) (int, bool) {
	for _, kv := range m.data {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return 0, false
}

// Delete removes a key
func (m *MyMap) Delete(key string) {
	for i, kv := range m.data {
		if kv.Key == key {
			m.data = append(m.data[:i], m.data[i+1:]...)
			return
		}
	}
}

func main() {
	m := &MyMap{}
	m.Set("apple", 10)
	m.Set("banana", 20)
	m.Set("apple", 15) // update

	val, ok := m.Get("apple")
	fmt.Println("apple:", val, ok) // apple: 15 true

	m.Delete("banana")
	fmt.Println(m.data) // [{apple 15}]
}

Example 2: Custom Hash Map (Buckets + Hash Function)
package main

import (
	"fmt"
	"hash/fnv"
)

type KeyValue struct {
	Key   string
	Value int
}

type Bucket struct {
	items []KeyValue
}

type MyHashMap struct {
	buckets []Bucket
}

func NewMyHashMap(size int) *MyHashMap {
	return &MyHashMap{buckets: make([]Bucket, size)}
}

func (m *MyHashMap) hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % len(m.buckets)
}

func (m *MyHashMap) Set(key string, value int) {
	index := m.hash(key)
	for i, kv := range m.buckets[index].items {
		if kv.Key == key {
			m.buckets[index].items[i].Value = value
			return
		}
	}
	m.buckets[index].items = append(m.buckets[index].items, KeyValue{Key: key, Value: value})
}

func (m *MyHashMap) Get(key string) (int, bool) {
	index := m.hash(key)
	for _, kv := range m.buckets[index].items {
		if kv.Key == key {
			return kv.Value, true
		}
	}
	return 0, false
}

func main() {
	m := NewMyHashMap(5)
	m.Set("apple", 10)
	m.Set("banana", 20)
	m.Set("cherry", 30)

	val, ok := m.Get("banana")
	fmt.Println("banana:", val, ok) // banana: 20 true

	fmt.Println(m.buckets) // You can see buckets + collisions
}


TL;DR

Maps are hash tables with buckets.

Each bucket holds up to 8 key/value pairs + tiny “top-hash” bytes.

Iteration order is intentionally non-deterministic—don’t rely on it.

Resizing is incremental (no big stop-the-world rehash).

Maps are not safe for concurrent writes (use a mutex or sync.Map).

The mental model (simple, accurate enough)

hmap (the header)
 ├─ count         // # of entries
 ├─ B             // log2(# buckets) → buckets = 2^B
 ├─ hash0         // per-map hash seed
 └─ buckets ──▶ [ b0, b1, b2, ... ]

bmap (one bucket)
 ├─ tophash[8]    // top 8 bits of each key’s hash
 ├─ keys[8]       // 0..7
 ├─ values[8]     // 0..7
 └─ overflow ──▶  next bucket if this one fills
Why 8? It’s a sweet spot for cache locality vs collision handling.

How a key lands in a bucket

1.Compute hash(key) using a per-map seed (hash0).

2.Take the low B bits → pick bucket index.

3.Scan up to 8 slots; use tophash to quickly skip non-matches.

4.If full, follow overflow chain.

Why iteration looks “random”

When you do:
for k, v := range m { /* ... */ }
Go does not guarantee order. The runtime uses randomized start positions and probing so you never (accidentally) depend on order. That keeps your code correct and reduces certain attack surfaces.


Proof (run twice):
m := map[string]int{"a":1, "b":2, "c":3, "d":4}
for k, v := range m { fmt.Println(k, v) }


You’ll see different orders across runs.

Need order? Sort the keys:

keys := make([]string, 0, len(m))
for k := range m { keys = append(keys, k) }
sort.Strings(keys)
for _, k := range keys { fmt.Println(k, m[k]) }


Resizing without pain (incremental grow)

When the load gets high, Go doubles buckets (B++).
But it doesn’t rehash everything at once. Buckets are evacuated lazily as you touch them. That keeps latency smooth.


Practical pro tips

✅ Key uniqueness: setting the same key overwrites the value.

✅ Missing keys return zero value:

Takeaway

Think of Go’s map as a carefully tuned hash table: buckets of 8, randomized iteration, incremental growth, and great cache behavior. Use it freely; sort only when you truly need order; guard it in concurrent writes.

If this helped, drop a “map ❤️” in the comments and I’ll post a follow-up with visuals of buckets, overflow, and tophash!