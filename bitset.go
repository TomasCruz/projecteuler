package projecteuler

import "fmt"

type Bitset[T Int32Plus] struct {
	Slice        []T
	Bitsize      int
	ElementCount int
}

// NewBitset construct new Bitset, all the bits set to 0 (false)
func NewBitset[T Int32Plus](n T, bitsize int) Bitset[T] {
	return Bitset[T]{
		Slice:        make([]T, (int(n)+bitsize-1)/bitsize),
		Bitsize:      bitsize,
		ElementCount: int(n),
	}
}

// Get returns bool value on index
func (b Bitset[T]) Get(index int) bool {
	pos := index / b.Bitsize
	j := index % b.Bitsize
	return (b.Slice[pos] & (T(1) << j)) != 0
}

// All returns set of values in Bitset
func (b Bitset[T]) All() map[int]struct{} {
	m := map[int]struct{}{}

	nPos := 0
	for pos := 0; pos < len(b.Slice); pos++ {
		if b.Slice[pos] == 0 {
			continue
		}

		bit := T(1)
		for j := 0; j < b.Bitsize; j++ {
			if b.Slice[pos]&bit != 0 {
				m[nPos+j] = struct{}{}
			}
			bit <<= 1
		}

		nPos += b.Bitsize
	}

	return m
}

// Set sets value on index
func (b Bitset[T]) Set(index int, value bool) {
	pos := index / b.Bitsize
	j := index % b.Bitsize

	if value {
		b.Slice[pos] |= T(1) << j
	} else {
		k := j + 1
		m := ((T(1) << k) - 1) & b.Slice[pos]
		b.Slice[pos] = ((b.Slice[pos] >> k) << k) + m
	}
}

// Clone clones the Bitset
func (b Bitset[T]) Clone() Bitset[T] {
	sl := make([]T, len(b.Slice))
	for i := range b.Slice {
		sl[i] = b.Slice[i]
	}

	return Bitset[T]{
		Slice:   sl,
		Bitsize: b.Bitsize,
	}
}

func (b Bitset[T]) Union(other Bitset[T]) (Bitset[T], error) {
	if b.Bitsize != other.Bitsize {
		return Bitset[T]{}, fmt.Errorf("bitsets have to have same bitsize")
	}

	bTCount := b.ElementCount / b.Bitsize
	otherTCount := other.ElementCount / other.Bitsize
	if bTCount < otherTCount {
		return other.Union(b)
	}

	sl := make([]T, bTCount)
	pos := 0
	for ; pos < otherTCount; pos++ {
		sl[pos] = b.Slice[pos] | other.Slice[pos]
	}
	for ; pos < bTCount; pos++ {
		sl[pos] = b.Slice[pos]
	}

	return Bitset[T]{
		Slice:        sl,
		Bitsize:      b.Bitsize,
		ElementCount: b.ElementCount,
	}, nil
}
