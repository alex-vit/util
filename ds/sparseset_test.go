package ds

import (
	"fmt"
	"slices"
	"testing"
)

func TestSparseSetAdd(t *testing.T) {
	sp := NewSparseSet[string]()
	sp.Put(1, "1")
	sp.Put(2, "2")

	want := []string{"1", "2"}
	got := sp.Values()

	if !slices.Equal(want, got) {
		t.Fatalf("expected values %v but got %v", want, got)
	}
}

func TestSparseSetReplace(t *testing.T) {
	sp := NewSparseSet[string]()

	sp.Put(1, "1")
	if v, ok := sp.Get(1); !ok || v != "1" {
		t.FailNow()
	}

	sp.Put(1, "2")
	if v, ok := sp.Get(1); !ok || v != "2" {
		t.FailNow()
	}

	if sp.Len() > 1 {
		t.FailNow()
	}
}

func TestSparseSetGet(t *testing.T) {
	sp := NewSparseSet[string]()

	for k := range 10 {
		v := fmt.Sprintf("%d-value", k)
		sp.Put(k, v)
	}

	for k := range 10 {
		want := fmt.Sprintf("%d-value", k)
		got, ok := sp.Get(k)
		if !ok || want != got {
			t.Fatalf("expected %d to contain %s but was %s (found=%v)", k, want, got, ok)
		}
	}

	_, ok := sp.Get(2 * defaultCapacity)
	if ok {
		t.Fatalf("expected no value at key %d", 2*defaultCapacity)
	}
}

func TestSparseSetAddBeyondCapacity(t *testing.T) {
	sp := NewSparseSet[string]()
	if defaultCapacity != cap(sp.valueIndices) {
		t.Fatalf("wanted capacity %d but was %d", defaultCapacity, cap(sp.valueIndices))
	}

	kvs := []struct {
		key   int
		value string
	}{
		{2 * defaultCapacity, "in a galaxy"},
		{3 * defaultCapacity, "far far away"},
	}

	for _, kv := range kvs {
		sp.Put(kv.key, kv.value)
	}
	for _, kv := range kvs {
		got, ok := sp.Get(kv.key)
		if !ok || kv.value != got {
			t.Fatalf("expected %d to contain %s but was %s", kv.key, kv.value, got)
		}
	}
}

func TestSparseSetDelete(t *testing.T) {
	sp := NewSparseSet[int]()
	sp.Put(1, 5)
	sp.Delete(1)
	if _, ok := sp.Get(1); ok {
		t.FailNow()
	}
}

func TestSparseSetDeleteLast(t *testing.T) {
	sp := NewSparseSet[int]()
	sp.Put(1, 1)
	sp.Put(2, 2)
	sp.Delete(2)
	// happy if no panic

	if _, ok := sp.Get(2); ok {
		t.FailNow()
	}
}

// TestSparseSetGetAfterDeletes tests that value indices are updated correctly as the underlying Swapback moves things around.
func TestSparseSetGetAfterDeletes(t *testing.T) {
	sp := NewSparseSet[string]()
	for i := range 10 {
		k := i + 1
		sp.Put(k, fmt.Sprintf("%d-value", k))
	}

	toDelete := []int{5, 7}
	for _, k := range toDelete {
		sp.Delete(k)
	}

	for i := range 10 {
		k := i + 1
		if slices.Contains(toDelete, k) {
			if _, ok := sp.Get(k); ok {
				t.Fatalf("expected %d to be deleted but it was found", k)
			}
			continue
		}
		want := fmt.Sprintf("%d-value", k)
		got, ok := sp.Get(k)
		if !ok || want != got {
			t.Fatalf("expected %d to contain %v but was %v", k, want, got)
		}
	}
}

func TestSparseSetEntries(t *testing.T) {
	sp := NewSparseSet[string]()
	sp.Put(3, "three")
	sp.Put(7, "seven")

	entries := sp.Entries()
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries but got %d", len(entries))
	}
}
