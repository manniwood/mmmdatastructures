package stack

import "testing"

func TestCreate(t *testing.T) {
	s := New()
	if s.top != -1 {
		t.Error("Expected top to be -1, got ", s.top)
	}
}

func TestPush(t *testing.T) {
	s := New()
	s.Push("5")
	if s.top != 0 {
		t.Error("Expected top to be 0, got ", s.top)
	}
	str, err := s.Peek()
	if str != "5" {
		t.Error("Expected top stack value to be 5, got ", str)
	}
	if err != nil {
		t.Error("Expected err to be nil, got ", err)
	}
	// Peek again; should still be there.
	str, err = s.Peek()
	if str != "5" {
		t.Error("Expected top stack value to be 5, got ", str)
	}
	if err != nil {
		t.Error("Expected err to be nil, got ", err)
	}
}

func TestFill(t *testing.T) {
	s := New()
	for i := 1; i <= 32; i++ {
		s.Push(string(i))
	}
	s.Push("33")
	if s.capacity != 64 {
		t.Error("Expected capacity to double")
	}
}

func TestDrain(t *testing.T) {
	s := New()
	for i := 1; i <= 32; i++ {
		s.Push(string(i))
	}
	var str string
	var err error
	for i := 32; i > 0; i-- {
		str, err = s.Pop()
		if str != string(i) {
			t.Error("Expected str to be ", i, ", got ", str)
		}
	}
	str, err = s.Pop()
	if err == nil {
		t.Error("Expected err to be present")
	}
	for i := 1; i < 35; i++ {
		s.Push(string(i))
		str, err = s.Pop()
		if str != string(i) {
			t.Error("Expected str to be ", i, ", got ", str)
		}
	}
}
