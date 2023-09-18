package projecteuler

import "strings"

// BigInt is a struct holding slice of digits in reversed order
type BigInt struct {
	digits []byte
}

// MakeBigIntFromInt constructs BigInt out of int
func MakeBigIntFromInt(input int) (result BigInt) {
	result.digits = make([]byte, 0)

	for i := 0; input > 0; i++ {
		result.digits = append(result.digits, byte(input%10))
		input /= 10
	}

	return
}

// MakeBigInt constructs BigInt out of string
func MakeBigInt(input string) (result BigInt) {
	l := len(input)
	result.digits = make([]byte, l)

	for i := l; i > 0; i-- {
		result.digits[l-i] = byte(input[i-1] - '0')
	}

	return
}

// MakeZeroBigInt constructs zero BigInt
func MakeZeroBigInt() BigInt {
	return MakeBigInt("0")
}

// DigitCount returns digit count
func (bi BigInt) DigitCount() int {
	return len(bi.digits)
}

// Clone clones BigInt
func (bi BigInt) Clone() (result *BigInt) {
	result = &BigInt{digits: make([]byte, len(bi.digits))}
	copy(result.digits, bi.digits)
	return
}

// AddBigInts adds
func AddBigInts(one BigInt, two BigInt) (result BigInt) {
	if len(one.digits) < len(two.digits) {
		return AddBigInts(two, one)
	}

	if len(one.digits) != len(two.digits) {
		for i := len(two.digits); i < len(one.digits); i++ {
			two.digits = append(two.digits, byte(0))
		}
	}

	carry := byte(0)
	l := len(one.digits)
	result.digits = make([]byte, l)

	for i := 0; i < l; i++ {
		currDigit := carry + one.digits[i] + two.digits[i]

		if currDigit > 9 {
			currDigit -= 10
			carry = 1
		} else {
			carry = 0
		}

		result.digits[i] = currDigit
	}

	if carry == 1 {
		result.digits = append(result.digits, byte(1))
	}

	return
}

// AddTo adds to bi
func (bi *BigInt) AddTo(rhs BigInt) {
	result := AddBigInts(*bi, rhs)
	bi.digits = make([]byte, len(result.digits))
	copy(bi.digits, result.digits)
}

// Subtract subtracts from bi
func (bi *BigInt) Subtract(rhs BigInt) {
	var rDigits []byte
	biLength := len(bi.digits)
	rhsLength := len(rhs.digits)
	if rhsLength > biLength {
		return
	} else if rhsLength == biLength {
		rDigits = rhs.digits
	} else {
		rDigits = make([]byte, biLength)
		copy(rDigits, rhs.digits)
	}

	prev := byte(0)
	for i := 0; i < biLength; i++ {
		x := bi.digits[i] - prev - rDigits[i]

		prev = byte(0)
		if x > 10 {
			x += 10
			prev = byte(1)
		}

		bi.digits[i] = x
	}

	bi.removeLeadingZeroes()
}

// String returns string representation
func (bi BigInt) String() string {
	var sb strings.Builder

	for i := len(bi.digits); i > 0; i-- {
		sb.WriteRune(rune(bi.digits[i-1] + '0'))
	}

	return sb.String()
}

// Int returns int64 value of bi. If bi respresents too big number, int64 will overflow.
func (bi BigInt) Int() int64 {
	num := int64(0)
	for i := len(bi.digits); i > 0; i-- {
		num *= 10
		num += int64(bi.digits[i-1])
	}

	return num
}

// Digits returns copy of reversed digits
func (bi BigInt) Digits() []byte {
	digits := make([]byte, len(bi.digits))
	copy(digits, bi.digits)
	return digits
}

// DigitSum returns sum of the digits
func (bi BigInt) DigitSum() int {
	retValue := 0

	for _, x := range bi.digits {
		retValue += int(x)
	}

	return retValue
}

// MulPowTen multiplies BigInt with power of ten
func (bi *BigInt) MulPowTen(pow int) {
	if pow == 0 || (len(bi.digits) == 1 && bi.digits[0] == 0) {
		return
	}

	newDigits := make([]byte, len(bi.digits)+pow)
	for i := 0; i < len(bi.digits); i++ {
		newDigits[i+pow] = bi.digits[i]
	}

	bi.digits = newDigits
}

// MulBigInts multiplies BigInts
func MulBigInts(one BigInt, two BigInt) (result BigInt) {
	result.digits = make([]byte, 1)

	for i := 0; i < len(two.digits); i++ {
		temp := one.Clone()
		temp.MultiplyByDigit(two.digits[i])
		temp.MulPowTen(i)
		result = AddBigInts(result, *temp)
	}

	return
}

// PowBigInt returns power of BigInt
func (bi *BigInt) PowBigInt(pow int) {
	if pow == 0 {
		bi.digits = make([]byte, 1)
		bi.digits[0] = 1
		return
	} else if pow == 1 {
		return
	}

	temp := bi.Clone()
	for i := 1; i < pow; i++ {
		m := MulBigInts(*temp, *bi)
		temp = &m
	}

	bi.digits = temp.digits
}

