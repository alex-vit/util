package ds

import (
	"slices"
	"testing"
)

func TestSetAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(5)
	if !slices.Equal([]int{5}, s.Values()) {
		t.FailNow()
	}
	s.Add(5)
	if !slices.Equal([]int{5}, s.Values()) {
		t.FailNow()
	}
	s.Add(7)

	//TODO I think this is flaky? compare contents
	if !slices.Equal([]int{5, 7}, s.Values()) {
		t.FailNow()
	}
}

func TestSetRemove(t *testing.T) {
	s := NewSet[int]()
	if s.Remove(5) {
		t.FailNow()
	}
	s.Add(5)
	if !s.Remove(5) {
		t.FailNow()
	}
}

func TestSetContains(t *testing.T) {
	s := NewSet[int]()
	if s.Contains(666) {
		t.FailNow()
	}
	s.Add(666)
	if !s.Contains(666) {
		t.FailNow()
	}
	s.Remove(666)
	if s.Contains(666) {
		t.FailNow()
	}
}
