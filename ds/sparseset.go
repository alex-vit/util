package ds

import (
	"math"

	"github.com/alex-vit/util"
)

const defaultCapacity = 100

// SparseSet is an int-keyed associative container with O(1) insert, lookup, and delete.
// Keys are non-negative ints (typically entity IDs). The dense array stores values contiguously,
// enabling cache-friendly iteration via Entries/Values, while the sparse array maps keys to
// dense indices. Get returns (value V, found bool) rather than a pointer.
type SparseSet[V any] struct {
	// sparse array: maps key → index into values (-1 = absent)
	valueIndices []int
	// dense array: stores (key, value) pairs contiguously
	values *Swapback[Tup[int, V]]
}

func NewSparseSet[V any]() *SparseSet[V] {
	valueIndices := make([]int, defaultCapacity)
	for i := range valueIndices {
		valueIndices[i] = -1
	}
	return &SparseSet[V]{
		valueIndices: valueIndices,
		values:       NewSwapbackPro[Tup[int, V]](defaultCapacity),
	}
}

// Put inserts or replaces the value for key.
func (sp *SparseSet[V]) Put(key int, value V) {
	sp.Delete(key)

	valueIndex := sp.values.Len()
	sp.values.Add(Tup[int, V]{key, value})

	if key > cap(sp.valueIndices)-1 {
		var newCap int = int(math.Round(float64(key) * 1.618))
		newValueIndices := make([]int, newCap)
		for i := range newValueIndices {
			newValueIndices[i] = -1
		}
		copy(newValueIndices, sp.valueIndices)
		sp.valueIndices = newValueIndices
	}

	sp.valueIndices[key] = valueIndex
}

// Get returns the value for key and whether it was found.
func (sp *SparseSet[V]) Get(key int) (value V, found bool) {
	if key >= len(sp.valueIndices) {
		return value, false
	}

	valueIndex := sp.valueIndices[key]
	if valueIndex == -1 {
		return value, false
	}

	return sp.values.Get(valueIndex).B, true
}

// GetAt returns a value at index i. You should prefer Get by id / key.
func (sp *SparseSet[V]) GetAt(i int) *V {
	return &sp.values.Get(i).B
}

// Delete removes the value for key. No-op if key is absent.
func (sp *SparseSet[V]) Delete(key int) {
	if key >= len(sp.valueIndices) {
		return
	}

	valueIndex := sp.valueIndices[key]
	if valueIndex == -1 {
		return
	}

	sp.values.Delete(valueIndex)
	sp.valueIndices[key] = -1

	if valueIndex >= sp.values.Len() {
		return
	}

	// Swapback has now moved the last value to this same index. Get the key this value belongs to.
	keyOfMovedValue := sp.values.Get(valueIndex).A
	// Update the value index for that key to point to the now moved value.
	sp.valueIndices[keyOfMovedValue] = valueIndex
}

// Entries returns all (key, value) pairs. The slice is the internal dense array — do not modify.
func (sp *SparseSet[V]) Entries() []Tup[int, V] {
	return sp.values.a
}

// Values returns a new slice containing just the values (without keys).
func (sp *SparseSet[V]) Values() []V {
	return util.Map(sp.values.a, func(tup Tup[int, V]) V { return tup.B })
}

func (sp *SparseSet[V]) Len() int {
	return sp.values.Len()
}
