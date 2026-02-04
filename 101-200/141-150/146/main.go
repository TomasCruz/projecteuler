package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 146; Investigating a Prime Pattern
The smallest positive integer n for which the numbers n^2 + 1, n^2 + 3, n^2 + 7, n^2 + 9, n^2 + 13, and n^2 + 27 are consecutive primes is 10.
The sum of all such integers n below one-million is 1242490.

What is the sum of all such integers n below 150 million?
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
		limit = 150000000
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := uint64(args[0].(int))

	setNs := []uint64{}
	primes := projecteuler.PrimesEratosthenes(limit, nil)
	primeSet := constructprimeSet(primes)
	belowPrimeIndex := 8
	mMaps := constructModuloMaps(belowPrimeIndex, primes)
	prod, tails := constructTails(limit*limit+28, belowPrimeIndex, primes, mMaps)

	n := uint64(0)
	for i := uint64(0); n < limit; i += prod {
		for _, t := range tails {
			m := t + i

			n = m * 10
			if n >= limit {
				break
			}
			sq := n * n

			if arePrimes(n, sq, primes) {
				// considering n^2 mod 2 = 0, n^2 mod 3 = 1, n^2 mod 5 = 0, n^2 mod 7 = 2:
				// n^2 + {even} mod 2 = 0
				// n^2 + 17 mod 3 = 0, n^2 + 23 mod 3 = 0
				// n^2 + 15 mod 5 = 0, n^2 + 25 mod 5 = 0
				// n^2 + 19 mod 7 = 0
				// so only n^2 + 21 need to be checked for primality

				if !isPrime(sq+21, limit, primes, primeSet) {
					setNs = append(setNs, n)
				}
			}
		}
	}

	sum := uint64(0)
	for _, n := range setNs {
		sum += n
	}

	result = strconv.FormatUint(sum, 10)
	return
}

func arePrimes(n, sq uint64, primes []uint64) bool {
	i := 3
	for ; primes[i] < n; i++ {
		// for x = n*n + d, d in [1,3,7,9,13,27], p|x
		// (n*n + d) mod p = 0, n*n mod p = p-d, which shouldn't be in [1,3,7,9,13,27]

		modulo := sq % primes[i]
		ds := []uint64{1, 3, 7, 9, 13, 27}
		j := 0

		for ; j < 6; j++ {
			np := primes[i]
			if i < 9 {
				for np < ds[j] {
					np += primes[i]
				}
			}
			ds[j] = np - ds[j]

			if modulo == ds[j] {
				break
			}
		}

		if j < 6 {
			break
		}
	}

	return primes[i] >= n
}

func isPrime(x, limit uint64, primes []uint64, primeSet map[uint64]struct{}) bool {
	if x < limit {
		_, present := primeSet[x]
		return present
	}

	for i := 3; primes[i]*primes[i] <= x; i++ {
		if x%primes[i] == 0 {
			return false
		}
	}

	return true
}

func constructprimeSet(primes []uint64) map[uint64]struct{} {
	ret := map[uint64]struct{}{}
	for _, p := range primes {
		ret[p] = struct{}{}
	}
	return ret
}

