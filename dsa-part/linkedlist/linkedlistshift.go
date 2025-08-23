//Write a function that takes in the head of a Singly Linked List and an integer k, shifts the list in place (i.e., doesn't create a brand new list) by k positions, and returns its new head.
/*
if k=2 1. I want to reverse list 0 to n exm => 5->4->3->2->1->0 2. next swtape i have to reverse k position like 4->5-3-2-1-0 3. the i reverce k posint to n position like 4->5->0->1->2->3
*/

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) append(value int) {
	newNode := &Node{value: value}
	newNode.next = list.head
	list.head = newNode

}

func (list *LinkedList) print() {
	curr := list.head
	for curr != nil {
		fmt.Print(curr.value)
		if curr.next != nil {
			fmt.Print("->")
		}

		curr = curr.next

	}

	fmt.Println("->nil")
}

func (list *LinkedList) shift(k int) {

	length := 1

	tail := list.head

	for tail.next != nil {
		tail = tail.next
		length++
	}

	// k is grater than length or less than zero
	k = k % length

	newTail := list.head

	for i := 0; i < length-k-1; i++ {
		newTail = newTail.next
	}

	newhead := newTail.next
	newTail.next = nil

	tail.next = list.head
	list.head = newhead
}

func main() {
	list := &LinkedList{}
	for i := 5; i >= 0; i-- {
		list.append(i)
	}

	list.print()
	list.shift(2)
	list.print()
}
