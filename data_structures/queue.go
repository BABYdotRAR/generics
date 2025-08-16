package ds

import "sync"

// NewQueue retrieves a new generic queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

// Queue (thread-safe) is a generic FIFO queue
type Queue[T any] struct {
	lock  sync.Mutex
	size  int
	front *node[T]
	rear  *node[T]
}

// Enqueue adds val to the end of the queue
func (q *Queue[T]) Enqueue(val T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	newNode := &node[T]{
		value: val,
		next:  nil,
	}

	if q.isEmpty() {
		q.front, q.rear = newNode, newNode
		q.size = 1
		return
	}

	q.rear.next = newNode
	q.rear = newNode
	q.size++
}

// Dequeue retrieves the first element in the queue and removes it. 
// It returns the zero value for T and false if the queue is empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.isEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	res := q.front
	q.front = res.next
	if q.isEmpty() {
		q.rear = nil
	}

	q.size--
	res.next = nil

	return res.value, true
}

// Peek only retrieves the first element in the queue without removing it. 
// It returns the zero value for T and false if the queue is empty.
func (q *Queue[T])Peek() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.isEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	return q.front.value, true
}

// Size returns the total number of elements in the queue
func (q *Queue[T])Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.size
}

// Clear resets the queue
func (q *Queue[T])Clear() {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.front, q.rear = nil, nil
	q.size = 0
}

// IsEmpty returns whether the queue contains no elements
func (q *Queue[T])IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.isEmpty()
}

func (q *Queue[T])isEmpty() bool {
	return q.front == nil
}