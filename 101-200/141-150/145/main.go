package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 145; Reversible Numbers
Some positive integers n have the property that the sum [n + reverse(n)] consists entirely of odd (decimal) digits. For instance,
36 + 63 = 99 and 409 + 904 = 1313. We will call such numbers "reversible"; so 36, 63, 409, and 904 are reversible. Leading zeroes are
not allowed in either n or reverse(n).

There are 120 reversible numbers below one-thousand.

How many reversible numbers are there below one-billion (10^9)?
*/

func main() {
	var limit int

	if len(os.Args) > 1 {
		limit64, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("bad argument")
		}

		limit = int(limit64)
	} else {
		limit = 9
	}

	projecteuler.Timed(calc, limit)
}

type intSet map[int]struct{}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	reversibles := intSet{}
	possibleDigitPairs := constructPossibleDigitPairs()

	for i := 2; i <= limit; i++ {
		masks := constructMasks(i)
		for _, m := range masks {
			constructReversibles(m, possibleDigitPairs, reversibles)
		}
	}

	result = strconv.Itoa(len(reversibles))
	return
}

func constructMasks(numDigits int) []string {
	raw := []string{}

	half := numDigits / 2
	raw = constructMasksRec(half, "", raw)

	isOdd := false
	if numDigits%2 == 1 {
		isOdd = true
	}

	ret := []string{}
	for _, m := range raw {
		if isOdd && m[half-1] != 'c' {
			continue
		}

		var sb strings.Builder
		for j := 0; j < half; j++ {
			sb.WriteByte(m[half-1-j])
		}
		if isOdd {
			prevHalfMinusOne := byte('n')
			if half > 1 {
				prevHalfMinusOne = m[half-2]
			}
			sb.WriteByte(prevHalfMinusOne)
		}
		sb.WriteString(m)
		m = sb.String()

		if m[numDigits-2] == 'c' {
			continue
		}

		prevCarry := "n" + m[:len(m)-1]

		i := 0
		for ; i < half; i++ {
			cr := 0
			if prevCarry[i] == 'c' {
				cr = 1
			}

			cl := 0
			if prevCarry[numDigits-1-i] == 'c' {
				cl = 1
			}

			if cl != cr {
				break
			}
		}

		if i == half {
			ret = append(ret, m)
		}
	}

	return ret
}

func constructMasksRec(remainingDigits int, curr string, ret []string) []string {
	if remainingDigits == 0 {
		ret = append(ret, curr)
		return ret
	}

	ret = constructMasksRec(remainingDigits-1, "c"+curr, ret)
	return constructMasksRec(remainingDigits-1, "n"+curr, ret)
}

func constructPossibleDigitPairs() map[string][]intSet {
	ret := map[string][]intSet{}
	ret["nc"] = make([]intSet, 10)
	ret["nn"] = make([]intSet, 10)
	ret["cn"] = make([]intSet, 10)
	ret["cc"] = make([]intSet, 10)

	for k := range ret {
		for i := 0; i < 10; i++ {
			ret[k][i] = intSet{}
		}
	}

	// nc = {0:-1 1:-1 2:9 3:8 4:7,9 5:6,8 6:5,7,9 7:4,6,8 8:3,5,7,9 9:2,4,6,8}
	for i := 0; i < 10; i++ {
		for j := 11 - i; j < 10; j += 2 {
			ret["nc"][i][j] = struct{}{}
		}
	}

	// nn = {0:1,3,5,7,9 1:0,2,4,6,8 2:1,3,5,7 3:0,2,4,6 4:1,3,5 5:0,2,4 6:1,3 7:0,2 8:1 9:0}
	odd := true
	for i := 0; i < 10; i++ {
		odd = !odd
		start := 1
		if odd {
			start = 0
		}
		for j := start; j < 10-i; j += 2 {
			ret["nn"][i][j] = struct{}{}
		}
	}

	// cn = {0:0,2,4,6,8 1:1,3,5,7 2:0,2,4,6 3:1,3,5 4:0,2,4 5:1,3 6:0,2 7:1 8:0 9:-1}
	odd = true
	for i := 0; i < 10; i++ {
		odd = !odd
		start := 0
		if odd {
			start = 1
		}
		for j := start; j < 9-i; j += 2 {
			ret["cn"][i][j] = struct{}{}
		}
	}

	// cc = {0:-1 1:9 2:8 3:7,9 4:6,8 5:5,7,9 6:4,6,8 7:3,5,7,9 8:2,4,6,8 9:1,3,5,7,9}
	for i := 0; i < 10; i++ {
		for j := 10 - i; j < 10; j += 2 {
			ret["cc"][i][j] = struct{}{}
		}
	}

	return ret
}

