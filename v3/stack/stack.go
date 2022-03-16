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

type NegativeStackCapacityError struct {
	msg string
}

func (e *NegativeStackCapacityError) Error() string {
	return e.msg
}

type ResizeStackCapacityError struct {
	msg string
}

func (e *ResizeStackCapacityError) Error() string {
	return e.msg
}

var StackCapacityExceeded = errors.New("Stack Capacity Exceeded")
var StackEmpty = errors.New("Stack Empty")

// Stack holds the data and state of the stack.
type Stack[T constraints.Ordered] struct {
	data []T
	// top is the topmost index of data[] that holds an element.
	top      int
	capacity int
}

// New returns a new empty stack of the default capacity.
func New[T constraints.Ordered]() (*Stack[T], error) {
	return NewWithCapacity[T](DefaultCapacity)
}

// NewWithCapacity returns a new empty stack with the requested capacity.
func NewWithCapacity[T constraints.Ordered](capacity int) (*Stack[T], error) {
	if capacity < 1 {
		return nil, &NegativeStackCapacityError{
			msg: fmt.Sprintf("capacity %d is zero or negative", capacity),
		}
	}
	return &Stack[T]{
		data: make([]T, capacity, capacity),
		// When the stack is empty, top == -1, whereas when the stack contains
		// one element, top == 0, the "0th" element of data[].
		top:      -1,
		capacity: capacity,
	}, nil
}

// Push pushes an element onto the stack. It returns an error if the size
// of the stack cannot be grown any more to accommodate
// the added element.
func (s *Stack[T]) Push(elem T) error {
	// "s.top+2" seems weird at first, but look at it this way:
	// if we have a s.data[] of capacity 1, then it has one slot with index 0.
	// And s.top begins pointing at index -1 so that we can increment
	// s.top whenever we push; so the first-pushed element would increment
	// s.top to 0. So before we push an element, if we take current top
	// (which is -1 in this example) and want to know if it will exceed
	// capacity, we have to add 2, because capacity is always max index + 1.
	if s.top+2 > s.capacity {
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
	s.top++
	s.data[s.top] = elem
	return nil
}

// Size returns the current size of the stack. It also tells you how many
// slots are being used in the slice that backs the stack.
func (s *Stack[T]) Size() int {
	// If we have a slice of size 1, it has only one index, 0.
	// So if our slice has that "0th" slot full, s.top points to index 0.
	// So adding 1 to s.top gives us the size of the stack.
	return s.top + 1
}

// Len is a synonym for Size(), to mimic the len() built-in used for slices.
func (s *Stack[T]) Len() int {
	return s.Size()
}

// Cap returns the current capacity of the slice that backs the stack.
// Cap is named Cap to mimic the built-in cap() command used for slices.
func (s *Stack[T]) Cap() int {
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
		return &ResizeStackCapacityError{
			msg: fmt.Sprintf("New capacity %d is not larger than current capacity %d", newCapacity, s.capacity),
		}
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
	if s.top == -1 {
		var zero T
		return zero, StackEmpty
	}
	elem := s.data[s.top]
	s.top--
	return elem, nil
}

// Peek returns stack's top element but does not remove it.
// If the stack is empty, an error is returned.
func (s *Stack[T]) Peek() (T, error) {
	if s.top == -1 {
		var zero T
		return zero, StackEmpty
	}
	return s.data[s.top], nil
}
