// Package maxheap implements a binary max heap for strings.
//
// It is an express design decision to hard-code
// this max heap just for the string type rather than for
// the empty interface.
//
package maxheap

import (
	"github.com/manniwood/mmmdatastructures"
	"github.com/pkg/errors"
)

// MaxHeap holds the data and state of the max heap.
type MaxHeap struct {
	data     []string
	capacity int
	size     int
}

// DefaultCapacity is the default capacity of the max heap
// when constructed using New() instead of NewWithCapacity().
const DefaultCapacity = 32

// New returns a new empty max heap for strings of the default capacity.
func New() (h *MaxHeap) {
	return NewWithCapacity(DefaultCapacity)
}

// NewWithCapacity returns a new empty max heap for strings with the requested capacity
// rounded up to the next power of two.
func NewWithCapacity(requested int) (h *MaxHeap) {
	power := 1
	for power < requested {
		power *= 2
		if power < 0 {
			// looks like we wrapped
			power = mmmdatastructures.MaxInt
			break
		}
	}
	h = new(MaxHeap)
	h.data = make([]string, power, power)
	h.capacity = power
	h.size = 0
	return h
}

// Insert inserts a string onto the max heap. Returns an error if the size
// of the max heap cannot be grown any more to accommodate
// the added string.
func (h *MaxHeap) Insert(str string) error {
	if h.size+1 > h.capacity {
		newCapacity := h.capacity * 2
		// if newCapacity became negative, we have exceeded
		// our capacity by doing one bit-shift too far
		if newCapacity < 0 {
			return errors.New("Capacity exceeded")
		}
		// NOTE: Purposefully not concerning ourselves
		// with the error returned from Resize here, because
		// we know our newCapacity is larger than q.capacity.
		h.resize(newCapacity)
	}
	// Increase the size of the max heap. Usefully, the size
	// is also the new last index into the backing slice.
	// Put our new value there. Then, bubble the new value
	// up, swapping it with its parent, until it is in the
	// correct position in the max heap.
	h.size++
	h.data[h.size] = str
	child := h.size
	for parent := child / 2; parent > 0; parent = child / 2 {
		if h.data[child] > h.data[parent] {
			h.data[child], h.data[parent] = h.data[parent], h.data[child]
		}
		child = parent
	}
	return nil
}

// Size returns the current size
// of the max heap. It also tells you how many
// slots are being used in the slice that
// backs the max heap.
func (h *MaxHeap) Size() int {
	return h.size
}

// Capacity returns the current capacity
// of the slice that backs the queue.
func (h *MaxHeap) Capacity() int {
	return h.capacity
}

// resize resizes the underlying slice that backs
// the max heap. It is made private, because we
// want to enforce resize only being called with
// a capacity that is twice the size of the previous
// capacity.
func (h *MaxHeap) resize(newCapacity int) error {
	if newCapacity <= h.capacity {
		return errors.Errorf("New capacity %d is not larger than current capacity %d", newCapacity, h.capacity)
	}
	newData := make([]string, newCapacity, newCapacity)
	for i := 0; i < len(h.data); i++ {
		newData[i] = h.data[i]
	}
	h.capacity = newCapacity
	h.data = newData
	return nil
}

// Peek returns the largest value from the top of the
// heap, without removing it.
func (h *MaxHeap) Peek() (string, error) {
	if h.size == 0 {
		return "", errors.New("Heap empty")
	}
	return h.data[1], nil
}

// Delete returns the largest value from the top of the
// heap, deleting it.
func (h *MaxHeap) Delete() (string, error) {
	if h.size == 0 {
		return "", errors.New("Heap empty")
	}
	max := h.data[1]
	// Take the last item in the heap and make it the
	// new root, even though this is almost certainly
	// not the largest element...
	h.data[1] = h.data[h.size]
	h.size--

	parent := 1
	// ...and "sink" it down to its correct level in the heap.
	sink(h.data, parent, h.size)

	return max, nil
}

func sink(data []string, parent int, size int) {
	for parent*2 <= size {
		// Make child the index of the larger of the parent's two children.
		// But, only check the right child when one exists, otherwise we
		// are reading past the end of the slice.
		child := parent * 2
		if child+1 <= size && data[child+1] > data[child] {
			child++
		}
		// swap the child with the parent if the child is larger
		if data[parent] < data[child] {
			data[child], data[parent] = data[parent], data[child]
		} else {
			break
		}
		parent = child
	}
}

// Sort performs an in-place heap sort on the provided slice of string
func Sort(data []string) {
	if data == nil || len(data) <= 2 {
		return
	}
	size := len(data) - 1
	// Turn into a maxheap
	for i := size / 2; i >= 1; i-- {
		sink(data, i, size)
	}
	// Move max val to the end of the array and then re-heapify all of the array
	// except for the max at the end. Then move new max val to second-last slot
	// of the array and re-heapify. The move the new max val to the third-last
	// slot of the array...
	for size > 1 {
		data[1], data[size] = data[size], data[1]
		size--
		sink(data, 1, size)
	}
}
