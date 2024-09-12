package structures

import "errors"

type Stack struct {
	counter  int
	elements []int
}

// NewStack initializes a new stack with a given size capacity.
// The stack is backed by a slice of integers with pre-allocated space.
func NewStack(size int) *Stack {
	return &Stack{
		elements: make([]int, size),
	}
}

// Push adds a new value to the top of the stack.
// Time complexity: O(1)
func (s *Stack) Push(value int) ([]int, error) {
	if (s.counter == len(s.elements)) {
		return nil, errors.New("stack is full")
	}
	s.elements[s.counter] = value
	s.counter++
	return s.elements, nil
}

// Pop removes and returns the top element of the stack.
// Time complexity: O(1)
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	s.counter--
	return s.elements[s.counter], nil
}

// IsEmpty checks whether the stack is empty (i.e., has no elements).
// Time complexity: O(1)
func (s *Stack) IsEmpty() bool {
	return s.counter == 0
}

// Peak returns the top element of the stack without removing it.
// Time complexity: O(1)
func (s *Stack) Peak() int {
	return s.elements[s.counter-1]
}

// Print returns a slice containing the current elements in the stack,
// up to the current size indicated by counter.
// Time complexity: O(1)
func (s *Stack) Print() []int {
	return s.elements[:s.counter]
}
