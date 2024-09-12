package structures

import (
	"errors"
)

type Queue interface {
	Enqueue(int) (int, error)
	Dequeue() (int, error)
}

// QueueWithArray is a circular queue implementation backed by an array.
type QueueWithArray struct {
	front          int
	rear           int
	elements       []int
	filledCapacity int
}

// Enqueue adds an element to the rear of the queue.
// Time Complexity: O(1) - Inserting an element in a circular queue is done in constant time.
func (q *QueueWithArray) Enqueue(value int) (int, error) {
	if q.filledCapacity == len(q.elements) {
		return 0, errors.New("queue is full")
	}

	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)
	q.filledCapacity++

	return value, nil
}

// Dequeue removes an element from the front of the queue.
// Time Complexity: O(1) - Removing an element from a circular queue is done in constant time.
func (q *QueueWithArray) Dequeue() (int, error) {
	if q.filledCapacity == 0 {
		return 0, errors.New("queue is empty")
	}

	removedElement := q.elements[q.front] 
	q.elements[q.front] = 0
	q.front = (q.front + 1) % len(q.elements)
	q.filledCapacity--

	return removedElement, nil
}

// QueueWithStack is a queue implementation using two stacks.
type QueueWithStack struct {
	enqueueStack *Stack
	dequeueStack *Stack
}

// Enqueue adds an element to the queue using the enqueueStack.
// Time Complexity: O(1) - Pushing onto a stack is done in constant time.
func (q *QueueWithStack) Enqueue(value int) (int, error) {
	_, err := q.enqueueStack.Push(value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// Dequeue removes an element from the front of the queue, transferring elements from enqueueStack
// to dequeueStack if needed.
// Time Complexity: Amortized O(1), Worst-case O(n) - Each element is moved between the stacks
// at most once, so the amortized complexity is O(1).
func (q *QueueWithStack) Dequeue() (int, error) {
	if q.dequeueStack.IsEmpty() {
		for !q.enqueueStack.IsEmpty() {
			element, err := q.enqueueStack.Pop()
			if err != nil {
				
				return 0, err
			}
			q.dequeueStack.Push(element)
		}
	}

	if q.dequeueStack.IsEmpty() {
		return 0, errors.New("queue is empty")
	}

	element, err := q.dequeueStack.Pop()
	if err != nil {
		return 0, err
	} 

	return element, nil
}

// NewQueue creates a new queue of the given type ("array" or "stack") and size.
func NewQueue(queuetype string, size int) (Queue, error) {
	if size <= 0 {
		return nil, errors.New("invalid queue size")
	}

	switch queuetype{
	case "array":
		return &QueueWithArray{
			elements: make([]int, size),
		}, nil
	case "stack":
		return &QueueWithStack{
			enqueueStack: NewStack(size),
			dequeueStack: NewStack(size),
		}, nil
	default:
		return nil , errors.New("invalid queue type")
	}
}
