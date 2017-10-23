// Package intqueue implements a queue for ints.
//
// The internal representation is a slice of ints
// that gets used as a circular buffer.
// This is instead of a more traditional approach
// that would use a linked list of nodes.
// The assumption is that contiguous slabs of RAM
// will generally provide more performance over pointers
// to nodes around the heap.
//
// There is a downside: whereas enqueueing to a
// linked list is always O(1), enqueueing here will
// be O(1) except for when the internal slice of ints
// has to be resized; then, enqueueing will be O(n)
// where n is the size of the queue before being resized.
//
// Therefore, when asking for a new instance of the
// queue, pick a capacity that you think won't need to grow.
package intqueue

import "github.com/pkg/errors"

// IntQueue holds the data and state of the queue.
type IntQueue struct {
	data     []int
	head     int
	tail     int
	capacity int
	length   int
}

// Creates a new empty queue for ints and returns a pointer to it.
func New() (q *IntQueue) {
	q = new(IntQueue)
	q.data = make([]int, 32, 32)
	q.head = -1
	q.tail = -1
	q.capacity = 32
	q.length = 0
	return q
}

// Enqueues an int. Returns an error if the size
// of the queue cannot be grown any more to accommodate
// the added int.
func (q *IntQueue) Enqueue(i int) error {
	if q.length+1 > q.capacity {
		new_capacity := q.capacity << 1
		// if new_cap became negative, we have exceeded
		// our capacity by doing one bit-shift too far
		if new_capacity < 0 {
			return errors.New("Capacity exceeded")
		}
		q.resize(new_capacity)
	}
	q.length++
	q.head++
	if q.head == q.capacity {
		q.head = 0
	}
	q.data[q.head] = i
	return nil
}

// Head can be earlier in array than tail, so
// we can't just copy; we could overwrite the tail.
// Instead, we may as well copy the queue in order
// into the new array. The Dequeue() method gives us
// every element in the correct order already, so we
// just leverage that.
func (q *IntQueue) resize(new_capacity int) {
	new_data := make([]int, new_capacity, new_capacity)
	var err error
	var i int
	for err = nil; err == nil; i, err = q.Dequeue() {
		new_data = append(new_data, i)
	}
	q.head = q.length - 1
	q.tail = 0
	q.capacity = new_capacity
	q.data = new_data
}

// Enqueues an int. Returns an error of the queue is empty.
func (q *IntQueue) Dequeue() (int, error) {
	if q.length-1 < 0 {
		return 0, errors.New("Queue empty")
	}
	q.length--
	q.tail++
	if q.tail == q.capacity {
		q.tail = 0
	}
	return q.data[q.tail], nil
}
