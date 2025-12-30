package ds

type Swapback[E any] struct {
	a []E
}

func NewSwapback[E any]() *Swapback[E] {
	return &Swapback[E]{
		a: []E{},
	}
}

func NewSwapbackPro[E any](initCapacity int) *Swapback[E] {
	return &Swapback[E]{
		a: make([]E, 0, initCapacity),
	}
}

func (s *Swapback[E]) Add(e E) {
	s.a = append(s.a, e)
}

func (s *Swapback[E]) Get(i int) *E {
	return &s.a[i]
}

// Delete deletes the item at index i.
// Returns imoved - the index of the item that now takes the removed item's place.
func (s *Swapback[E]) Delete(i int) (imoved int) {
	last := len(s.a) - 1
	// place the item to remove last
	s.a[i], s.a[last] = s.a[last], s.a[i]
	// null removed element to avoid memory leaks
	clear(s.a[last:])
	// trim the array to exclude the last, removed element
	s.a = s.a[:last]
	return last
}

func (s *Swapback[E]) Len() int { return len(s.a) }
