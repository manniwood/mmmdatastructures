package stack

import (
	"fmt"
	"testing"
)

func TestCreateInt(t *testing.T) {
	s := New[int]()
	if s.top != -1 {
		t.Error("Expected top to be -1, got ", s.top)
	}
}

func TestPushInt(t *testing.T) {
	s := New[int]()
	s.Push(5)
	if s.top != 0 {
		t.Error("Expected top to be 0, got ", s.top)
	}
	integer, err := s.Peek()
	if integer != 5 {
		t.Error("Expected top stack value to be 5, got ", integer)
	}
	if err != nil {
		t.Error("Expected err to be nil, got ", err)
	}
	// Peek again; should still be there.
	integer, err = s.Peek()
	if integer != 5 {
		t.Error("Expected top stack value to be 5, got ", integer)
	}
	if err != nil {
		t.Error("Expected err to be nil, got ", err)
	}
}

func TestFillInt(t *testing.T) {
	s := New[int]()
	for i := 1; i <= 32; i++ {
		s.Push(i)
	}
	s.Push(33)
	if s.capacity != 64 {
		t.Error("Expected capacity to double")
	}
}

func TestDrainInt(t *testing.T) {
	s := New[int]()
	for i := 1; i <= 32; i++ {
		s.Push(i)
	}
	var integer int
	var err error
	for i := 32; i > 0; i-- {
		integer, err = s.Pop()
		if integer != i {
			t.Error("Expected integer to be ", i, ", got ", integer)
		}
	}
	integer, err = s.Pop()
	if err == nil {
		t.Error("Expected err to be present")
	}
	for i := 1; i < 35; i++ {
		s.Push(i)
		integer, err = s.Pop()
		if integer != i {
			t.Error("Expected integer to be ", i, ", got ", integer)
		}
	}
}

func TestCreateString(t *testing.T) {
	s := New[string]()
	if s.top != -1 {
		t.Error("Expected top to be -1, got ", s.top)
	}
}

func TestPushString(t *testing.T) {
	s := New[string]()
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

func TestFillString(t *testing.T) {
	s := New[string]()
	for i := 1; i <= 32; i++ {
		s.Push(fmt.Sprint(i))
	}
	s.Push("33")
	if s.capacity != 64 {
		t.Error("Expected capacity to double")
	}
}

func TestDrainString(t *testing.T) {
	s := New[string]()
	for i := 1; i <= 32; i++ {
		s.Push(fmt.Sprint(i))
	}
	var str string
	var err error
	for i := 32; i > 0; i-- {
		str, err = s.Pop()
		if str != fmt.Sprint(i) {
			t.Error("Expected str to be ", i, ", got ", str)
		}
	}
	str, err = s.Pop()
	if err == nil {
		t.Error("Expected err to be present")
	}
	for i := 1; i < 35; i++ {
		s.Push(fmt.Sprint(i))
		str, err = s.Pop()
		if str != fmt.Sprint(i) {
			t.Error("Expected str to be ", i, ", got ", str)
		}
	}
}
