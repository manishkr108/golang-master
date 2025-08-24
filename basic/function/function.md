Rules for Declaring Generic Functions in Real Projects
1. When logic is identical across multiple types

If the function’s body is the same for int, float64, string, struct, etc., then generics make sense.

Example: Min, Max, Sum, Map, Filter, Reduce functions.

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}


👉 Works for int, float64, string, etc., without code duplication.


When working with collections (slices, maps, channels)

If you often write the same iteration logic for []int, []string, []struct{} → use generics.

func Map[T any, R any](items []T, fn func(T) R) []R {
	result := make([]R, len(items))
	for i, v := range items {
		result[i] = fn(v)
	}
	return result
}


👉 Replaces MapInt, MapString, MapUser boilerplate.

When constraints naturally exist

Use generics when only a specific group of types is valid.

Example: math operations (only numeric types make sense):

type Number interface {
	~int | ~int64 | ~float64
}

func Add[T Number](a, b T) T {
	return a + b
}


👉 Prevents accidental use with unsupported types.

When creating libraries or reusable utilities

In real projects, you don’t want FindInt, FindString, FindUser — one generic Find is better.

func Find[T any](items []T, match func(T) bool) (T, bool) {
	for _, v := range items {
		if match(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}


👉 Reusable across domains.

Avoid generics when:

Business logic is tied to a specific type (e.g., reversing a User struct list is not same as reversing a string).

It adds complexity without benefit.

You’d use type assertions or reflection anyway (bad design → stick to normal functions).

🚦 Rule of Thumb (Interview & Real-World)

Utility/helper functions → ✅ good candidate.

Math or data structure operations (stack, queue, tree, graph, sorting, searching) → ✅ generics shine.

Business logic (invoice calculation, API request handling, domain-specific models) → ❌ keep type-specific.


Rule / Guideline for Declaring Generic Functions

You should define a generic function when:

The logic is identical for multiple types.
Example: reversing a slice of int, string, float64, or rune—the algorithm is the same, only the type changes.

You want type safety + reusability.
Instead of writing separate functions like ReverseInts([]int), ReverseStrings([]string), etc., you can write one Reverse[T any]([]T).

You don’t want to sacrifice performance.
Generics in Go are compiled to concrete types at compile time (no runtime reflection overhead like in interface{}).

Avoiding code duplication.
If you see yourself copying the same logic for multiple types → time to consider generics.



Where You Need Generics in Real Projects

Here are some real-world examples:

1. Utility Functions

Instead of writing:

func MaxInt(a, b int) int { ... }
func MaxFloat(a, b float64) float64 { ... }
func MaxString(a, b string) string { ... }


➡️ Write one generic function:

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}


✅ Works for int, float64, string directly.

2. Data Structures (like LinkedList, Stack, Queue)

Without generics, you’d do:

type IntStack struct { data []int }
type StringStack struct { data []string }


➡️ With generics:

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(val T) {
	s.data = append(s.data, val)
}
func (s *Stack[T]) Pop() T {
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last
}


✅ One stack implementation for all types.


3. Algorithms (Sorting, Searching, Reverse, Map/Filter/Reduce)

Generic reverse:

func Reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}


✅ Works with []int, []string, []rune, []float64.


4. Database / JSON Handling

You often fetch JSON or DB rows into structs. Instead of writing functions for each struct type, use generics:

func DecodeJSON[T any](data []byte) (T, error) {
	var result T
	err := json.Unmarshal(data, &result)
	return result, err
}


✅ Works for User, Product, Order… any struct.