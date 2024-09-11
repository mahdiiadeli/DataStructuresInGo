package linkedlist

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
// This operation involves updating the last node's next pointer and last pointer, or setting both pointers if the list is empty.
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
// This operation always involves setting the new node's next pointer to the current first node and updating the first pointer.
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
// This operation requires traversing the list to find the second-to-last node, then updating its next pointer to nil and the last pointer.
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
// This operation involves updating the first pointer to the next node, and if necessary, resetting the last pointer if the list becomes empty.
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
// This operation involves traversing the list from the beginning to find the node with the given value.
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
// This operation involves scanning through the linked list to find the node with the specified value.
func (ls *LinkedList) Contains(value int) bool {
	return ls.IndexOf(value) != -1
}

// Size returns the number of elements in the linked list
// Time Complexity: O(1)
// This operation retrieves the size of the linked list from a stored attribute.
func (ls *LinkedList) Size() int {
	return ls.size
}

// ToArray converts the linked list to a slice of integers
// Time Complexity: O(n)
// This operation involves traversing the entire linked list to copy each node's value into a new slice.
func (ls *LinkedList) ToArray() []int {
	arr := make([]int, ls.size, ls.size)
	currentNode := ls.first
	arrIndex := 0

	for currentNode != nil {
		arr[arrIndex] = currentNode.value
		currentNode = currentNode.next
		arrIndex++
	}

	return arr
}
