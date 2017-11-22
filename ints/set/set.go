// Package set implements a set for strings.
// It's just syntactic sugar around map[string]struct{}
package set

type Set map[int]struct{}

func New() Set {
	m := make(map[int]struct{})
	return Set(m)
}

func (s Set) Has(k int) bool {
	_, ok := s[k]
	return ok
}

func (s Set) Put(k int) {
	s[k] = struct{}{}
}

func (s Set) Del(k int) {
	delete(s, k)
}

func (s Set) PutSlice(ks []int) {
	for _, k := range ks {
		s[k] = struct{}{}
	}
}