/*
if n^2 is odd, all the numbers that should be consecutive primes would be even. More formally,
n^2 mod 2 = a, n^2+1,n^2+3,n^2+7,n^2+9,n^2+13,n^2+27 (SET) mod 2 = a+1,a+1,a+1,a+1,a+1,a+1 => a+1 != 0 => a=0
n^2 mod 5 = a, SET mod 5 = a+1,a+3,a+2,a+4,a+3,a+2 !=0 => a=0

Since n^2 is divisible by 2 and 5, n^2 is divisible by 10, which also means n is divisible by 10.

n = 10 * m

2:
mods m   [0,1]
mods m^2 [0,1]
mods n^2 [0,0]
allowed n^2 [0]
100*m^2 =(2) 0*m^2 [0,0] in [0], m in [0,1]

3:
mods m   [0,1,2]
mods m^2 [0,1,1]
mods n^2 [0,1,1]
allowed n^2 [1]
100*m^2 =(3) 1*m^2 [0,1,1] in [1], m in [1,2]

5:
mods m   [0,1,2,3,4]
mods m^2 [0,1,4,4,1]
mods n^2 [0,0,0,0,0]
allowed n^2 [0]
100*m^2 =(5) 0*m^2 [0,0,0,0,0] in [0], m in [0,1,2,3,4]

7:
mods m [0,1,2,3,4,5,6]
mods m^2 [0,1,4,2,2,4,1]
mods n^2 [0,2,1,4,4,1,2]
allowed n^2 [2]
100*m^2 =(7) 2*m^2 [0,2,1,4,4,1,2] in [2], m in [1,6]

11:
mods m   [0,1,2,3,4,5,6,7,8,9,10]
mods m^2 [0,1,4,9,5,3,3,5,9,4,1]
mods n^2 [0,1,4,9,5,3,3,5,9,4,1]
allowed n^2 [0,1,5,3]
100*m^2 =(11) 1*m^2 [0,1,4,9,5,3,3,5,9,4,1] in [0,1,5,3], m in [0,1,4,5,6,7,10]
*/
func constructModuloMaps(belowPrimeIndex int, primes []uint64) []map[uint64]struct{} {
	ret := make([]map[uint64]struct{}, belowPrimeIndex)

	for i := 3; i < belowPrimeIndex; i++ {
		mm := make([]uint64, primes[i])
		nn := make([]uint64, primes[i])

		hun := 100 % primes[i]
		for j := range primes[i] {
			mm[j] = (j * j) % primes[i]
			nn[j] = (j * j * hun) % primes[i]
		}

		res := map[uint64]struct{}{}
		for index, x := range nn {
			if (x+1)%primes[i] == 0 || (x+3)%primes[i] == 0 || (x+7)%primes[i] == 0 ||
				(x+9)%primes[i] == 0 || (x+13)%primes[i] == 0 || (x+27)%primes[i] == 0 {
				continue
			}
			res[uint64(index)] = struct{}{}
		}

		ret[i] = res
	}

	return ret
}

/*
1,22,29,34,43,50,55,62,71,76

77 // belowPrimeIndex == 5

2  [0,1]
3  [1,2]
5  [0,1,2,3,4]
7  [1,6]
11 [0,1,4,5,6,7,10]
13 [1,3,4,9,10,12]
17 [0,1,2,4,5,6,11,12,13,15,16]

11: 0,1,4,5,6,7,10, 11,12,15,16,17,18,21, 22,23,26,27,28,29,32, 33,34,37,38,39,40,43, 44,45,48,49,50,51,54, 55,56,59,60,61,62,65, 66,67,70,71,72,73,76
7
1,6,15,22,27,29,34,43,48,50,55,62,71,76
5
1,6,15,22,27,29,34,43,48,50,55,62,71,76
3
1,22,29,34,43,50,55,62,71,76
2
1,22,29,34,43,50,55,62,71,76
2310
0,1,4,5,6,7,10
1,6 15,22 27,29 34,43 48,50 55,62 71,76
1,6,15,22,27,29,34,43,48,50,55,62,71,76, 78,83,92,99,104,106,111,120,125,132,139,148,153 155,160,169,...384
1,22,29,34,43,50,55,62,71,76,83,92,104,106,125,139,148,153,155,160,169...1149
*/
func constructTails(tailLimit uint64, belowPrimeIndex int, primes []uint64, mMaps []map[uint64]struct{}) (uint64, []uint64) {
	prod := uint64(3)
	for i := 3; i < belowPrimeIndex; i++ {
		prod *= primes[i]
	}

	loopLimit := tailLimit
	if loopLimit > prod {
		loopLimit = prod
	}

	tails := []uint64{}

	for i := uint64(1); i < loopLimit; i = nextRoot(i) {
		if i%3 == 0 {
			continue
		}

		j := 4
		for ; j < belowPrimeIndex; j++ {
			if _, present := mMaps[j][i%primes[j]]; !present {
				break
			}
		}

		if j == belowPrimeIndex {
			tails = append(tails, i)
		}
	}

	return prod, tails
}

func nextRoot(i uint64) uint64 {
	if i%7 == 1 {
		i += 5
	} else {
		i += 2
	}
	return i
}
