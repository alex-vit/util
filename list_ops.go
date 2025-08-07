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

func FilterInPlace[E any](s []E, keep func(E) bool) []E {
	var n int
	for i := range s {
		if keep(s[i]) {
			s[n] = s[i]
			n++
		}
	}
	// clear to prevent memory leaks
	clear(s[n:])
	s = s[:n]
	return s
}
