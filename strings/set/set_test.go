package set

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	s := New()
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
		s.Del(k)
	}
	for i := 4; i <= 6; i++ {
		k := strconv.Itoa(i)
		if s.Has(k) {
			t.Errorf("Did not expect %v to be in the set", k)
		}
	}
}
