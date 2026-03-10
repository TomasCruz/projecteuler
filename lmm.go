package projecteuler

import (
	"math"
	"slices"
)

func LMM(d, n int64) (PellTriplet, []PellTriplet) {
	one, minusOne := getOneAndMinusOne(d)
	solutionSet := map[PellTriplet]struct{}{}

	primes := Primes(int(n), nil)
	sqDivisors := squareDivisors(int(n), primes)

	for _, f := range sqDivisors {
		m := n / (int64(f) * int64(f))
		mHalf := float64(abs(m)) / 2.0

		for z := int64(math.Ceil(-mHalf)); z <= int64(mHalf); z++ {
			zSq := z * z
			dModM := d % m
			if zSq%m != dModM {
				continue
			}

			_, values := PQa(z, m, d)
			for i := range values[3] {
				g := int64(values[3][i])
				b := int64(values[2][i])
				currN := g*g - 5*b*b

				var t PellTriplet
				switch currN {
				case n:
					t = PellTriplet{A: g, B: b, C: currN}
				case m:
					t = PellTriplet{A: int64(f) * g, B: int64(f) * b, C: int64(f) * int64(f) * currN}
				case -m:
					if minusOne != nil {
						t = ComposePellTriplets(d, PellTriplet{A: g, B: b, C: currN}, *minusOne)
						t = PellTriplet{A: int64(f) * t.A, B: int64(f) * t.B, C: int64(f) * int64(f) * t.C}
					}
				default:
					continue
				}

				if t.A < 0 || t.B < 0 {
					continue
				}

				solutionSet[t] = struct{}{}
			}
		}
	}

	unprunned := make([]PellTriplet, 0, len(solutionSet))
	for k := range solutionSet {
		unprunned = append(unprunned, k)
	}
	slices.SortFunc(unprunned, TripletSortFunc)

	ret := make([]PellTriplet, 0, len(solutionSet))
	ret = pruneEquivalents(d, n, unprunned)

	return one, ret
}

func PQa(p, q, d int64) (int64, [][]int64) {
	t := abs(q)
	if (d-p*p)%t != 0 {
		p *= t
		q *= t
		d *= t * t
	}

	dRoot := math.Sqrt(float64(d))

	a := []int64{}
	A := []int64{}
	B := []int64{}
	G := []int64{}
	a0, a1 := int64(0), int64(1)
	b0, b1 := int64(1), int64(0)
	g0, g1 := -p, q
	p1, q1 := p, q
	P := []int64{p}
	Q := []int64{q}

	period := int64(0)
	pe := 0
	ps := 0
	var pr, qr int64

	for {
		t := int64(math.Floor((float64(p1) + dRoot) / float64(q1)))
		a = append(a, t)

		a0, a1 = a1, t*a1+a0
		b0, b1 = b1, t*b1+b0
		g0, g1 = g1, t*g1+g0
		A = append(A, a1)
		B = append(B, b1)
		G = append(G, g1)

		p1 = t*q1 - p1
		q1 = (d - p1*p1) / q1
		P = append(P, p1)
		Q = append(Q, q1)

		plus := (float64(p1) + dRoot) / float64(q1)
		minus := (float64(p1) - dRoot) / float64(q1)

		if period != 0 {
			if p1 == pr && q1 == qr {
				pe = len(a)
				u := a[ps:pe]
				uA := A[ps:pe]
				uB := B[ps:pe]
				uG := G[ps:pe]
				uP := P[ps:pe]
				uQ := Q[ps:pe]
				return period, [][]int64{u, uA, uB, uG, uP, uQ}
			} else if ps == 0 {
				period++
			}
		} else if plus > 1 && -1 < minus && minus < 0 {
			pr, qr = p1, q1
			period = 1
		}
	}
}

func ComposePellTriplets(d int64, one, two PellTriplet) PellTriplet {
	return PellTriplet{
		A: one.A*two.A + d*one.B*two.B,
		B: one.A*two.B + one.B*two.A,
		C: one.C * two.C,
	}
}

type PellTriplet struct {
	A, B, C int64
}

type powers struct {
	base int
	exp  []int
}

func TripletSortFunc(a, b PellTriplet) int {
	switch {
	case a.A < b.A:
		return -1
	case a.A > b.A:
		return 1
	default:
		return 0
	}
}

func getOneAndMinusOne(d int64) (PellTriplet, *PellTriplet) {
	_, values := PQa(0, 1, d)
	l := len(values[3])

	var o PellTriplet
	var m *PellTriplet

	gotOne := false
	gotMinusOne := false
	for i := range l {
		g := int64(values[3][i])
		b := int64(values[2][i])
		n := g*g - d*b*b

		if !gotOne && n == int64(1) {
			o = PellTriplet{A: g, B: b, C: n}
			gotOne = true
		}

		if !gotMinusOne && n == int64(-1) {
			m = &PellTriplet{A: g, B: b, C: n}
			gotMinusOne = true
		}

		// fmt.Printf("%d^2 - %d*%d^2 = %d\n", g, d, b, n)
	}

	if !gotOne && gotMinusOne {
		o = ComposePellTriplets(d, *m, *m)
	}

	return o, m
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func pruneEquivalents(d, n int64, fund []PellTriplet) []PellTriplet {
	l := len(fund)
	eq := make([]bool, l)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if pellTripletsAreEquivalent(d, n, fund[i], fund[j]) {
				eq[j] = true
			}
		}
	}

	ret := make([]PellTriplet, 0, l)
	for i := range fund {
		if !eq[i] {
			ret = append(ret, fund[i])
		}
	}

	return ret
}

func pellTripletsAreEquivalent(d, n int64, a, b PellTriplet) bool {
	// (x,y) and (r,s) are equivalent if (xr – Dys)/N and (xs – yr)/N are both integers
	if (a.A*b.A-d*a.B*b.B)%n == 0 && (a.A*b.B-a.B*b.A)%n == 0 {
		return true
	}

	return false
}

func squareDivisors(n int, primes []int) []int {
	factors, _ := Factorize(n, primes)

	factorPossibilities := []powers{}
	j := 0
	for k, v := range factors {
		maxSqExp := v / 2
		if maxSqExp != 0 {
			factorPossibilities = append(factorPossibilities, powers{})
			factorPossibilities[j].base = k
			factorPossibilities[j].exp = make([]int, maxSqExp+1)
			for i := 0; i <= maxSqExp; i++ {
				prod := 1
				for range i {
					prod *= k
				}
				factorPossibilities[j].exp[i] = prod
			}
			j++
		}
	}

	slices.SortFunc(factorPossibilities, func(a, b powers) int {
		switch {
		case a.base < b.base:
			return -1
		case a.base > b.base:
			return 1
		default:
			return 0
		}
	})

	divCount := 1
	for _, f := range factorPossibilities {
		divCount *= len(f.exp) - 1
	}
	divCount++

	divisors := make([]int, 0, divCount)
	divisors = mulFactorsRec(factorPossibilities, 0, 1, divisors)
	slices.Sort(divisors)

	return divisors
}

func mulFactorsRec(factorPossibilities []powers, factorIndex, currProd int, divisors []int) []int {
	if factorIndex == len(factorPossibilities) {
		divisors = append(divisors, currProd)
		return divisors
	}

	for _, p := range factorPossibilities[factorIndex].exp {
		currFactor := currProd * p
		divisors = mulFactorsRec(factorPossibilities, factorIndex+1, currFactor, divisors)
	}

	return divisors
}