// IsPalindrome returns true iff bi is a palindrome
func (bi BigInt) IsPalindrome() bool {
	limit := len(bi.digits) / 2
	for i := 0; i < limit; i++ {
		if bi.digits[i] != bi.digits[len(bi.digits)-i-1] {
			return false
		}
	}

	return true
}

// ReverseDigits reverses digits of bi
func (bi *BigInt) ReverseDigits() {
	limit := len(bi.digits) / 2
	for i := 0; i < limit; i++ {
		bi.digits[i], bi.digits[len(bi.digits)-i-1] = bi.digits[len(bi.digits)-i-1], bi.digits[i]
	}
}

// CompareBigInts returns -1, 0, 1 for l < r, l == r, l > r respectively
func CompareBigInts(l, r BigInt) int {
	lDigits := l.digits
	rDigits := r.digits
	lLen := len(lDigits)
	rLen := len(rDigits)

	if lLen < rLen {
		return -1
	} else if lLen > rLen {
		return 1
	}

	for i := lLen - 1; i >= 0; i-- {
		if lDigits[i] < rDigits[i] {
			return -1
		} else if lDigits[i] > rDigits[i] {
			return 1
		}
	}

	return 0
}

// Concatenate appends rhs digits to bi
func (bi *BigInt) Concatenate(rhs BigInt) {
	bi.digits = append(rhs.digits, bi.digits...)
}

func (bi *BigInt) MultiplyByDigit(d byte) {
	if d == 0 {
		bi.digits = []byte{0}
		return
	} else if d == 1 {
		return
	}

	carry := byte(0)
	for i := 0; i < len(bi.digits); i++ {
		currDigit := carry + bi.digits[i]*d

		if currDigit > 9 {
			carry = currDigit / 10
			currDigit %= 10
		} else {
			carry = 0
		}

		bi.digits[i] = currDigit
	}

	if carry != 0 {
		bi.digits = append(bi.digits, carry)
	}
}

func (bi BigInt) SquareRoot(decimalCount int) (root BigInt, decimalIndex int) {
	// 101 decimals
	// https://www.freecodecamp.org/news/find-square-root-of-number-calculate-by-hand/
	// STEP 1: Separate The Digits Into Pairs
	// STEP 2: Find The Largest Integer
	// STEP 3: Now Subtract That Integer
	// STEP 4: Let's Move To The Next Pair
	// STEP 5: Find The Right Match
	// STEP 6: Subtract Again

	digitSquares := make([]int, 10)
	for i := 0; i < 10; i++ {
		digitSquares[i] = i * i
	}

	pairs := bi.pairs()
	pairLength := len(pairs)
	pairIndex := 0
	endOfPairsHit := false
	twenty := MakeBigIntFromInt(20)
	remainder := MakeZeroBigInt()

	root = MakeZeroBigInt()
	for di := 0; di < decimalCount; {
		currPair := pair{}
		if !endOfPairsHit {
			if pairIndex == pairLength {
				endOfPairsHit = true
				decimalIndex = root.DigitCount()
			} else {
				currPair = pairs[pairIndex]
			}
		}
		pairIndex++

		remainder.MulPowTen(2)
		remainder.AddTo(MakeBigIntFromInt(currPair.value()))

		i := 1
		res := MulBigInts(root, twenty)
		prevProduct := MakeZeroBigInt()
		for ; ; i++ {
			if i == 10 {
				break
			}

			iBigInt := MakeBigIntFromInt(i)
			rClone := res.Clone()
			rClone.AddTo(iBigInt)
			product := MulBigInts(*rClone, iBigInt)
			// if rClone*i>remainder { i--; break}
			cmp := CompareBigInts(product, remainder)
			if cmp == 1 {
				break
			}
			prevProduct = product
		}

		remainder.Subtract(prevProduct)

		root.MulPowTen(1)
		root.AddTo(MakeBigIntFromInt(i - 1))

		if endOfPairsHit {
			di++
		}
	}

	return
}

type pair struct {
	a int
	b int
}

// value returns value of p
func (p pair) value() int {
	return 10*p.a + p.b
}

// pairs returns list of pair of bi's  digits
func (bi BigInt) pairs() []pair {
	l := len(bi.digits)
	start := l - 1
	first := pair{}
	if l%2 == 1 {
		first.b = int(bi.digits[start])
		start--
	} else {
		first.a = int(bi.digits[start])
		first.b = int(bi.digits[start-1])
		start -= 2
	}

	pairs := make([]pair, (start+1)/2+1)
	pairs[0] = first
	for pairIndex := 1; start > 0; start -= 2 {
		pairs[pairIndex] = pair{
			a: int(bi.digits[start]),
			b: int(bi.digits[start-1]),
		}
		pairIndex++
	}

	return pairs
}

// removeLeadingZeroes removes leading zeroes from bi
func (bi *BigInt) removeLeadingZeroes() {
	i := len(bi.digits) - 1
	for ; i >= 0 && bi.digits[i] == 0; i-- {
	}
	bi.digits = bi.digits[:i+1]
}
