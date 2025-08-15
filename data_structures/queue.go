package ds

import "sync"

type Queue[T any] struct {
	lock  sync.Mutex
	size  int
	front *node[T]
	rear  *node[T]
}

func (q *Queue[T]) Enqueue(val T) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.size++
	node := node[T]{
		value: val,
		next:  nil,
	}

	if q.front == nil {
		q.front = &node
		q.rear = &node
		return
	}

	q.rear.next = &node
	q.rear = &node
}

func (q *Queue[T])Dequeue() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.front == nil {
		var zeroValue T
		return zeroValue, false
	}

	q.size--
	if q.front == q.rear {
		q.rear = nil
	}
	res := q.front
	q.front = res.next
	return res.value, true
}