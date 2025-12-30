package ds

type BitSet64 uint64

func (b *BitSet64) Set(n uint8) {
	*b |= 1 << n
}

func (b *BitSet64) Has(n uint8) bool {
	return (*b & (1 << n)) != 0
}

func (b *BitSet64) Clear(n uint8) {
	*b &= ^(1 << n)
}

func (b *BitSet64) Toggle(n uint8) {
	*b ^= 1 << n
}
