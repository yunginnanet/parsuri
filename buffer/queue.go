package buffer

import (
	"github.com/gammazero/deque"
	"reflect"
	"sync"
)

type Queue[T any] struct {
	dq       *deque.Deque[T]
	mu       *sync.RWMutex
	capacity int
	zero     T
}

func NewQueue[T any](initialCap int) *Queue[T] {
	dq := &deque.Deque[T]{}
	dq.SetBaseCap(initialCap)
	tv := new(T)
	return &Queue[T]{
		dq:       dq,
		mu:       &sync.RWMutex{},
		zero:     *tv,
		capacity: initialCap,
	}
}

func (q *Queue[T]) isZero(val T) bool {
	return reflect.DeepEqual(val, q.zero)
}

func (q *Queue[T]) Push(item T) {
	q.mu.Lock()
	q.dq.PushFront(item)
	q.mu.Unlock()
}

// Pop returns the last item in the queue and removes it.
//
// If the item is "zero", it will try to pop again until
// a non-zero item is found or the queue is empty.
//
// If the queue is empty, it returns the zero value of T and false.
func (q *Queue[T]) Pop() (T, bool) {
	q.mu.RLock()
	if q.dq.Len() == 0 {
		var z T
		q.mu.RUnlock()
		return z, false
	}
	val := q.dq.PopBack()
	q.mu.RUnlock()
	if q.isZero(val) {
		return q.Pop()
	}
	return val, true
}

func (q *Queue[T]) PopFront() (T, bool) {
	q.mu.RLock()
	if q.dq.Len() == 0 {
		var z T
		q.mu.RUnlock()
		return z, false
	}
	val := q.dq.PopFront()
	q.mu.RUnlock()
	if q.isZero(val) {
		return q.PopFront()
	}
	return val, true
}

func (q *Queue[T]) Clear() {
	q.mu.Lock()
	q.dq.Clear()
	q.mu.Unlock()
}

func (q *Queue[T]) Len() int {
	q.mu.RLock()
	l := q.dq.Len()
	q.mu.RUnlock()
	return l
}
