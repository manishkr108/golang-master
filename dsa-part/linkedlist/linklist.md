Write a function that takes in the head of a Singly Linked List and an integer k, shifts the list in place (i.e., doesn't create a brand new list) by k positions, and returns its new head.

if k=2 1. I want to reverse list 0 to n exm => 5->4->3->2->1->0 2. next swtape i have to reverse k position like 4->5-3-2-1-0 3. the i reverce k posint to n position like 4->5->0->1->2->


first we need to create a Node type struct 
1. Node struct
type Node struct{
    value int
    next *Node
}

🔹 Why this is needed?

Every linked list is built from nodes.

Each node must store two things:

value → the actual data (like 10, 20, 30).

next → a pointer that tells where the next node is in memory.

👉 Without next, the nodes would not be linked. That’s why it’s called a linked list.

secound point is you need to create LinkedList see belo

type LinkedList struct{
    head *Node
}


2. Create LinkedList struct
type LinkedList struct {
    head *Node
}


🔹 Why this is needed?

A linked list must have an entry point.

That entry point is the head.

head points to the first node in the list.

If the list is empty → head = nil.

👉 Think of head as the front door of your list. Without it, you can’t reach any nodes.

3. Add function to insert at the end
func (list *LinkedList) Insert(value int) {
    newNode := &Node{value: value}

    if list.head == nil {
        // if list is empty, new node becomes head
        list.head = newNode
        return
    }

    curr := list.head
    for curr.next != nil {  // go until last node
        curr = curr.next
    }

    // connect last node to new node
    curr.next = newNode
}

Line by line explanation:

newNode := &Node{value: value}
→ Create a fresh node in memory with given value. Initially next = nil.

if list.head == nil { ... }
→ If no nodes exist, make the new node the first node (head).
(Because the first inserted node is always the head).

curr := list.head
→ Start from the head (entry point).

for curr.next != nil { curr = curr.next }
→ Walk through all nodes until the last one (the one whose next == nil).

curr.next = newNode
→ Attach the new node at the end by updating the last node’s next.

4. Function to display list
func (list *LinkedList) Display() {
    curr := list.head
    for curr != nil {
        fmt.Print(curr.value)
        if curr.next != nil {
            fmt.Print(" -> ")
        }
        curr = curr.next
    }
    fmt.Println()
}


🔹 Line by line explanation:

curr := list.head
→ Start from the head.

for curr != nil
→ Visit each node until you fall off the list (nil).

fmt.Print(curr.value)
→ Print the data of the current node.

if curr.next != nil { fmt.Print(" -> ") }
→ Print arrow only if another node exists after this.

curr = curr.next
→ Move to the next node.

fmt.Println()
→ Print new line at the end for neat output.

5. Main function
func main() {
    list := LinkedList{}

    list.Insert(10)
    list.Insert(20)
    list.Insert(30)

    list.Display()
}


