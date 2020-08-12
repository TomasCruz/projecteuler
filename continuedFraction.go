package projecteuler

import (
	"fmt"
	"math"
	"math/big"
	"strings"
)

type (
	// RootIntElement holds Head (integer part) and Fractions, rootIntFraction
	RootIntElement struct {
		Head      int64
		Fractions rootIntFraction
	}

	// ContinuedFraction holds
	ContinuedFraction struct {
		root      int
		RootFloor int
		primes    []int
		Head      RootIntElement
		Fractions []RootIntElement
	}

	rootIntPart struct {
		q int
		i int
	}

	rootIntFraction struct {
		num   rootIntPart
		denom rootIntPart
	}
)

// MakeContinuedFraction creates and returns c, a continued fraction representation of sqrt(x).
// primes is a slice containing primes. It is necessary for gcd calculations important for reducing fractions
func MakeContinuedFraction(x int, primes []int) (c ContinuedFraction) {
	c.root = x
	c.RootFloor = int(math.Floor(math.Sqrt(float64(x))))
	c.primes = primes
	c.Head.Head = int64(c.RootFloor)
	c.Head.Fractions = rootIntFraction{num: makeRIPart(0, 1), denom: makeRIPart(1, -c.RootFloor)}

	c.Fractions = append(c.Fractions, c.next(c.Head))
	for {
		nextRie := c.next(c.Fractions[len(c.Fractions)-1])
		if nextRie == c.Fractions[0] {
			break
		}

		c.Fractions = append(c.Fractions, nextRie)
	}

	return
}

// String returns concise receiver's string representation
func (c ContinuedFraction) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[%d; (%d", c.Head.Head, c.Fractions[0].Head))
	for i := 1; i < len(c.Fractions); i++ {
		sb.WriteString(fmt.Sprintf(",%d", c.Fractions[i].Head))
	}
	sb.WriteString(")]")
	return sb.String()
}

// Period returns a period of receiver
func (c ContinuedFraction) Period() int {
	return len(c.Fractions)
}

// Convergent calculates BigIntFraction value of elementCount convergent
func (c ContinuedFraction) Convergent(elementCount int) BigIntFraction {
	elements := make([]*RootIntElement, 0, elementCount+1)
	elements = append(elements, &c.Head)
	for i := 0; i < elementCount; i++ {
		elements = append(elements, &c.Fractions[i%len(c.Fractions)])
	}

	return CalcElements(elements)
}

func (c ContinuedFraction) next(a RootIntElement) (ri RootIntElement) {
	c.rationalize(&a.Fractions)
	a.Fractions.reduce(c.primes)
	q := (a.Fractions.num.q*c.RootFloor + a.Fractions.num.i) / a.Fractions.denom.i
	ri.Head = int64(q)
	ri.Fractions.denom = a.Fractions.denom
	ri.Fractions.num.q = a.Fractions.num.q
	ri.Fractions.num.i = a.Fractions.num.i - q*a.Fractions.denom.i
	ri.Fractions.invert()
	return
}

func (c ContinuedFraction) rationalize(r *rootIntFraction) {
	if r.denom.q == 0 {
		return
	}

	// (r.num.q*sqrt(c.r) + r.num.i)*(r.denom.q*sqrt(c.r) - r.denom.i)
	// q[r.num.i*r.denom.q - r.num.q*r.denom.i] + i[r.num.q*r.denom.q*c.r-r.num.i*r.denom.i]
	r.num.q, r.num.i = r.num.i*r.denom.q-r.num.q*r.denom.i, r.num.q*r.denom.q*c.root-r.num.i*r.denom.i

	// (r.denom.q*sqrt(c.r) + r.denom.i)*(r.denom.q*sqrt(c.r) - r.denom.i)
	// r.denom.q*r.denom.q*c.r - r.denom.i*r.denom.i
	r.denom.i = r.denom.q*r.denom.q*c.root - r.denom.i*r.denom.i
	r.denom.q = 0
}

/*
// Mul calculates and returns res, a product of a and b, rootIntPart structs
// containing a term involving sqrt(c.r), c bieng the reciever
func (c ContinuedFraction) Mul(a, b rootIntPart) (res rootIntPart) {
	// [a.q*sqrt(c.root) + a.i] * [b.q*sqrt(r) + b.i]
	// a.q*b.q*c.root + (a.q*b.i+b.q*a.i)*sqrt(r) + a.i*b.i
	ai := a.q*b.q*c.root + a.i*b.i
	aq := a.q*b.i + b.q*a.i
	res = rootIntPart{q: aq, i: ai}
	return
}
*/

// CalcElements calculates and returns BigIntFraction value of the continued fraction represented by elements
func CalcElements(elements []*RootIntElement) BigIntFraction {
	lastIndex := len(elements) - 1

	prev := MakeFraction(big.NewInt(elements[lastIndex].Head), big.NewInt(1))

	for i := lastIndex; i > 0; i-- {
		calcElem(&prev, elements[i-1])
	}

	return prev
}

func makeRIPart(q, i int) (ri rootIntPart) {
	ri.q, ri.i = q, i
	return
}

func (r rootIntPart) sqDiffComplement() (res rootIntPart) {
	res.q, res.i = r.q, -r.i
	return
}

func (r *rootIntFraction) invert() {
	r.num, r.denom = r.denom, r.num
}

func (r *rootIntFraction) reduce(primes []int) {
	x := gcd(r.num.q, r.num.i, r.denom.i, primes)
	r.num.q /= x
	r.num.i /= x
	r.denom.i /= x
}

func gcd(a, b, c int, primes []int) int {
	fa, _ := Factorize(a, primes)
	fb, _ := Factorize(b, primes)
	fc, _ := Factorize(c, primes)
	resMap := make(map[int]int)
	for p, e := range fa {
		eb, ok := fb[p]
		if !ok {
			continue
		}

		ec, ok := fc[p]
		if !ok {
			continue
		}

		resMap[p] = min(e, eb, ec)
	}

	res := 1
	for p, e := range resMap {
		for i := 0; i < e; i++ {
			res *= p
		}
	}

	return res
}

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}

	return m
}

func calcElem(prev *BigIntFraction, el *RootIntElement) {
	prev.Invert()
	prev.AddInt(el.Head)
	return
}
