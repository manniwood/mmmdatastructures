// Package stringqueue implements a queue for strings.
//
// It is an express design decision to hard-code
// this queue just for the int type rather than for
// the empty interface.
//
// The internal representation is a slice of strings
// that gets used as a circular buffer.
// This is instead of a more traditional approach
// that would use a linked list of nodes.
// The assumption is that contiguous slabs of RAM
// will generally provide more performance over pointers
// to nodes potentially scattered about the heap.
//
// There is a downside: whereas enqueueing to a
// linked list is always O(1), enqueueing here will
// be O(1) except for when the internal slice of strings
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
package stringqueue

import "github.com/pkg/errors"

// Queue holds the data and state of the queue.
type Queue struct {
	data     []string
	head     int
	tail     int
	capacity int
	length   int
}

// DefaultCapacity is the default capacity of the queue
// when constructed using New() instead of NewWithCapacity().
const DefaultCapacity = 32

// New returns a new empty queue for strings of the default capacity.
func New() (q *Queue) {
	return NewWithCapacity(DefaultCapacity)
}

// NewWithCapacity returns a new empty queue for strings with the requested capacity.
func NewWithCapacity(capacity int) (q *Queue) {
	q = new(Queue)
	q.data = make([]string, capacity, capacity)
	q.head = -1
	q.tail = -1
	q.capacity = capacity
	q.length = 0
	return q
}

// Enqueue enqueues a string. Returns an error if the size
// of the queue cannot be grown any more to accommodate
// the added string.
func (q *Queue) Enqueue(i string) error {
	if q.length+1 > q.capacity {
		newCapacity := q.capacity * 2
		// if new_cap became negative, we have exceeded
		// our capacity by doing one bit-shift too far
		if newCapacity < 0 {
			return errors.New("Capacity exceeded")
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

// Length tells you the current length
// of the queue. It also tells you how many
// slots are being used in the slice that
// backs the queue.
func (q *Queue) Length() int {
	return q.length
}

// Capacity tells you the current capacity
// of the slice that backs the queue.
func (q *Queue) Capacity() int {
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
func (q *Queue) Resize(newCapacity int) error {
	if newCapacity <= q.capacity {
		return errors.Errorf("New capacity %d is not larger than current capacity %d", newCapacity, q.capacity)
	}
	new_data := make([]string, newCapacity, newCapacity)
	var err error
	var i string
	// Because we are using the slice as a ring buffer,
	// head can be earlier in array than tail, so
	// it would be strange to just copy the old (possibly
	// partially wrapped) slice into the new slice.
	// Instead, we may as well copy the queue in order
	// into the new slice. The Dequeue() method gives us
	// every element in the correct order already, so we
	// just leverage that.
	for err = nil; err == nil; i, err = q.Dequeue() {
		new_data = append(new_data, i)
	}
	q.head = q.length - 1
	q.tail = 0
	q.capacity = newCapacity
	q.data = new_data
	return nil
}

// Dequeue dequeues a string. It returns the dequeued string
// or an error of the queue is empty.
func (q *Queue) Dequeue() (string, error) {
	if q.length-1 < 0 {
		return "", errors.New("Queue empty")
	}
	q.length--
	q.tail++
	if q.tail == q.capacity {
		q.tail = 0
	}
	return q.data[q.tail], nil
}
