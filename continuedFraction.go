package projecteuler

import (
	"fmt"
	"math"
	"strings"
)

type (
	// RootIntElement holds Head (integer part) and Fractions, rootIntFraction
	RootIntElement struct {
		Head      int
		Fractions rootIntFraction
	}

	// ContinuedFraction holds
	ContinuedFraction struct {
		r         int
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
	c.r = x
	c.RootFloor = int(math.Floor(math.Sqrt(float64(x))))
	c.primes = primes
	c.Head.Head = c.RootFloor
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

func (c ContinuedFraction) next(a RootIntElement) (ri RootIntElement) {
	c.rationalize(&a.Fractions)
	a.Fractions.reduce(c.primes)
	q := (a.Fractions.num.q*c.RootFloor + a.Fractions.num.i) / a.Fractions.denom.i
	ri.Head = q
	ri.Fractions.denom = a.Fractions.denom
	ri.Fractions.num.q = a.Fractions.num.q
	ri.Fractions.num.i = a.Fractions.num.i - q*a.Fractions.denom.i
	ri.Fractions.invert()
	return
}

func (c ContinuedFraction) calcConvergent(elementCount int, Fractions func(int) *RootIntElement) BigIntFraction {
	elements := make([]*RootIntElement, 0, elementCount+1)
	elements = append(elements, &c.Head)
	for i := 0; i < elementCount; i++ {
		elements = append(elements, Fractions(i))
	}

	return CalcElements(elements)
}

func (c ContinuedFraction) rationalize(r *rootIntFraction) {
	if r.denom.q == 0 {
		return
	}

	// (r.num.q*sqrt(c.r) + r.num.i)*(r.denom.q*sqrt(c.r) - r.denom.i)
	// q[r.num.i*r.denom.q - r.num.q*r.denom.i] + i[r.num.q*r.denom.q*c.r-r.num.i*r.denom.i]
	r.num.q, r.num.i = r.num.i*r.denom.q-r.num.q*r.denom.i, r.num.q*r.denom.q*c.r-r.num.i*r.denom.i

	// (r.denom.q*sqrt(c.r) + r.denom.i)*(r.denom.q*sqrt(c.r) - r.denom.i)
	// r.denom.q*r.denom.q*c.r - r.denom.i*r.denom.i
	r.denom.i = r.denom.q*r.denom.q*c.r - r.denom.i*r.denom.i
	r.denom.q = 0
}

/*
// Mul calculates and returns res, a product of a and b, rootIntPart structs
// containing a term involving sqrt(c.r), c bieng the reciever
func (c ContinuedFraction) Mul(a, b rootIntPart) (res rootIntPart) {
	// [r.q*sqrt(r) + r.i] * [b.q*sqrt(r) + b.i]
	// r.q*b.q*r.r + (r.q*b.i+b.q*r.i)*sqrt(r) + r.i*b.i
	ai := a.q*b.q*c.r + a.i*b.i
	aq := a.q*b.i + b.q*a.i
	res = rootIntPart{q: aq, i: ai}
	return
}
*/

// CalcElements calculates and returns BigIntFraction value of the continued fraction represented by elements
func CalcElements(elements []*RootIntElement) BigIntFraction {
	lastIndex := len(elements) - 1

	prev := MakeFraction(MakeBigIntFromInt(elements[lastIndex].Head), MakeBigIntFromInt(1))

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
