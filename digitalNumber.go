package projecteuler

// DigitalNumber has number and it's digits
type DigitalNumber struct {
	x        int
	digits   []byte
	digitMap map[byte]struct{}
}

// NewDigitalNumber constructs new DigitalNumber
func NewDigitalNumber(x int) (newDn DigitalNumber) {
	newDn = DigitalNumber{x: x, digitMap: make(map[byte]struct{})}

	for x > 0 {
		currDigit := byte(x % 10)
		newDn.digits = append(newDn.digits, currDigit)
		newDn.digitMap[currDigit] = struct{}{}
		x /= 10
	}

	return
}

// X returns number
func (dn DigitalNumber) X() int {
	return dn.x
}

// Digits returns digits (reference)
func (dn DigitalNumber) Digits() []byte {
	return dn.digits
}

// DigitMap returns digitMap (reference)
func (dn DigitalNumber) DigitMap() map[byte]struct{} {
	return dn.digitMap
}

// DigitCount returns number of digits
func (dn DigitalNumber) DigitCount() int {
	return len(dn.digits)
}

// NumberFromDigits calculates number from its reversed digits
func NumberFromDigits(digits []byte) (value int) {
	pow, value := 1, 0

	for i := 0; i < len(digits); i++ {
		value += pow * int(digits[i])
		pow *= 10
	}

	return
}
