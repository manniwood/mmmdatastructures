package queue

import "github.com/pkg/errors"

// IntQueue holds the data and state of the queue.
type IntQueue struct {
	Data     []int
	Head     int
	Tail     int
	Capacity int
	Length   int
}

// Creates a new empty queue for ints and returns a pointer to it.
func NewInt() (q *IntQueue) {
	q = new(IntQueue)
	q.Data = make([]int, 32, 32)
	q.Head = -1
	q.Tail = -1
	q.Capacity = 32
	q.Length = 0
	return q
}

// Enqueues an int. Returns an error if the size
// of the queue cannot be grown any more to accommodate
// the added int.
func (q *IntQueue) Enqueue(i int) error {
	if q.Length+1 > q.Capacity {
		new_cap := q.Capacity << 1
		// if new_cap became negative, we have exceeded
		// our capacity by doing one bit-shift too far
		if new_cap < 0 {
			return errors.New("Capacity exceeded")
		}
		q.resize(new_cap)
	}
	q.Length++
	q.Head++
	if q.Head == q.Capacity {
		q.Head = 0
	}
	q.Data[q.Head] = i
	return nil
}

// Head can be earlier in array than tail, so
// we can't just copy; we could overwrite the tail.
// Instead, we may as well copy the queue in order
// into the new array.
func (q *IntQueue) resize(new_cap int) {
	new_data := make([]int, new_cap, new_cap)
	var err error
	var i int
	for err = nil; err == nil; i, err = q.Dequeue() {
		new_data = append(new_data, i)
	}
	q.Head = q.Length - 1
	q.Tail = 0
	q.Capacity = new_cap
	q.Data = new_data
}

// Enqueues an int. Returns an error of the queue is empty.
func (q *IntQueue) Dequeue() (int, error) {
	if q.Length-1 < 0 {
		return 0, errors.New("Queue empty")
	}
	q.Length--
	q.Tail++
	if q.Tail == q.Capacity {
		q.Tail = 0
	}
	return q.Data[q.Tail], nil
}
