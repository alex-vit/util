package util

func Map[A, B any](as []A, f func(A) B) []B {
	bs := make([]B, 0, len(as))
	for _, a := range as {
		bs = append(bs, f(a))
	}
	return bs
}

func Filter[A any](as []A, keep func(A) bool) (filtered []A) {
	for _, a := range as {
		if keep(a) {
			filtered = append(filtered, a)
		}
	}
	return filtered
}

// TODO FilterInPlace
