package util

import (
	"fmt"
	"strings"
)

type Set[V comparable] struct {
	m map[V]struct{}
}

func NewSet[V comparable]() *Set[V] {
	return &Set[V]{m: make(map[V]struct{})}
}

func NewSetOf[V comparable](vs ...V) *Set[V] {
	s := &Set[V]{m: make(map[V]struct{})}
	for _, v := range vs {
		s.Add(v)
	}
	return s
}

func (s *Set[V]) Contains(v V) bool {
	_, found := s.m[v]
	return found
}

func (s *Set[V]) Add(v V) (added bool) {
	if s.Contains(v) {
		return false
	} else {
		s.m[v] = struct{}{}
		return true
	}
}

func (s *Set[V]) Remove(v V) (removed bool) {
	if s.Contains(v) {
		delete(s.m, v)
		return true
	} else {
		return false
	}
}

func (s *Set[V]) Len() int {
	return len(s.m)
}

func (s *Set[V]) Values() (values []V) {
	for k := range s.m {
		values = append(values, k)
	}
	return values
}

func (s *Set[V]) Clear() {
	clear(s.m)
}

func (s Set[V]) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	var i int
	for v := range s.m {
		if i == 0 {
			sb.WriteString(fmt.Sprintf("%v", v))
		} else {
			sb.WriteString(fmt.Sprintf(", %v", v))
		}
		i++
	}
	sb.WriteString("}")
	return sb.String()
}
