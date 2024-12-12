package goutil

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
