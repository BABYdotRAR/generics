package ds

import "sync"

// NewStack retrieves a new generic stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Stack (thread-safe) is a generic LIFO stack
type Stack[T any] struct {
	lock sync.Mutex
	size int
	top  *node[T]
}

type node[T any] struct {
	value T
	next  *node[T]
}

// IsEmpty returns whether the stack contains no elements
func (s *Stack[T]) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.isEmpty()
}

func (s *Stack[T]) isEmpty() bool {
	return s.top == nil
}

// Push adds elem at the top of the stack
func (s *Stack[T]) Push(elem T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.top = &node[T]{
		value: elem,
		next:  s.top,
	}
	s.size++
}

// Pop removes and returns the top element of the stack.
// It returns the zero value of T and false if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	elem := s.top
	s.size--
	s.top = elem.next
	return elem.value, true
}

// Top returns the top element of the stack without removing it.
// It returns the zero value of T and false if the stack is empty
func (s *Stack[T]) Top() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return s.top.value, true
}

// Size returns the current number of elements in the stack
// O(1) cost
func (s *Stack[T]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.size
}

// Clear removes all elements from the stack
func (s *Stack[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.top = nil
	s.size = 0
}
