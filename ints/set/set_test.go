package set

import (
	"testing"
)

func Test(t *testing.T) {
	s := New()
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
