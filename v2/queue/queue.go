// Package queue implements a queue.
//
// The internal representation is a slice of T
// that gets used as a circular buffer.
// This is instead of a more traditional approach
// that would use a linked list of nodes.
// The assumption is that contiguous slabs of RAM
// will generally provide more performance over pointers
// to nodes potentially scattered about the heap.
//
// There is a downside: whereas enqueueing to a
// linked list is always O(1), enqueueing here will
// be O(1) except for when the internal slice of ints
// has to be resized; then, enqueueing will be O(n)
// where n is the size of the queue before being resized.
//
// Therefore, when asking for a new instance of the
// queue, use NewWithCapacity() to pick a capacity that you
// think won't need to grow.
//
// When the queue does need to grow, it always uses a capacity
// that is twice the current capacity. Enqueue() will do this
// doubling for you automatically.
//
// However, if you would like to grow the backing slice
// yourself, to have control over 1) when the size is increased,
// and 2) how much larger the backing slice grows, you can use
// Resize() directly. If your code needs to ask the current
// capacity and length of the queue, Capacity() and Length()
// will provide those numbers.
package queue

import (
	"constraints"
	"errors"
	"fmt"
)

// DefaultCapacity is the default capacity of the queue
// when constructed using New() instead of NewWithCapacity().
const DefaultCapacity = 32

var QueueCapacityExceeded = errors.New("Queue Capacity Exceeded")
var QueueEmpty = errors.New("Queue Empty")

// Queue holds the data and state of the queue.
type Queue[T constraints.Ordered] struct {
	data     []T
	head     int
	tail     int
	capacity int
	length   int
}

// New returns a new empty queue for ints of the default capacity.
func New[T constraints.Ordered]() (q *Queue[T]) {
	return NewWithCapacity[T](DefaultCapacity)
}

// NewWithCapacity returns a new empty queue for ints with the requested capacity.
func NewWithCapacity[T constraints.Ordered](capacity int) (q *Queue[T]) {
	return &Queue[T]{
		data: make([]T, capacity, capacity),
		head: -1,
		tail: -1,
		capacity: capacity,
		length: 0,
	}
}

// Enqueue enqueues an int. Returns an error if the size
// of the queue cannot be grown any more to accommodate
// the added int.
func (q *Queue[T]) Enqueue(i T) error {
	if q.length+1 > q.capacity {
		newCapacity := q.capacity * 2
		// if newCapacity became negative, we have exceeded
		// our capacity by doing one bit-shift too far
		if newCapacity < 0 {
			return QueueCapacityExceeded
		}
		// NOTE: Purposefully not concerning ourselves
		// with the error returned from Resize here, because
		// we know our newCapacity is larger than q.capacity.
		q.Resize(newCapacity)
	}
	q.length++
	q.head++
	if q.head == q.capacity {
		q.head = 0
	}
	q.data[q.head] = i
	return nil
}

// EnqueueSlice enqueues a slice of ints. Returns an error
// if the size of the queue cannot be grown any more to accommodate
// the added ints.
func (q *Queue[T]) EnqueueSlice(sl []T) error {
	for _, i := range sl {
		err := q.Enqueue(i)
		if err != nil {
			return err
		}
	}
	return nil
}

// Length returns the current length
// of the queue. This is the same as the number of
// slots used in the slice that
// backs the queue.
func (q *Queue[T]) Length() int {
	return q.length
}

// Empty returns true if the queue is empty,
// false otherwise.
func (q *Queue[T]) Empty() bool {
	return q.length == 0
}

// Capacity returns the current capacity
// of the slice that backs the queue.
func (q *Queue[T]) Capacity() int {
	return q.capacity
}

// Resize resizes the underlying slice that backs
// the queue. The Enqueue method calls this automatically
// when the backing slice is full, but feel free to use
// this method preemptively if your calling code has a
// good time to do this resizing. Also, the Enqueue method
// uses a new backing slice that is twice the size of the
// old one; but if you call Resize yourself, you can pick
// whatever new size you want.
func (q *Queue[T]) Resize(newCapacity int) error {
	if newCapacity <= q.capacity {
		return fmt.Errorf("New capacity %d is not larger than current capacity %d", newCapacity, q.capacity)
	}
	newData := make([]T, newCapacity, newCapacity)
	var err error
	var i T
	// Because we are using the slice as a ring buffer,
	// head can be earlier in array than tail, so
	// it would be strange to just copy the old (possibly
	// partially wrapped) slice into the new slice.
	// Instead, we may as well copy the queue in order
	// into the new slice. The Dequeue() method gives us
	// every element in the correct order already, so we
	// just leverage that.
	for err = nil; err == nil; i, err = q.Dequeue() {
		newData = append(newData, i)
	}
	q.head = q.length - 1
	q.tail = 0
	q.capacity = newCapacity
	q.data = newData
	return nil
}

// Dequeue dequeues an int. It returns the dequeued int
// or an error if the queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	if q.length-1 < 0 {
		var zero T
		return zero, QueueEmpty
	}
	q.length--
	q.tail++
	if q.tail == q.capacity {
		q.tail = 0
	}
	return q.data[q.tail], nil
}
