package misc

type Set[E comparable] map[E]struct{}

func NewSet[E comparable](vals ...E) Set[E] {
	s := Set[E]{}
	s.Add(vals...)
	return s
}

func (s Set[E]) Add(vals ...E) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

func (s Set[E]) AddSet(s2 Set[E]) {
	for k := range s2 {
		s[k] = struct{}{}
	}
}
