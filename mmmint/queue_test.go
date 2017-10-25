package mmmint

import "testing"

func TestCreate(t *testing.T) {
	q := New()
	if q.head != -1 {
		t.Error("Expected Head to be -1, got ", q.head)
	}
	if q.tail != -1 {
		t.Error("Expected Tail to be -1, got ", q.tail)
	}
}

func TestEnqueue(t *testing.T) {
	q := New()
	q.Enqueue(5)
	if q.tail != -1 {
		t.Error("Expected Tail to be -1, got ", q.tail)
	}
	if q.head != 0 {
		t.Error("Expected Tail to be 0, got ", q.tail)
	}
}

func TestFill(t *testing.T) {
	q := New()
	for i := 1; i <= 32; i++ {
		q.Enqueue(i)
	}
	q.Enqueue(33)
	if q.capacity != 64 {
		t.Error("Expected capacity to double")
	}
}

func TestDrain(t *testing.T) {
	q := New()
	for i := 1; i <= 32; i++ {
		q.Enqueue(i)
	}
	var i int
	var err error
	for j := 0; j < 32; j++ {
		i, err = q.Dequeue()
		if i != j+1 {
			t.Error("Expected i to be ", j, ", got ", i)
		}
	}
	i, err = q.Dequeue()
	if err == nil {
		t.Error("Expected err to be present")
	}
	for j := 1; j < 35; j++ {
		q.Enqueue(j)
		i, err = q.Dequeue()
		if i != j {
			t.Error("Expected i to be ", j, ", got ", i)
		}
	}
}
