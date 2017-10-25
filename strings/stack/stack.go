// Package stack implements a stack for strings.
//
// It is an express design decision to hard-code
// this stack just for the string type rather than for
// the empty interface.
//
package stack

import "github.com/pkg/errors"

// Stack holds the data and state of the stack.
type Stack struct {
	data     []string
	top      int
	capacity int
	size     int
}

// DefaultCapacity is the default capacity of the stack
// when constructed using New() instead of NewWithCapacity().
const DefaultCapacity = 32

// New returns a new empty stack for strings of the default capacity.
func New() (s *Stack) {
	return NewWithCapacity(DefaultCapacity)
}

// NewWithCapacity returns a new empty stack for strings with the requested capacity.
func NewWithCapacity(capacity int) (s *Stack) {
	s = new(Stack)
	s.data = make([]string, capacity, capacity)
	s.top = -1
	s.capacity = capacity
	s.size = 0
	return s
}

// Push pushes a string onto the stack. Returns an error if the size
// of the stack cannot be grown any more to accommodate
// the added string.
func (s *Stack) Push(i string) error {
	if s.size+1 > s.capacity {
		newCapacity := s.capacity * 2
		// if newCapacity became negative, we have exceeded
		// our capacity by doing one bit-shift too far
		if newCapacity < 0 {
			return errors.New("Capacity exceeded")
		}
		// NOTE: Purposefully not concerning ourselves
		// with the error returned from Resize here, because
		// we know our newCapacity is larger than q.capacity.
		s.Resize(newCapacity)
	}
	s.size++
	s.top++
	s.data[s.top] = i
	return nil
}

// Size returns the current size
// of the stack. It also tells you how many
// slots are being used in the slice that
// backs the stack.
func (s *Stack) Size() int {
	return s.size
}

// Capacity returns the current capacity
// of the slice that backs the queue.
func (s *Stack) Capacity() int {
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
func (s *Stack) Resize(newCapacity int) error {
	if newCapacity <= s.capacity {
		return errors.Errorf("New capacity %d is not larger than current capacity %d", newCapacity, s.capacity)
	}
	newData := make([]string, newCapacity, newCapacity)
	for i := 0; i < len(s.data); i++ {
		newData[i] = s.data[i]
	}
	s.capacity = newCapacity
	s.data = newData
	return nil
}

// Pop pops the string off the top of the stack. It returns the popped string
// or an error of the stack is empty.
func (s *Stack) Pop() (string, error) {
	if s.size-1 < 0 {
		return "", errors.New("Stack empty")
	}
	str := s.data[s.top]
	s.size--
	s.top--
	return str, nil
}

// Peek returns the string off the top of the stack
// but does not remove it.
// If the stack is empty, an error is returned.
func (s *Stack) Peek() (string, error) {
	if s.size-1 < 0 {
		return "", errors.New("Stack empty")
	}
	return s.data[s.top], nil
}
