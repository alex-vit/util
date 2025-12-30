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
	got := []string{}
	for _, v := range sp.values.a {
		got = append(got, v.B)
	}

	if !slices.Equal(want, got) {
		t.Fatalf("expected values of %v but was %v", want, sp.values)
	}
}

func TestSparseSetReplace(t *testing.T) {
	sp := NewSparseSet[string]()

	sp.Put(1, "1")
	if "1" != *sp.Get(1) {
		t.FailNow()
	}

	sp.Put(1, "2")
	if "2" != *sp.Get(1) {
		t.FailNow()
	}

	if sp.values.Len() > 1 {
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
		got := *sp.Get(k)
		if want != got {
			t.Fatalf("expected %d to contain %s but was %s", k, want, got)
		}
	}

	if nil != sp.Get(2*defaultCapacity) {
		t.Fatalf("expected no value at 666 but was %v", sp.Get(666))
	}
}

func TestGetOrPutZero(t *testing.T) {
	sp := NewSparseSet[string]()
	if nil != sp.Get(4) {
		t.FailNow()
	}

	var zero string
	if zero != *sp.GetOrPutZero(4) {
		t.FailNow()
	}

	sp.Put(4, "4")
	if "4" != *sp.GetOrPutZero(4) {
		t.FailNow()
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
		if kv.value != *sp.Get(kv.key) {
			t.Fatalf("expected %d to contain %s but was %s", kv.key, kv.value, *sp.Get(kv.key))
		}
	}
}

func TestSparseSetDelete(t *testing.T) {
	sp := NewSparseSet[int]()
	sp.Put(1, 5)
	sp.Delete(1)
	if nil != sp.Get(1) {
		t.FailNow()
	}
}

func TestSparseSetDeleteLast(t *testing.T) {
	sp := NewSparseSet[int]()
	sp.Put(1, 1)
	sp.Put(2, 2)
	sp.Delete(2)
	// happy if no panic

	if nil != sp.Get(2) {
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
			continue
		}
		want := fmt.Sprintf("%d-value", k)
		got := sp.Get(k)
		if want != *got {
			t.Fatalf("expected %d to contain %v but was %v", k, want, got)
		}
	}

}
