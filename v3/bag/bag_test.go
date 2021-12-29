package bag

import (
	"reflect"
	"sort"
	"testing"
)

func TestInt(t *testing.T) {
	b := New[int]()
	b.PutSlice([]int{1, 2, 3, 3})
	b.Put(4)
	b.Put(5)
	b.Put(6)
	b.Put(6)
	got := []int{}
	b.Iter(func(elem int){
		got = append(got, elem)
	})
	sort.Ints(got)
	want := []int{1, 2, 3, 3, 4, 5, 6, 6}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected want %#v to equal got %#v", want, got)
	}

	if !b.Has(3) {
		t.Errorf("should have 3")
	}

	b.Delete(3)
	if !b.Has(3) {
		t.Errorf("should have 3 after one deletion")
	}

	b.Delete(3)
	if b.Has(3) {
		t.Errorf("should NOT have 3 after two deletions")
	}
}

