package ds

import (
	"testing"
)

func TestBitSet(t *testing.T) {
	var bs BitSet64

	for x := range 64 {
		if (x & 1) != 0 {
			bs.Set(uint8(x))
		}
	}

	for x := range 64 {
		want := (x & 1) != 0
		if want != bs.Has(uint8(x)) {
			t.FailNow()
		}
	}

}

func TestBitSetToggle(t *testing.T) {
	var bs BitSet64
	bs.Set(1)

	bs.Toggle(1)
	if bs.Has(1) {
		t.FailNow()
	}

	bs.Toggle(1)
	if !bs.Has(1) {
		t.FailNow()
	}
}

func TestBitSetClear(t *testing.T) {
	var bs BitSet64
	bs.Set(1)

	bs.Clear(1)
	if bs.Has(1) {
		t.FailNow()
	}

	bs.Clear(1)
	if bs.Has(1) {
		t.FailNow()
	}

}
