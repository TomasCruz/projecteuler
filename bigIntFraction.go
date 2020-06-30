package projecteuler

// BigIntFraction struct
type BigIntFraction struct {
	numerator, denominator *BigInt
}

// MakeFraction creates and returns fr
func MakeFraction(numerator, denominator BigInt) (fr BigIntFraction) {
	fr.numerator = numerator.Clone()
	fr.denominator = denominator.Clone()

	return
}

// AddInt adds x to receiver
func (fr *BigIntFraction) AddInt(x int) {
	toAdd := MakeBigIntFromInt(x)
	toAdd = MulBigInts(toAdd, *fr.denominator)
	fr.numerator.AddTo(toAdd)
}

// Numerator returns receiver's numerator
func (fr BigIntFraction) Numerator() BigInt {
	return *fr.numerator
}

// Denominator returns receiver's denominator
func (fr BigIntFraction) Denominator() BigInt {
	return *fr.denominator
}

// Invert flips receiver's numerator and denominator
func (fr *BigIntFraction) Invert() {
	fr.numerator, fr.denominator = fr.denominator, fr.numerator
}
