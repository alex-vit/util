package util

type Tup[A, B any] struct {
	A A
	B B
}

// D destructures the tuple into two variables
func (t *Tup[A, B]) D() (A, B) {
	return t.A, t.B
}
