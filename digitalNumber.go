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

// DigitOccurencies returns the map of digit occurencies
func (dn DigitalNumber) DigitOccurencies() (occurencies map[byte]int) {
	occurencies = make(map[byte]int)
	for _, d := range dn.Digits() {
		if occ, ok := occurencies[d]; !ok {
			occurencies[d] = 1
		} else {
			occurencies[d] = occ + 1
		}
	}

	return
}

// SameDigitSet returns true iff d2 has the same multiset of digits as dn
func (dn DigitalNumber) SameDigitSet(d2 DigitalNumber) bool {
	o1 := dn.DigitOccurencies()
	o2 := d2.DigitOccurencies()

	if o1 == nil || o2 == nil || len(o1) != len(o2) {
		return false
	}

	for d, occ := range o1 {
		if occ2, ok := o2[d]; !ok {
			return false
		} else if occ != occ2 {
			return false
		}
	}

	return true
}
