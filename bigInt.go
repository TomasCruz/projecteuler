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
func CloneBigInt(rhs BigInt) (result BigInt) {
	result.digits = make([]byte, len(rhs.digits))
	copy(result.digits, rhs.digits)
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

// AddTo adds to
func (bi *BigInt) AddTo(rhs BigInt) {
	result := AddBigInts(*bi, rhs)
	bi.digits = make([]byte, len(result.digits))
	copy(bi.digits, result.digits)
}

// String returns string representation
func (bi BigInt) String() string {
	var sb strings.Builder

	for i := len(bi.digits); i > 0; i-- {
		sb.WriteRune(rune(bi.digits[i-1] + '0'))
	}

	return sb.String()
}

// DigitSum returns sum of the digits
func (bi BigInt) DigitSum() int {
	retValue := 0

	for _, x := range bi.digits {
		retValue += int(x)
	}

	return retValue
}

func (bi *BigInt) mulDigit(d byte) {
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

// MulPowTen multiplies BigInt with power of ten
func (bi *BigInt) MulPowTen(pow int) {
	if pow == 0 {
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
		temp := CloneBigInt(one)
		tempPtr := &temp
		tempPtr.mulDigit(two.digits[i])
		tempPtr.MulPowTen(i)
		result = AddBigInts(result, *tempPtr)
	}

	return
}

// PowBigInt rewturns power of BigInt
func (bi *BigInt) PowBigInt(pow int) {
	if pow == 0 {
		bi.digits = make([]byte, 1)
		bi.digits[0] = 1
		return
	} else if pow == 1 {
		return
	}

	temp := CloneBigInt(*bi)
	for i := 1; i < pow; i++ {
		temp = MulBigInts(temp, *bi)
	}

	bi.digits = temp.digits
}
