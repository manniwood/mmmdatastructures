package set

import (
	"strconv"
	"testing"
)

func TestInt(t *testing.T) {
	s := New[int]()
	s.PutSlice([]int{1, 2, 3})
	for i := 4; i <= 6; i++ {
		s.Put(i)
	}
	for i := 1; i <= 6; i++ {
		if !s.Has(i) {
			t.Errorf("Expected %v to be in the set", i)
		}
	}
	for i := 4; i <= 6; i++ {
		s.Delete(i)
	}
	for i := 4; i <= 6; i++ {
		if s.Has(i) {
			t.Errorf("Did not expect %v to be in the set", i)
		}
	}
}

func TestString(t *testing.T) {
	s := New[string]()
	s.PutSlice([]string{"1", "2", "3"})
	for i := 4; i <= 6; i++ {
		k := strconv.Itoa(i)
		s.Put(k)
	}
	for i := 1; i <= 6; i++ {
		k := strconv.Itoa(i)
		if !s.Has(k) {
			t.Errorf("Expected %v to be in the set", k)
		}
	}
	for i := 4; i <= 6; i++ {
		k := strconv.Itoa(i)
		s.Delete(k)
	}
	for i := 4; i <= 6; i++ {
		k := strconv.Itoa(i)
		if s.Has(k) {
			t.Errorf("Did not expect %v to be in the set", k)
		}
	}
}
