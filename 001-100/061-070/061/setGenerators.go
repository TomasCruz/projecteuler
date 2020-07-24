package main

func generateTriangles() figurateMap {
	// Triangle 	  	P(3,n)=n(n+1)/2 	  	1, 3, 6, 10, 15, ...
	// 1000 <= n(n+1)/2 <= 9999, 2000 <= n(n+1) <= 19998
	// n^2 + n - 2000 >= 0
	// [-1+sqrt(1+8000)]/2, 44.2
	// n >= 45

	// n^2 + n - 19998 <= 0
	// [-1+sqrt(1+79992)]/2, 140.9
	// n <= 140

	s := make(figurateMap)

	t := figurateNumber{x: 45 * (45 + 1) / 2, t: figurateTriangle}
	s[t.x] = t
	for i := 46; i <= 140; i++ {
		t.x += i
		s[t.x] = t
	}

	return s
}

func generateSquares() figurateMap {
	// Square 	  		P(4,n)=n^2 	  			1, 4, 9, 16, 25, ...
	// 1000 <= n^2 <= 9999
	// n >= sqrt(1000), n >= 31.6, n >= 32
	// n <= sqrt(9999), n <= 99.9, n <= 99

	s := make(figurateMap)

	for i := 32; i <= 99; i++ {
		n := i * i
		t := figurateNumber{x: n, t: figurateSquare}

		s[t.x] = t
	}

	return s
}

func generatePentagonal() figurateMap {
	// Pentagonal 	  	P(5,n)=n(3n−1)/2 	  	1, 5, 12, 22, 35, ...
	// 1000 <= n(3n−1)/2 <= 9999, 2000 <= n(3n−1) <= 19998
	// 3n^2-n-2000 >= 0
	// [1+sqrt(1+24000)]/6, 25.9
	// n >= 26
	// 3n^2-n-19998 >= 0
	// [1+sqrt(1+239976)]/6, 81.8
	// n <= 81

	s := make(figurateMap)

	for i := 26; i <= 81; i++ {
		n := i * (3*i - 1) / 2
		t := figurateNumber{x: n, t: figuratePentagonal}
		s[t.x] = t
	}

	return s
}

func generateHexagonal() figurateMap {
	// Hexagonal 	  	P(6,n)=n(2n−1) 	  		1, 6, 15, 28, 45, ...
	// 1000 <= n(2n−1) <= 9999
	// 2n^2-n-1000 >= 0
	// [1+sqrt(1+8000)]/4, 22.6
	// n >= 23
	// 2n^2-n-9999 >= 0
	// [1+sqrt(1+79992)]/4, 70.9
	// n <= 70

	s := make(figurateMap)

	for i := 23; i <= 70; i++ {
		n := i * (2*i - 1)
		t := figurateNumber{x: n, t: figurateHexagonal}
		s[t.x] = t
	}

	return s
}

func generateHeptagonal() figurateMap {
	// Heptagonal 	  	P(7,n)=n(5n−3)/2 	  	1, 7, 18, 34, 55, ...
	// 1000 <= n(5n−3)/2 <= 9999, 2000 <= n(5n−3) <= 19998
	// 5n^2-3n-2000 >= 0
	// [3+sqrt(9+40000)]/10, 20.3
	// n >= 21
	// 5n^2-3n-19998 >= 0
	// [3+sqrt(9+399960)]/10, 63.5
	// n <= 63

	s := make(figurateMap)

	for i := 21; i <= 63; i++ {
		n := i * (5*i - 3) / 2
		t := figurateNumber{x: n, t: figurateHeptagonal}
		s[t.x] = t
	}

	return s
}

func generateOctagonal() figurateMap {
	// Octagonal 	  	P(8,n)=n(3n−2) 	  		1, 8, 21, 40, 65, ...
	// 1000 <= n(3n−2) <= 9999
	// 3n^2-2n-1000 >= 0
	// [2+sqrt(4+12000)]/6, 18.5
	// n >= 19
	// 3n^2-2n-9999 >= 0
	// [2+sqrt(4+119988)]/6, 58.1
	// n <= 58

	s := make(figurateMap)

	for i := 19; i <= 58; i++ {
		n := i * (3*i - 2)
		t := figurateNumber{x: n, t: figurateOctagonal}
		s[t.x] = t
	}

	return s
}
