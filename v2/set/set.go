// Package set implements a set.
// It's just syntactic sugar around map[T]struct{}
package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return Set[T]{}
}

func (s Set[T]) Has(elem T) bool {
	_, ok := s[elem]
	return ok
}

func (s Set[T]) Put(elem T) {
	s[elem] = struct{}{}
}

func (s Set[T]) Delete(elem T) {
	delete(s, elem)
}

func (s Set[T]) PutSlice(elements []T) {
	for _, elem := range elements {
		s[elem] = struct{}{}
	}
}
