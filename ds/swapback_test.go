package ds

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestAdd(t *testing.T) {
	sl := []int{}
	sw := NewSwapback[int]()
	for x := range 5 {
		sl = append(sl, x)
		sw.Add(x)
	}
	if !slices.Equal(sl, sw.a) {
		t.Fatalf("expected slice %v to equal swapback content %v", sl, sw.a)
	}
}

func TestDelete(t *testing.T) {
	sw := NewSwapback[int]()
	for x := range 5 {
		sw.Add(x + 1)
	}

	last := len(sw.a) - 1
	moved := sw.Delete(1)
	if last != moved {
		t.Fatalf("should have moved %d but was %d", last, moved)
	}

	newContent := []int{1, 5, 3, 4}
	if !slices.Equal(newContent, sw.a) {
		t.Fatalf("expected contents of %v but was %v", newContent, sw.a)
	}

	newLen := len(newContent)
	if newLen != len(sw.a) {
		t.Fatalf("should have trimmed length to %d but was %d", newLen, cap(sw.a))
	}

	extended := sw.a[:newLen+1]
	var zeroed int
	if extended[len(extended)-1] != zeroed {
		t.Fatalf("expected deleted element to get zeroed but was %d", extended[len(extended)-1])
	}
}

func TestDeleteAll(t *testing.T) {
	sw := NewSwapback[int]()
	for x := range 5 {
		sw.Add(x + 1)
	}
	for sw.Len() > 0 {
		sw.Delete(rand.IntN(sw.Len()))
	}
	// this is redundant because a true test fail should be a panic
	if sw.Len() != 0 {
		t.Fatalf("failed to delete all, %d items remained", sw.Len())
	}
}

func TestDeleteAllCustomCapacity(t *testing.T) {
	sw := NewSwapbackPro[int](100)
	for x := range 5 {
		sw.Add(x + 1)
	}
	for sw.Len() > 0 {
		sw.Delete(rand.IntN(sw.Len()))
	}
	// this is redundant because a true test fail should be a panic
	if sw.Len() != 0 {
		t.Fatalf("failed to delete all, %d items remained", sw.Len())
	}
}
