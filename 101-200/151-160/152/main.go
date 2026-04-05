package main

import (
	"log"
	"math"
	"os"
	"slices"
	"strconv"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 152; Sums of Square Reciprocals
There are several ways to write the number 1/2 as a sum of square reciprocals using distinct integers.

For instance, the numbers {2,3,4,5,7,12,15,20,28,35} can be used:

1/2 = 1/2^2 + 1/3^2 + 1/4^2 + 1/5^2 +
	1/7^2 + 1/12^2 + 1/15^2 + 1/20^2 +
	1/28^2 + 1/35^2

In fact, only using integers between 2 and 45 inclusive, there are exactly three ways to do it,
the remaining two being: {2,3,4,6,7,9,10,20,28,35,36,45} and {2,3,4,6,7,9,12,15,28,30,35,36,45}.

How many ways are there to write 1/2 as a sum of reciprocals of squares using distinct integers between 2 and 80 inclusive?
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
		limit = 80
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	primesAll := projecteuler.Primes(limit/2, nil)
	primes := calcRelevantPrimes(limit, primesAll)

	// ai <= log80/log(pi)
	l := map[int]int{}
	for i := range primes {
		l[primes[i]] = int(math.Floor(math.Log(float64(limit)) / math.Log(float64(primes[i]))))
	}

	bigNumber := projecteuler.MultiplyFactors(l)
	bigNumber *= bigNumber

	target := bigNumber / 2

	numbers := make([]int64, limit+1)
	for i := 2; i <= limit; i++ {
		numbers[i] = bigNumber / int64(i*i)
	}

	divisors := projecteuler.FindDivisors(l)
	slices.Sort(divisors)

	i := 0
	for i < len(divisors) && divisors[i] <= limit {
		i++
	}
	divisors = divisors[1:i] // cut off 1 and divisors greater than limit

	upperStartIndex, _ := slices.BinarySearch(divisors, limit/2)
	upperStartIndex -= 3 // fastest execution

	upperDivisors := divisors[upperStartIndex:]
	upperSums := calcUpperSums(upperDivisors, numbers)

	combos := findCombos(divisors, numbers, target, upperStartIndex, upperSums)

	result = strconv.Itoa(len(combos))
	return
}

func calcRelevantPrimes(limit int, primesAll []int) []int {
	ret := []int{2} // 2 must be among numbers in the solution, so 2 must be among relevant primes
	for i := 1; i < len(primesAll); i++ {
		p := primesAll[i]
		if calcRelevantPrimesRec(p, p*p, limit/p, 1, 0, 1) {
			ret = append(ret, p)
		}
	}

	return ret
}

func calcRelevantPrimesRec(p, pSq, pLimit, index, currNom, currDenom int) bool {
	if index > pLimit {
		return false
	}

	if currNom != 0 && currNom%pSq == 0 {
		return true
	}

	if calcRelevantPrimesRec(p, pSq, pLimit, index+1, currNom, currDenom) {
		return true
	}

	// cN/(cD*p)^2 + 1/(index*p)^2 == (cN*index^2 + cD^2) /(cD*index*p)^2
	if calcRelevantPrimesRec(p, pSq, pLimit, index+1, currNom*index*index+currDenom*currDenom, currDenom*index) {
		return true
	}

	return false
}

type bsf struct {
	x  int64
	bs projecteuler.Bitset[int]
}

func cmpInts(a, b int64) int {
	switch {
	case a > b:
		return 1
	case a < b:
		return -1
	default:
		return 0
	}
}

func calcUpperSums(upperDivisors []int, numbers []int64) []bsf {
	pow2 := 1
	for range len(upperDivisors) {
		pow2 *= 2
	}
	ret := make([]bsf, 0, pow2)
	ret = calcUpperSumsRec(upperDivisors, numbers, 0, 0, projecteuler.NewBitset(len(upperDivisors), 64), ret)
	slices.SortFunc(ret, func(a, b bsf) int {
		return cmpInts(a.x, b.x)
	})

	return ret
}

func calcUpperSumsRec(upperDivisors []int, numbers []int64, index int, currSum int64, currBS projecteuler.Bitset[int], ret []bsf) []bsf {
	if index == len(upperDivisors) {
		ret = append(ret, bsf{x: currSum, bs: currBS.Clone()})
		return ret
	}

	ret = calcUpperSumsRec(upperDivisors, numbers, index+1, currSum, currBS, ret)
	currBS.Set(index, true)
	ret = calcUpperSumsRec(upperDivisors, numbers, index+1, currSum+numbers[upperDivisors[index]], currBS, ret)
	currBS.Set(index, false)

	return ret
}

func findCombos(divisors []int, numbers []int64, target int64, upperStartIndex int, upperSums []bsf) [][]int {
	ret := [][]int{}
	ret = findCombosRec(divisors, numbers, upperSums, target, 1, upperStartIndex, numbers[2], []int{2}, ret)

	return ret
}

