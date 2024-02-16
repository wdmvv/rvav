package calc

type Stack[T any] struct {
	elems []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (s *Stack[T]) Push(key T) {
	s.elems = append(s.elems, key)
}

func (s *Stack[T]) Top() T {
	return s.elems[len(s.elems)-1]
}

func (s *Stack[T]) Pop() T {
	var x T
	x, s.elems = s.elems[len(s.elems)-1], s.elems[:len(s.elems)-1]
	return x
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Stack[T]) Len() int{
	return len(s.elems)
}