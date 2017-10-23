package queue

import "testing"

func TestCreate(t *testing.T) {
	q := NewInt()
	if q.Head != -1 {
		t.Error("Expected Head to be -1, got ", q.Head)
	}
	if q.Tail != -1 {
		t.Error("Expected Tail to be -1, got ", q.Tail)
	}
}

func TestEnqueue(t *testing.T) {
	q := NewInt()
	q.Enqueue(5)
	if q.Tail != -1 {
		t.Error("Expected Tail to be -1, got ", q.Tail)
	}
	if q.Head != 0 {
		t.Error("Expected Tail to be 0, got ", q.Tail)
	}
}

func TestFill(t *testing.T) {
	q := NewInt()
	for i := 1; i <= 32; i++ {
		q.Enqueue(i)
	}
	q.Enqueue(33)
	if q.Capacity != 64 {
		t.Error("Expected capacity to double")
	}
}

func TestDrain(t *testing.T) {
	q := NewInt()
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
