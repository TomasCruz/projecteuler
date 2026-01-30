package main

type nfcm struct {
	limit, digit    int
	s, m, e, primes []int
	mask            string
}

func newNFCM(limit, digit int, mask string, primes []int) nfcm {
	// s for non-current non-zero digit, m for non-current digit, e for odd, non-current non-5 digits
	s := []int{}
	m := []int{}
	e := []int{}
	for i := 0; i < 10; i++ {
		if i == digit {
			continue
		}

		if i != 0 {
			s = append(s, i)
		}

		if i%2 == 1 && i != 5 {
			e = append(e, i)
		}

		m = append(m, i)
	}

	return nfcm{
		limit:  limit,
		digit:  digit,
		s:      s,
		m:      m,
		e:      e,
		primes: primes,
		mask:   mask,
	}
}

func (n nfcm) genPrimesRec(from int, soFar int64, ncm map[int64]struct{}) {
	if n.limit == from {
		if n.isPrime(soFar) {
			ncm[soFar] = struct{}{}
		}
		return
	}

	soFar *= int64(10)

	switch n.mask[from] {
	case 'd':
		n.genPrimesRec(from+1, soFar+int64(n.digit), ncm)
	case 's':
		for _, i := range n.s {
			n.genPrimesRec(from+1, soFar+int64(i), ncm)
		}
	case 'm':
		for _, i := range n.m {
			n.genPrimesRec(from+1, soFar+int64(i), ncm)
		}
	case 'e':
		for _, i := range n.e {
			n.genPrimesRec(from+1, soFar+int64(i), ncm)
		}
	}
}

func (n nfcm) isPrime(x int64) bool {
	for _, p := range n.primes {
		if x%int64(p) == 0 {
			return false
		}
	}
	return true
}
