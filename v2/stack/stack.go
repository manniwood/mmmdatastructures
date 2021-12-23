// Package stack implements a stack.
package stack

import (
	"constraints"
	"errors"
	"fmt"
)

// DefaultCapacity is the default capacity of the stack
// when constructed using New() instead of NewWithCapacity().
const DefaultCapacity = 32

var StackCapacityExceeded = errors.New("Stack Capacity Exceeded")
var StackEmpty = errors.New("Stack Empty")

// Stack holds the data and state of the stack.
type Stack[T constraints.Ordered] struct {
	data     []T
	top      int
	capacity int
	size     int
}

// New returns a new empty stack of the default capacity.
func New[T constraints.Ordered]() (s *Stack[T]) {
	return NewWithCapacity[T](DefaultCapacity)
}

// NewWithCapacity returns a new empty stack with the requested capacity.
func NewWithCapacity[T constraints.Ordered](capacity int) (s *Stack[T]) {
	return &Stack[T]{
		data: make([]T, capacity, capacity),
		top: -1,
		capacity: capacity,
		size: 0,
	}
}

// Push pushes an element onto the stack. It returns an error if the size
// of the stack cannot be grown any more to accommodate
// the added element.
func (s *Stack[T]) Push(elem T) error {
	if s.size+1 > s.capacity {
		newCapacity := s.capacity * 2
		// If newCapacity became negative, we have exceeded
		// our capacity.
		if newCapacity < 0 {
			return StackCapacityExceeded
		}
		// NOTE: We are purposefully not concerning ourselves
		// with the error returned from Resize here, because
		// we know our newCapacity is larger than q.capacity.
		s.Resize(newCapacity)
	}
	s.size++
	s.top++
	s.data[s.top] = elem
	return nil
}

// Size returns the current size
// of the stack. It also tells you how many
// slots are being used in the slice that
// backs the stack.
func (s *Stack[T]) Size() int {
	return s.size
}

// Capacity returns the current capacity
// of the slice that backs the stack.
func (s *Stack[T]) Capacity() int {
	return s.capacity
}

// Resize resizes the underlying slice that backs
// the stack. The Push method calls this automatically
// when the backing slice is full, but feel free to use
// this method preemptively if your calling code has a
// good time to do this resizing. Also, the Push method
// uses a new backing slice that is twice the size of the
// old one; but if you call Resize yourself, you can pick
// whatever new size you want.
func (s *Stack[T]) Resize(newCapacity int) error {
	if newCapacity <= s.capacity {
		return fmt.Errorf("New capacity %d is not larger than current capacity %d", newCapacity, s.capacity)
	}
	newData := make([]T, newCapacity, newCapacity)
	for i := 0; i < len(s.data); i++ {
		newData[i] = s.data[i]
	}
	s.capacity = newCapacity
	s.data = newData
	return nil
}

// Pop pops the top element off the stack. It returns the popped element
// or an error of the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	if s.size-1 < 0 {
		var zero T
		return zero, StackEmpty
	}
	elem := s.data[s.top]
	s.size--
	s.top--
	return elem, nil
}

// Peek returns stack's top element but does not remove it.
// If the stack is empty, an error is returned.
func (s *Stack[T]) Peek() (T, error) {
	if s.size-1 < 0 {
		var zero T
		return zero, StackEmpty
	}
	return s.data[s.top], nil
}
