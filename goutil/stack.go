package goutil

import (
	"sync"
)

type Stack[T any] struct {
	keys []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (s *Stack[T]) Push(key T) {
	s.keys = append(s.keys, key)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.keys) == 0
}

// Pop returns T and true if a value exists, false if the stack is empty
func (s *Stack[T]) Pop() (T, bool) {
	var x T
	if s.IsEmpty() {
		return x, false
	}
	x, s.keys = s.keys[len(s.keys)-1], s.keys[:len(s.keys)-1]
	return x, true
}

type SafeStack[T any] struct {
	mu    sync.Mutex
	items []T
	cond  *sync.Cond
}

func NewSafeStack[T any]() *SafeStack[T] {
	s := &SafeStack[T]{items: make([]T, 100)}
	s.cond = sync.NewCond(&s.mu)
	return s
}

func (s *SafeStack[T]) Push(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, item)
	s.cond.Signal()
}

func (s *SafeStack[T]) Pop() T {
	s.mu.Lock()
	defer s.mu.Unlock()

	for len(s.items) == 0 {
		s.cond.Wait()
	}

	var x T
	x, s.items = s.items[len(s.items)-1], s.items[:len(s.items)-1]
	return x
}
