package maxheap

import "testing"

func TestCreate(t *testing.T) {
	h := New()
	if h.size != 0 {
		t.Error("Expected size to be -1, got ", h.size)
	}
}

func TestInsert(t *testing.T) {
	h := New()
	var tests = []struct {
		insert string
		max    string
		want   []string
	}{
		{"05", "05", []string{"", "05"}},
		{"10", "10", []string{"", "10", "05"}},
		{"20", "20", []string{"", "20", "05", "10"}},
		{"07", "20", []string{"", "20", "07", "10", "05"}},
	}
	_, err := h.Peek()
	if err == nil {
		t.Error("Supposed to return error when peeking at empty max heap")
	}
	for _, test := range tests {
		h.Insert(test.insert)
		checkBackingSlice(t, test.want, h.data, h.size)
		i, _ := h.Peek()
		if i != test.max {
			t.Errorf("Expected max to be %v, got %v", test.insert, i)
		}
	}
}

func checkBackingSlice(t *testing.T, a []string, b []string, sz int) {
	expectedSize := len(a) - 1
	if sz != expectedSize {
		t.Error("Expected size to be %v, got %v", expectedSize, sz)
	}
	for i, x := range a {
		if a[i] != b[i] {
			t.Errorf("Expected %vth, element to be %v, got %v", i, x, b[i])
		}
	}
}

func TestDelete(t *testing.T) {
	h := New()
	inits := []string{"05", "10", "20", "07"}
	for _, i := range inits {
		h.Insert(i)
	}
	var tests = []struct {
		max  string
		want []string
	}{
		{"20", []string{"", "10", "07", "05"}},
		{"10", []string{"", "07", "05"}},
		{"07", []string{"", "05"}},
		{"05", []string{""}},
	}
	for _, test := range tests {
		i, _ := h.Delete()
		if i != test.max {
			t.Errorf("Expected max to be %v, got %v", test.max, i)
		}
		checkBackingSlice(t, test.want, h.data, h.size)
	}
	_, err := h.Delete()
	if err == nil {
		t.Error("Supposed to return error when deleting from empty max heap")
	}
}
