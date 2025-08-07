package util

import (
	"fmt"
	"iter"
	"strings"
)

type DefaultDict[K comparable, V any] struct {
	M  map[K]V
	mk func() V
}

func NewDefaultDict[K comparable, V any]() *DefaultDict[K, V] {
	mkZero := func() V {
		var v V
		return v
	}
	return NewDefaultDictF[K](mkZero)
}

func NewDefaultDictF[K comparable, V any](mk func() V) *DefaultDict[K, V] {
	return &DefaultDict[K, V]{M: make(map[K]V), mk: mk}
}

func (d *DefaultDict[K, V]) Get(k K) V {
	v, ok := d.M[k]
	if !ok {
		v = d.mk()
		d.M[k] = v
	}
	return v
}

func (d *DefaultDict[K, V]) Put(k K, v V) {
	d.M[k] = v
}

func (d *DefaultDict[K, V]) Clear() {
	clear(d.M)
}

func (d DefaultDict[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range d.M {
			if !yield(k) {
				return
			}
		}
	}
}

func (d DefaultDict[K, V]) String() string {
	var sb strings.Builder
	sb.WriteString("{\n")
	for k, v := range d.M {
		var ks any
		switch kv := any(k).(type) {
		case byte:
			ks = string(kv)
		case rune:
			ks = string(kv)
		default:
			ks = kv
		}
		sb.WriteString(fmt.Sprintf("  %v: %v\n", ks, v))
	}
	sb.WriteString("}")
	return sb.String()
}
