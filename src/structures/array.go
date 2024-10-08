// Package structures provides a simple dynamic array data structure implementation.
// It supports inserting elements with automatic resizing, removing elements by index,
// searching for elements, and printing the current array contents.
// The array dynamically grows in size when needed, but does not shrink automatically.

package structures

type Array struct {
	initialCapacity int
	filledCapacity  int
	items           []int
}

// NewArray initializes a new array with a given initial capacity.
func NewArray(initialCapacity int) *Array {
	return &Array{
		initialCapacity: initialCapacity,
		filledCapacity:  0,
		items:           make([]int, initialCapacity),
	}
}

// Insert adds a new item to the array, resizing if necessary.
// Time complexity (amortized): O(1)
func (arr *Array) Insert(item int) {
	if arr.filledCapacity == len(arr.items) {
		newItems := make([]int, len(arr.items)*2)
		copy(newItems, arr.items)
		arr.items = newItems
	}

	arr.items[arr.filledCapacity] = item
	arr.filledCapacity += 1
}

// RemoveAt removes an element from the array at a specific index.
// Time complexity: O(n) (due to element shifting)
func (arr *Array) RemoveAt(index int) error {
	if index < 0 || index >= len(arr.items) {
		panic("index out of range")
	}

	for i := index; i < len(arr.items)-1; i++ {
		arr.items[i] = arr.items[i+1]
	}
	arr.items = arr.items[:arr.filledCapacity-1]

	arr.filledCapacity--

	return nil
}

// IndexOf returns the index of a given item, or -1 if not found.
// Time complexity: O(n)
func (arr *Array) IndexOf(item int) int {
	for i, v := range arr.items {
		if v == item {
			return i
		}
	}

	return -1
}

// Print returns the array elements up to filledCapacity.
// Time complexity: O(1)
func (arr *Array) Print() []int {
	return arr.items
}
