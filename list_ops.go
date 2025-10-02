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

func Last[E any, S ~[]E](s S) E {
	return s[len(s)-1]
}

func LastN[E any, S ~[]E](s S, n int) S {
	return s[len(s)-n:]
}

func Clip[E any, S ~[]E](s *S, n int) {
	clear((*s)[len(*s)-n:])
	*s = (*s)[:len(*s)-n]
}

func RemoveLast[E any, S ~[]E](s *S) E {
	last := Last(*s)
	Clip(s, 1)
	return last
}

func RemoveLastN[E any, S ~[]E](s *S, n int) S {
	last := make(S, 0, n)
	// copy values bc the originals get zeroed out by [clip]
	last = append(last, LastN(*s, n)...)
	Clip(s, n)
	return last
}
