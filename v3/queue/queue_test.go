package queue

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCreateInt(t *testing.T) {
	q, _ := New[int]()
	if q.head != -1 {
		t.Error("Expected Head to be -1, got ", q.head)
	}
	if q.tail != -1 {
		t.Error("Expected Tail to be -1, got ", q.tail)
	}
}

func TestEnqueueInt(t *testing.T) {
	q, _ := New[int]()
	q.Enqueue(5)
	if q.tail != -1 {
		t.Error("Expected Tail to be -1, got ", q.tail)
	}
	if q.head != 0 {
		t.Error("Expected Tail to be 0, got ", q.tail)
	}
}

func TestFillInt(t *testing.T) {
	q, _ := New[int]()
	for i := 1; i <= 32; i++ {
		q.Enqueue(i)
	}
	q.Enqueue(33)
	if q.capacity != 64 {
		t.Error("Expected capacity to double")
	}
}

func TestDrainInt(t *testing.T) {
	q, _ := New[int]()
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
	if q.Len() != 0 {
		t.Error("Expected queue length to be 0")
	}
	if !q.Empty() {
		t.Error("Expected queue to be empty")
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

func TestEnqueueSliceInt(t *testing.T) {
	q, _ := New[int]()
	q.EnqueueSlice([]int{1, 2, 3})
	if q.Len() != 3 {
		t.Error("Expected queue length to be 3")
	}
	for j := 1; j <= 3; j++ {
		i, err := q.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if i != j {
			t.Errorf("Expected i to be >>%v<<, got >>%v<<", j, i)
		}
	}
}

func TestCreateString(t *testing.T) {
	q, _ := New[string]()
	if q.head != -1 {
		t.Error("Expected Head to be -1, got ", q.head)
	}
	if q.tail != -1 {
		t.Error("Expected Tail to be -1, got ", q.tail)
	}
}

func TestEnqueueString(t *testing.T) {
	q, _ := New[string]()
	q.Enqueue("5")
	if q.tail != -1 {
		t.Error("Expected Tail to be -1, got ", q.tail)
	}
	if q.head != 0 {
		t.Error("Expected Tail to be 0, got ", q.tail)
	}
}

func TestFillString(t *testing.T) {
	q, _ := New[string]()
	for i := 1; i <= 32; i++ {
		q.Enqueue(fmt.Sprint(i))
	}
	q.Enqueue("33")
	if q.capacity != 64 {
		t.Error("Expected capacity to double")
	}
}

func TestDrainString(t *testing.T) {
	q, _ := New[string]()
	for i := 1; i <= 32; i++ {
		q.Enqueue(fmt.Sprint(i))
	}
	var i string
	var err error
	for j := 0; j < 32; j++ {
		i, err = q.Dequeue()
		if i != fmt.Sprint(j+1) {
			t.Error("Expected i to be ", j, ", got ", i)
		}
	}
	if q.Len() != 0 {
		t.Error("Expected queue length to be 0")
	}
	if !q.Empty() {
		t.Error("Expected queue to be empty")
	}
	i, err = q.Dequeue()
	if err == nil {
		t.Error("Expected err to be present")
	}
	for j := 1; j < 35; j++ {
		q.Enqueue(fmt.Sprint(j))
		i, err = q.Dequeue()
		if i != fmt.Sprint(j) {
			t.Error("Expected i to be ", j, ", got ", i)
		}
	}
}

func TestEnqueueSliceString(t *testing.T) {
	q, _ := New[string]()
	q.EnqueueSlice([]string{"1", "2", "3"})
	if q.Len() != 3 {
		t.Error("Expected queue length to be 3")
	}
	for j := 1; j <= 3; j++ {
		i, err := q.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		expected := strconv.Itoa(j)
		if i != expected {
			t.Errorf("Expected i to be >>%v<<, got >>%v<<", expected, i)
		}
	}
}
