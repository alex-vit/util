package ds

import (
	"iter"
	"slices"
)

type Queue[V any] struct {
	a []V
}

func NewQueue[V any]() *Queue[V] {
	return &Queue[V]{[]V{}}
}

func NewQueueOf[V any](vs ...V) *Queue[V] {
	q := &Queue[V]{[]V{}}
	for _, v := range vs {
		q.Push(v)
	}
	return q
}

// An alias to [Queue.Push] becasue I keep forgetting it.
//
// Deprecated: use [Queue.Push].
func (q *Queue[V]) Add(v V) {
	q.Push(v)
}
func (q *Queue[V]) Push(v V) {
	q.a = append(q.a, v)
}

func (q *Queue[V]) Next() (v V, ok bool) {
	if len(q.a) == 0 {
		return v, false
	}

	v = q.a[0]
	q.a = q.a[1:]
	return v, true
}

func (q *Queue[V]) Iter() iter.Seq[V] {
	return func(yield func(V) bool) {
		for v, ok := q.Next(); ok; v, ok = q.Next() {
			if !yield(v) {
				return
			}
		}
	}
}

func (q Queue[V]) Len() int {
	return len(q.a)
}

func (q *Queue[V]) Clear() {
	clear(q.a)
}

type PQ[V any] struct {
	a    []V
	comp func(a, b V) int
}

func NewPQ[V any](comp func(a, b V) int) *PQ[V] {
	return &PQ[V]{[]V{}, comp}
}

func (q *PQ[V]) Push(v V) {
	q.a = append(q.a, v)
	slices.SortFunc(q.a, q.comp)
}

func (q *PQ[V]) Next() (v V, ok bool) {
	if len(q.a) == 0 {
		return v, false
	}

	v = q.a[0]
	q.a = q.a[1:]
	return v, true
}

func (q *PQ[V]) Clear() {
	clear(q.a)
}
