package ds

import (
	"math"
)

const defaultCapacity = 100

type SparseSet[V any] struct {
	// the 'sparse' bit
	valueIndices []int
	// 'dense' part
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

func (sp *SparseSet[V]) Put(key int, value V) *V {
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
	return sp.Get(key)
}

func (sp *SparseSet[V]) PutZero(key int) *V {
	var zero V
	return sp.Put(key, zero)
}

func (sp *SparseSet[V]) Get(key int) *V {
	if key >= len(sp.valueIndices) {
		return nil
	}

	valueIndex := sp.valueIndices[key]
	if valueIndex == -1 {
		return nil
	}

	return &sp.values.Get(valueIndex).B
}

// GetAt returns a value at index i. You should prefer Get by id / key.
func (sp *SparseSet[V]) GetAt(i int) *V {
	return &sp.values.Get(i).B
}

func (sp *SparseSet[V]) Contains(key int) bool {
	return sp.Get(key) != nil
}

func (sp *SparseSet[V]) GetOrPutZero(key int) *V {
	if !sp.Contains(key) {
		sp.PutZero(key)
	}
	return sp.Get(key)
}

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

func (sp *SparseSet[V]) Entries() []Tup[int, V] {
	return sp.values.a
}

func (sp *SparseSet[V]) Len() int {
	return sp.values.Len()
}
