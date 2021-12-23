// Package set implements a set for ints.
// It's just syntactic sugar around map[int]struct{}
package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	m := make(map[T]struct{})
	return Set[T](m)
	// return Set[T]{}
}

func (s Set[T]) Has(k T) bool {
	_, ok := s[k]
	return ok
}

func (s Set[T]) Put(k T) {
	s[k] = struct{}{}
}

func (s Set[T]) Delete(k T) {
	delete(s, k)
}

func (s Set[T]) PutSlice(ks []T) {
	for _, k := range ks {
		s[k] = struct{}{}
	}
}