func findCombosRec(divisors []int, numbers []int64, upperSums []bsf, target int64,
	index, upperStartIndex int, currSum int64, currCombo []int, ret [][]int) [][]int {

	if index == upperStartIndex {
		x := target - currSum
		bsfIndex, found := slices.BinarySearchFunc(upperSums, x, func(a bsf, x int64) int {
			return cmpInts(a.x, x)
		})

		if found {
			indexStart := bsfIndex
			for indexStart > 0 && upperSums[indexStart-1].x == x {
				indexStart--
			}

			indexEnd := bsfIndex
			for indexEnd < len(upperSums)-1 && upperSums[indexEnd+1].x == x {
				indexEnd++
			}

			for i := indexStart; i <= indexEnd; i++ {
				cc := extractSolution(upperSums, i, currCombo, divisors, upperStartIndex)
				ret = append(ret, cc)
			}
			// fmt.Println(cc)
		}

		return ret
	}

	nextSum := currSum + numbers[divisors[index]]

	if nextSum <= target {
		currCombo = append(currCombo, divisors[index])
		ret = findCombosRec(divisors, numbers, upperSums, target, index+1, upperStartIndex, nextSum, currCombo, ret)
		currCombo = currCombo[:len(currCombo)-1]
	}
	ret = findCombosRec(divisors, numbers, upperSums, target, index+1, upperStartIndex, currSum, currCombo, ret)

	return ret
}

func extractSolution(upperSums []bsf, index int, currCombo []int, divisors []int, upperStartIndex int) []int {
	rest := upperSums[index].bs.All()
	rSl := make([]int, 0, len(rest))
	for k := range rest {
		rSl = append(rSl, k)
	}
	slices.Sort(rSl)

	cc := make([]int, len(currCombo)+len(rest))
	i := 0
	for ; i < len(currCombo); i++ {
		cc[i] = currCombo[i]
	}
	for j := range rSl {
		cc[i+j] = divisors[rSl[j]+upperStartIndex]
	}
	return cc
}

// might be useful in the future
func calcTrigamma(numbers []float64) []float64 {
	trigamma := make([]float64, len(numbers)+1)
	trigamma[1] = math.Pi * math.Pi / 6.0
	for i := 2; i < len(trigamma); i++ {
		trigamma[i] = trigamma[i-1] - numbers[i-1]
	}

	return trigamma
}

/*
	trigamma(1) = pi^2/6
	trigamma(2) = pi^2/6 - 1 =(approx.) 0.644934067
	trigamma(3) =(approx.) 0.394934067
	trigamma(4) =(approx.) 0.283822956

	trigamma(n + 1) = trigamma(n) - (1/n^2)
	trigamma(n + 1) = pi^2/6 - 1 - sum[2..n](1 / i^2)

	trigamma(m + 1) - trigamma(n + 1) = pi^2/6 - 1 - sum[2..m](1 / i^2) - (pi^2/6 - 1 - sum[2..n](1 / i^2)) = -sum[n+1..m](1 / i^2)
	trigamma(n) - trigamma(m) = sum[n..m-1](1 / i^2)

	trigamma(n) - trigamma(limit+1) = sum[n..limit](1 / i^2)
	trigamma(2) - trigamma(n+1) = sum[2..n](1 / i^2)

	trigamma(3) - trigamma(n+1) = sum[3..n](1 / i^2) =(approx.) 0.394934067 < 0.5 => 2 must be among xi

	might be useful in the future:
	1/x + 1/{2x^2} < trigamma(x) < 1/x + 1/{2x^2} + 1/{6x^3}
	https://math.stackexchange.com/questions/4185455/upper-bound-on-trigamma-function-psi1x-for-x-in-bbb-n
	https://royalsocietypublishing.org/doi/10.1098/rspa.2017.0363

	........... ANALYSIS ..........

	https://en.wikipedia.org/wiki/Divisor_function

	Sz(n) = sum[d|n](d^z)
	S(-2)(n) = sum[d|n](d^(-2))

	L = P[1..n]xi, prime factors of xi being pi[1..r], powers of pi being ai
		sum[d|L](d^(-2)) = P[1..r](1 - pi^(ai+1)*(-2))/(1 - pi^(-2))

	A = subset d|L such that sum[d|A](d^(-2)) == 1/2 => sum[1..n]1/xi^(-2) == 1/2, multiply by P[1..n]xi^2
		=> 1/2 * P[1..n]xi^2 == sum[1..n](x1^2 * ... xi-1^2 * xi+1^2 * ... * xn^2) (both sides divisible by xi^2 for every i)
		=> for every i, (x1 * ... xi-1 * xi+1 * ... * xn)^2 is divisible by xi^2
		=> for every i, prime factors of xi also exist among factors of other xj (j in [1,n] and j != i)
		=> max pi <= limit/2 to have a multiple of factor of xi among other xj

	for every i, pi^ai <= 80,
	pi^ai <= 80 => log(pi^ai)/log(pi) = ai <= log80/log(pi)

	relevant primes:
	Let x1..xn be one of the solutions. 1/x1^2 + ... + 1/x(n-1)^2 = a/b.
	a/b + 1/xn^2 = (a*xn^2 + b)/b*xn^2 = 1/2
	if b is not divisible by xn^2, there will be factors of xn
	remaining in the denominator after reduction so the reduced fraction will not be 1/2.

	As previously shown, there can't be prime factors of x1..xn appearing only once,
	so there must be at least 2 multiples of a prime p among x1..xn.
	Considering the sum of multiples of p among x1..xn, and the sum of the other members of x1..xn,
	having a/b be the reduced sum of square reciprocals of the other members of x1..xn,
	b must be a divisor of lcm(the other members) meaning it's not divisible by p since the other members are not divisible by p.

	1/2 = a/b + Y/k*p^2 = (a*k*p^2 + b*Y)/(b*k*p^2) => b*Y must be divisible by p^2 => Y must be divisible by p^2.

	Only primes p that need to be considered are such that there are multiples of p such that Y/k*p^2 can be reduced by p^2.

	Divide problem into lower and upper sums.

	Examine for product of relevant prime factors (and powers) within above limits which subset of divisors sums to 1/2.

	P.S. reworked to deal with int64
*/
