package projecteuler

import "fmt"

// GCD returns greatest common divisor of m and n.
// GCD(m, n) == GCD(n, m % n)
func GCD(m, n int64) int64 {
	if m < n {
		m, n = n, m
	}

	for n > 0 {
		m, n = n, m%n
	}

	return m
}

type Fraction struct {
	Num   int64
	Denom int64
}

func NewFraction(num, denom int64) Fraction {
	return Fraction{
		Num:   num,
		Denom: denom,
	}
}

func (f Fraction) Less(rhs Fraction) bool {
	l := f.Num * rhs.Denom
	r := rhs.Num * f.Denom
	return l < r
}

func (f Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.Num, f.Denom)
}

// f.num == n1*k, f.denom == d1*k, where k == GCD(num,denom). After reducing, f.num == n1 and f.denom == d1
func (f *Fraction) Reduce() {
	k := GCD(f.Num, f.Denom)
	f.Num /= k
	f.Denom /= k
}
