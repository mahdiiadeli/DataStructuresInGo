// Package linkedlist provides a simple implementation of a singly linked list data structure.
// It supports operations to manage and manipulate the list, including adding and removing elements,
// searching for values, and converting the list to an array. The linked list dynamically grows
// or shrinks as needed, with efficient insertions and deletions at the beginning and end.

package structures

import "fmt"

type LinkedList struct {
	first *Node
	last  *Node
	size int
}

type Node struct {
	value int
	next  *Node
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// newNode creates a new node with the given value
func newNode(value int) *Node {
	return &Node{value: value}
}

// isEmpty checks if the linked list is empty
// Time Complexity: O(1)
func (ls *LinkedList) isEmpty() bool {
	return ls.first == nil
}

// AddLast adds a node with the given value at the end of the list
// Time Complexity: O(1)
func (ls *LinkedList) AddLast(value int) {
	newNode := newNode(value)
	if ls.isEmpty() {
		ls.first = newNode
		ls.last = newNode
	} else {
		ls.last.next = newNode
		ls.last = newNode
	}

	ls.size++
}

// AddFirst adds a node with the given value at the beginning of the list
// Time Complexity: O(1)
func (ls *LinkedList) AddFirst(value int) {
	newNode := newNode(value)
	if ls.isEmpty() {
		ls.first = newNode
		ls.last = newNode
	} else {
		newNode.next = ls.first
		ls.first = newNode
	}
	
	ls.size++
}

// RemoveLast removes the last node from the list
// Time Complexity: O(n)
func (ls *LinkedList) RemoveLast() {
	if ls.isEmpty() || ls.first == ls.last {
		ls.first = nil
		ls.last = nil
		ls.size = 0
		return
	}

	currentNode := ls.first
	for currentNode.next != ls.last {
		currentNode = currentNode.next
	}

	currentNode.next = nil
	ls.last = currentNode
	ls.size--
}

// RemoveFirst removes the first node from the list
// Time Complexity: O(1)
func (ls *LinkedList) RemoveFirst() {
	if ls.isEmpty() || ls.first == ls.last {
		ls.first = nil
		ls.last = nil
		ls.size = 0
		return
	}

	newFirstNode := ls.first.next
	ls.first.next = nil
	ls.first = newFirstNode
	ls.size--
}

// IndexOf returns the index of the first occurrence of the given value, or -1 if not found
// Time Complexity: O(n)
func (ls *LinkedList) IndexOf(value int) int {
	index := 0
	currentNode := ls.first
	for currentNode != nil {
		if currentNode.value == value {
			return index
		}
		currentNode = currentNode.next
		index++
	}

	return -1
}

// Contains checks if a value is present in the linked list
// Time Complexity: O(n)
func (ls *LinkedList) Contains(value int) bool {
	return ls.IndexOf(value) != -1
}

// Size returns the number of elements in the linked list
// Time Complexity: O(1)
func (ls *LinkedList) Size() int {
	return ls.size
}

// ToArray converts the linked list to a slice of integers
// Time Complexity: O(n)
func (ls *LinkedList) ToArray() []int {
	arr := make([]int, ls.size)
	currentNode := ls.first
	arrIndex := 0

	for currentNode != nil {
		arr[arrIndex] = currentNode.value
		currentNode = currentNode.next
		arrIndex++
	}

	return arr
}

// Reverse reverses the order of nodes in the linked list.
// Time Complexity: O(n) 
func (ls *LinkedList) Reverse() {
	if ls.isEmpty() || ls.first.next == nil {
		return
	}

	currentNode := ls.first
	var previousNode *Node

	for currentNode != nil {
		nextNode := currentNode.next
		currentNode.next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	ls.last = ls.first
	ls.first = previousNode
	ls.last.next= nil
}

// GetKthFromTheEnd returns the value of the k-th node from the end of the linked list.
// Time Complexity: O(n)
func (ls *LinkedList) GetKthFromTheEnd(k int) int {
	if k <= 0 {
		panic("K must be greater than 0")
	}
	
	lead := ls.first
	follow := ls.first

	for i := 0; i < k -1 ; i++ {
		fmt.Println(follow)
		fmt.Println(follow.next)
		follow = follow.next
		if follow == nil {
			panic("K out of range")
		}
	}

	for follow != ls.last {
		lead = lead.next
		follow = follow.next
	}

	return lead.value
}