Level 0: start = 0
[1, 2, 3]
├── Swap 0 with 0 → [1, 2, 3]
│   └── Level 1: start = 1
│       ├── Swap 1 with 1 → [1, 2, 3]
│       │   └── Level 2: start = 2
│       │       ├── Swap 2 with 2 → [1, 2, 3] ✅
│       ├── Swap 1 with 2 → [1, 3, 2]
│       │   └── Level 2: start = 2
│       │       ├── Swap 2 with 2 → [1, 3, 2] ✅
├── Swap 0 with 1 → [2, 1, 3]
│   └── Level 1: start = 1
│       ├── Swap 1 with 1 → [2, 1, 3]
│       │   └── Level 2: start = 2
│       │       ├── Swap 2 with 2 → [2, 1, 3] ✅
│       ├── Swap 1 with 2 → [2, 3, 1]
│       │   └── Level 2: start = 2
│       │       ├── Swap 2 with 2 → [2, 3, 1] ✅
├── Swap 0 with 2 → [3, 2, 1]
│   └── Level 1: start = 1
│       ├── Swap 1 with 1 → [3, 2, 1]
│       │   └── Level 2: start = 2
│       │       ├── Swap 2 with 2 → [3, 2, 1] ✅
│       ├── Swap 1 with 2 → [3, 1, 2]
│       │   └── Level 2: start = 2
│       │       ├── Swap 2 with 2 → [3, 1, 2] ✅


🌳 Build Recursion Tree

I’ll show as a tree diagram:

backtrack([1,2,3], start=0)
│
├─ i=0 → swap(0,0): [1,2,3]
│   backtrack([1,2,3], start=1)
│   │
│   ├─ i=1 → swap(1,1): [1,2,3]
│   │   backtrack([1,2,3], start=2)
│   │   │
│   │   ├─ i=2 → swap(2,2): [1,2,3]
│   │   │   backtrack([1,2,3], start=3) → ✅ Permutation: [1,2,3]
│   │   │   backtrack ends, undo swap(2,2) → [1,2,3]
│   │   │
│   │   └─ loop ends at start=2
│   │
│   └─ backtrack ends, undo swap(1,1) → [1,2,3]
│
│   ├─ i=2 → swap(1,2): [1,3,2]
│   │   backtrack([1,3,2], start=2)
│   │   │
│   │   ├─ i=2 → swap(2,2): [1,3,2]
│   │   │   backtrack([1,3,2], start=3) → ✅ Permutation: [1,3,2]
│   │   │   undo swap(2,2) → [1,3,2]
│   │   │
│   │   └─ loop ends
│   │
│   └─ undo swap(1,2) → [1,2,3]
│
└─ backtrack ends at start=1

    for i := start; i < len(arr); i++ {
			arr[start], arr[i] = arr[i], arr[start]
			backtrack(start + 1)
			arr[start], arr[i] = arr[i], arr[start]// This is undo part
		}

🔁 Why Undo Is Needed in Backtracking
When we swap elements to explore a new permutation, we’re modifying the original array. If we don’t undo that change, the next recursive path will start with a corrupted or altered array, leading to incorrect results.