package projecteuler

import "strings"

// BigInt is a struct holding slice of digits in reversed order
type BigInt struct {
	digits []byte
}

// MakeBigInt constructs BigInt out of string
func MakeBigInt(input string) (result BigInt, err error) {
	l := len(input)
	result.digits = make([]byte, l)

	for i := l; i > 0; i-- {
		result.digits[l-i] = byte(input[i-1] - '0')
	}

	return
}

// CloneBigInt clones BigInt
func CloneBigInt(rhs *BigInt) (result BigInt) {
	result.digits = make([]byte, len(rhs.digits))
	copy(result.digits, rhs.digits)
	return
}

// Add adds
func Add(one BigInt, two BigInt) (result BigInt) {
	if len(one.digits) < len(two.digits) {
		return Add(two, one)
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

// AddTo adds to
func (bi *BigInt) AddTo(rhs *BigInt) {
	result := Add(*bi, *rhs)
	bi.digits = make([]byte, len(result.digits))
	copy(bi.digits, result.digits)
}

// String returns string representation
func (bi *BigInt) String() string {
	var sb strings.Builder

	for i := len(bi.digits); i > 0; i-- {
		sb.WriteRune(rune(bi.digits[i-1] + '0'))
	}

	return sb.String()
}
