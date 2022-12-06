package set

type Set[K comparable] struct {
	elements map[K]struct{}
}

func New[K comparable]() Set[K] {
	return Set[K] {
		elements: make(map[K]struct{}),
	}
}

func Of[K comparable](e ...K) Set[K] {
	s := New[K]()

	for _, val := range e {
		s.Put(val)
	}

	return s
}

func (s Set[K]) Put(e K) {
	s.elements[e] = struct{}{}
}

func (s Set[K]) Has(e K) bool {
	_, ok := s.elements[e]
	return ok
}

func (s Set[K]) Len() int {
	return len(s.elements)
}


