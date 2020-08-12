package projecteuler

import (
	"math/big"
)

// BigIntFraction struct
type BigIntFraction struct {
	numerator, denominator *big.Int
}

// MakeFraction creates and returns fr
func MakeFraction(numerator, denominator *big.Int) (fr BigIntFraction) {
	fr.numerator = big.NewInt(1)
	fr.denominator = big.NewInt(1)
	fr.numerator.Set(numerator)
	fr.denominator.Set(denominator)

	return
}

// AddInt adds x to receiver
func (fr *BigIntFraction) AddInt(x int64) {
	toAdd := big.NewInt(x)
	toAdd.Mul(toAdd, fr.denominator)
	fr.numerator.Add(fr.numerator, toAdd)
}

// Numerator returns receiver's numerator
func (fr BigIntFraction) Numerator() *big.Int {
	return fr.numerator
}

// Denominator returns receiver's denominator
func (fr BigIntFraction) Denominator() *big.Int {
	return fr.denominator
}

// Invert flips receiver's numerator and denominator
func (fr *BigIntFraction) Invert() {
	fr.numerator, fr.denominator = fr.denominator, fr.numerator
}
