package projecteuler

// DigitalNumber is immutable structure containing number and it's digits
type DigitalNumber struct {
	x      int
	digits []byte
}

// NewDigitalNumber constructs new DigitalNumber
func NewDigitalNumber(x int) (newDn DigitalNumber) {
	newDn = DigitalNumber{x: x}

	for x > 0 {
		currDigit := byte(x % 10)
		newDn.digits = append(newDn.digits, currDigit)
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

// NonZeroDigits returns true iff digits don't contain zero
func (dn DigitalNumber) NonZeroDigits() bool {
	for _, d := range dn.digits {
		if d == 0 {
			return false
		}
	}

	return true
}

// DifferentDigits returns true iff digits are different among themselves
func (dn DigitalNumber) DifferentDigits() bool {
	usedDigits := make(map[byte]struct{})

	for _, d := range dn.Digits() {
		if _, ok := usedDigits[d]; ok {
			return false
		}

		usedDigits[d] = struct{}{}
	}

	return true
}

// DifferentDigitCompositions returns true iff DigitalNumber's digits are not in the set, and adds them to it if so.
// Caller is responsible for making the set
func (dn DigitalNumber) DifferentDigitCompositions(usedDigits map[byte]struct{}) (differentDigits bool) {
	for _, k := range dn.Digits() {
		if _, ok := usedDigits[k]; ok {
			return
		}
	}

	differentDigits = true
	for _, k := range dn.Digits() {
		usedDigits[k] = struct{}{}
	}

	return
}

// ReplaceDigits replaces digits in a number with ones from the map (index->new digit)
func (dn DigitalNumber) ReplaceDigits(replacements map[byte]byte) (value int) {
	pow := 1
	value = 0

	for i := 0; i < len(dn.digits); i++ {
		currDigit := int(dn.digits[i])
		if newDigit, ok := replacements[byte(len(dn.digits)-i-1)]; ok {
			currDigit = int(newDigit)
		}

		value += pow * currDigit
		pow *= 10
	}

	return
}
