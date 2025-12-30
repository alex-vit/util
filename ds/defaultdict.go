package ds

import (
	"fmt"
	"iter"
	"strings"
)

type DefaultDict[K comparable, V any] struct {
	m             map[K]V
	createDefault func() V
}

func NewDefaultDict[K comparable, V any]() *DefaultDict[K, V] {
	mkZero := func() V {
		var v V
		return v
	}
	return NewDefaultDictF[K](mkZero)
}

func NewDefaultDictF[K comparable, V any](createDefault func() V) *DefaultDict[K, V] {
	return &DefaultDict[K, V]{m: make(map[K]V), createDefault: createDefault}
}

func (d *DefaultDict[K, V]) Get(k K) V {
	v, ok := d.m[k]
	if !ok {
		v = d.createDefault()
		d.m[k] = v
	}
	return v
}

func (d *DefaultDict[K, V]) Put(k K, v V) {
	d.m[k] = v
}

func (d *DefaultDict[K, V]) Clear() {
	clear(d.m)
}

func (d DefaultDict[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range d.m {
			if !yield(k) {
				return
			}
		}
	}
}

func (d *DefaultDict[K, V]) Values() []V {
	values := make([]V, 0, len(d.m))
	for _, v := range d.m {
		values = append(values, v)
	}
	return values
}

func (d DefaultDict[K, V]) String() string {
	var sb strings.Builder
	sb.WriteString("{\n")
	for k, v := range d.m {
		var ks any
		switch kv := any(k).(type) {
		case byte:
			ks = string(kv)
		case rune:
			ks = string(kv)
		default:
			ks = kv
		}
		fmt.Fprintf(&sb, "  %v: %v\n", ks, v)
	}
	sb.WriteString("}")
	return sb.String()
}
