// Package bag implements a bag.
// It's just syntactic sugar around map[T]int
// to track how many copies of something are in the bag.
package bag

type Bag[T comparable] map[T]int

func New[T comparable]() Bag[T] {
	return Bag[T]{}
}

// Has can tell you if the bag contains at least one of elem.
func (b Bag[T]) Has(elem T) bool {
	_, ok := b[elem]
	return ok
}

// Put puts an element in the bag.
func (b Bag[T]) Put(elem T) {
	b[elem] = b[elem] + 1
}

func (b Bag[T]) Delete(elem T) {
	_, ok := b[elem]
	if !ok {
		// Element does not exist; do nothing.
		return
	}
	// Element exists; decrement.
	b[elem] = b[elem] - 1
	// If this was the last copy of this element, nuke it from the
	// backing map.
	if b[elem] == 0 {
		delete(b, elem)
	}
}

func (b Bag[T]) PutSlice(elements []T) {
	for _, elem := range elements {
		b.Put(elem)
	}
}

// Iter iterates through every element of the bag and calls
// function f using the element as an argument for T.
func (b Bag[T]) Iter(f func(elem T)) {
	for elem, count := range b {
		for i := count; i > 0; i-- {
			f(elem)
		}
	}
}
