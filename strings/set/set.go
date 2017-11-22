// Package set implements a set for strings.
// It's just syntactic sugar around map[string]struct{}
package set

type Set map[string]struct{}

func New() Set {
	m := make(map[string]struct{})
	return Set(m)
}

func (s Set) Has(k string) bool {
	_, ok := s[k]
	return ok
}

func (s Set) Put(k string) {
	s[k] = struct{}{}
}

func (s Set) Del(k string) {
	delete(s, k)
}

func (s Set) PutSlice(ks []string) {
	for _, k := range ks {
		s[k] = struct{}{}
	}
}