func constructReversibles(mask string, possibleDigitPairs map[string][]intSet, reversibles intSet) {
	prevCarry := "n" + mask[:len(mask)-1]
	positionSlices := make([][]intSet, len(mask))
	half := len(mask) / 2
	for i := 0; i < half; i++ {
		var sb strings.Builder
		sb.WriteByte(prevCarry[i])
		sb.WriteByte(mask[i])
		positionSlices[i] = possibleDigitPairs[sb.String()]
	}

	power := 1
	x := 0
	revX := 0
	constructReversiblesRec(positionSlices, 0, len(mask), power, x, revX, reversibles)
}

func constructReversiblesRec(positionSlices [][]intSet, position, lenMask, power, x, revX int, reversibles intSet) {
	half := lenMask / 2
	if position == half {
		finishNumbers(x, revX, lenMask, reversibles)
		return
	}

	for digit := 0; digit < 10; digit++ {
		// if (position == 0 || position == half-1) && digit == 0 {
		if position == 0 && digit == 0 {
			continue
		}

		for currDigit := range positionSlices[position][digit] {
			// if (position == 0 || position == half-1) && currDigit == 0 {
			if position == 0 && currDigit == 0 {
				continue
			}

			constructReversiblesRec(positionSlices, position+1, lenMask, power*10, x+power*digit, revX*10+currDigit, reversibles)
		}
	}
}

func finishNumbers(x, revX, lenMask int, reversibles intSet) {
	power := 1
	for i := 0; i < lenMask/2; i++ {
		power *= 10
	}

	if lenMask%2 == 1 {
		// this only works for mask[lenMask%2] being 'n'
		for i := 0; i < 5; i++ {
			one := power*10*revX + power*i + x
			two := reverseNumber(one)
			reversibles[one] = struct{}{}
			reversibles[two] = struct{}{}
		}
	} else {
		one := power*revX + x
		two := reverseNumber(one)
		reversibles[one] = struct{}{}
		reversibles[two] = struct{}{}
	}
}

func reverseNumber(x int) int {
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev
}

/*
	no 5, no 9 digits

	2:
		c1c0 - no, p1 will be even
		c1n0 - no, places 1 and 0 must be same per carry
		n1c0 - no
		n1n0 - 63

	3:
		c2c1c0 - no, s2 will be even, p2 = s2(s0) + c1
		c2c1n0 - no, s1 will be even (same digits, no c0)
		c2n1c0 - 409
		n2c1c0 - no, s2+c1 > s0 == s2
		c2n1n0 - no, s0 == s2
		n2c1n0 - no, p2 will be even (c1+s2)
		n2n1c0 - no, s0 == s2
		n2n1n0 - no, p1 = s1 (even)

	4:
		c3c2c1c0 - no, s0 == s3, but p3 = s3 + c2, so even
		c3c2c1n0 - no, s1 == s2, but p2 = s2 + c1, so even
		c3c2n1c0 - no, p2 = s2 = s1, and p1 = s1 + c0, so even
		c3n2c1c0 - no, p1 = s1(s2) + c0 producing c1, p2 = s2 + c1 so should produce c2
		n3c2c1c0 - no, p3 = s3(s0) + c2 must produce c3 like s0 did
		c3c2n1n0 - no, p2 = s2(s1), n2
		c3n2c1n0 - no, p2 = s2(s1)+c1, p1 = s1 (produces carry)
		n3c2c1n0 - no, (2)
		c3n2n1c0 - no, p2 misses c1, so even
		n3c2n1c0 - no, p2 (3)
		n3n2c1c0 - no, p2 = s2(s1) + c1, p1 = s1 + c0 (3)
		3ns
		n3n2n1c0 - no, p2 = s2(s1), p1 = s1 + c0 (3)
		same nncn
		ncnn - (2)
		cnnn - no, p3 (3)
		nnnn - 1726
		...

		middle must have previous carry
		one before last must be no carry (2)
		same sum plus same carry must produce same carry
		all must be odd

	ncwc = {0:-1 1:-1 2:9 3:8 4:7,9 5:6,8 6:5,7,9 7:4,6,8 8:3,5,7,9 9:2,4,6,8}
	ncnc = {0:1,3,5,7,9 1:0,2,4,6,8 2:1,3,5,7 3:0,2,4,6 4:1,3,5 5:0,2,4 6:1,3 7:0,2 8:1 9:0}
	wcnc = {0:1,3,5,7 1:0,2,4,6 2:1,3,5 3:0,2,4 4:1,3 5:0,2 6:1 7:0 8:-1 9:-1}
	wcwc = {0:9 1:8 2:7,9 3:6,8 4:5,7,9 5:4,6,8 6:3,5,7,9 7:2,4,6,8 8:1,3,5,7,9 9:0,2,4,6,8}
*/
