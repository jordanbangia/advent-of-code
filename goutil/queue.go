package goutil

type Queue[T any] struct {
	keys []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{nil}
}

func (q *Queue[T]) Push(i T) {
	q.keys = append(q.keys, i)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.keys) == 0
}

func (q *Queue[T]) Pop() (T, bool) {
	var x T
	if q.IsEmpty() {
		return x, false
	}
	x, q.keys = q.keys[0], q.keys[1:]
	return x, true
}
