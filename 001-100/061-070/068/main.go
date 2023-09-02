package main

import (
	"log"
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 68; Magic 5-gon Ring
Consider the following "magic" 3-gon ring, filled with the numbers 1 to 6, and each line adding to nine.

Working clockwise, and starting from the group of three with the numerically lowest external node (4,3,2 in this example),
each solution can be described uniquely. For example, the above solution can be described by the set: 4,3,2; 6,2,1; 5,1,3.

It is possible to complete the ring with four different totals: 9, 10, 11, and 12. There are eight solutions in total.
Total	Solution Set
9	4,2,3; 5,3,1; 6,1,2
9	4,3,2; 6,2,1; 5,1,3
10	2,3,5; 4,5,1; 6,1,3
10	2,5,3; 6,3,1; 4,1,5
11	1,4,6; 3,6,2; 5,2,4
11	1,6,4; 5,4,2; 3,2,6
12	1,5,6; 2,6,4; 3,4,5
12	1,6,5; 3,5,4; 2,4,6

By concatenating each group it is possible to form 9-digit strings; the maximum string for a 3-gon ring is 432621513.

Using the numbers 1 to 10, and depending on arrangements, it is possible to form 16- and 17-digit strings.
What is the maximum 16-digit string for a "magic" 5-gon ring?
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
		limit = 5
	}

	projecteuler.Timed(calc, limit)
}

type Triplet struct {
	A, B, C    int
	as, bs, cs string
	solution   string
}

func NewTriplet(a, b, c int) Triplet {
	return Triplet{
		A:  a,
		B:  b,
		C:  c,
		as: strconv.Itoa(a),
		bs: strconv.Itoa(b),
		cs: strconv.Itoa(c),
	}
}

func (t Triplet) String() string {
	return t.as + t.bs + t.cs
}

type TripletSlice []Triplet

func (a TripletSlice) Len() int           { return len(a) }
func (a TripletSlice) Less(i, j int) bool { return a[i].A < a[j].A }
func (a TripletSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ngon struct {
	vertices []int
	solution TripletSlice
}

func newNgon(vertices []int) ngon {
	return ngon{
		vertices: vertices,
	}
}

func (ng *ngon) solve() {
	n := len(ng.vertices)
	lineLow := (5*n + 3) / 2
	lineHigh := lineLow + n

	for line := lineLow; line <= lineHigh; line++ {
		allSet := map[int]struct{}{}
		for i := 1; i <= 2*n; i++ {
			allSet[i] = struct{}{}
		}

		vSet := map[int]struct{}{}
		for _, v := range ng.vertices {
			vSet[v] = struct{}{}
			delete(allSet, v)
		}

		ts := make([]Triplet, 0, n)
		for j := 0; j < n; j++ {
			secondIndex := j + 1
			if secondIndex >= n {
				secondIndex -= n
			}
			x := line - ng.vertices[j] - ng.vertices[secondIndex]

			if _, present := allSet[x]; !present {
				continue
			}
			delete(allSet, x)

			ts = append(ts, NewTriplet(x, ng.vertices[j], ng.vertices[secondIndex]))
		}

		if len(ts) != n {
			continue
		}

		b := make([]Triplet, n)
		copy(b, ts)
		sort.Sort(TripletSlice(b))

		minSlice := b[0]
		minSliceIndex := 0

		for i := 0; i < n; i++ {
			if ts[i] == minSlice {
				break
			}

			minSliceIndex++
		}

		ng.solution = append(ts[minSliceIndex:], ts[:minSliceIndex]...)
	}
}

func (ng ngon) hasSolution() bool {
	return ng.solution != nil
}

func (ng ngon) String() string {
	var sb strings.Builder
	for _, t := range ng.solution {
		sb.WriteString(t.String())
	}

	return sb.String()
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	combinations, _ := projecteuler.Combinations(byte(2*limit), byte(limit), nil)

	ngons := []*ngon{}
	for _, c := range combinations {
		ngons = append(ngons, ngonsFromCombination(c)...)
	}

	var wg sync.WaitGroup
	threadPerCoreGuard := make(chan int, runtime.GOMAXPROCS(runtime.NumCPU()))
	m := projecteuler.NewConcurrentStringSet()
	for i := range ngons {
		threadPerCoreGuard <- 1
		wg.Add(1)

		go func(ng *ngon) {
			defer wg.Done()

			<-threadPerCoreGuard
			ng.solve()
			if ng.hasSolution() {
				m.Write(ng.String())
			}

		}(ngons[i])
	}
	wg.Wait()

	solutions := m.Read()
	solutionLengthMap := map[int][]string{}
	for _, s := range solutions {
		solutionLengthMap[len(s)] = append(solutionLengthMap[len(s)], s)
	}

	iMin := math.MaxInt
	for i := range solutionLengthMap {
		if i < iMin {
			iMin = i
		}
	}

	lengthSolutions := solutionLengthMap[iMin]
	sort.Strings(lengthSolutions)
	result = lengthSolutions[len(lengthSolutions)-1]

	return
}

func ngonsFromCombination(c []byte) []*ngon {
	n := len(c) / 2
	populated := make([]byte, 0, n)
	for i, x := range c {
		if x == 1 {
			populated = append(populated, byte(i))
		}
	}

	var ngons []*ngon
	permutations := projecteuler.Permutations(byte(n-1), nil)
	for _, p := range permutations {
		verticeBytes := []byte{populated[0]}
		for i := range p {
			verticeBytes = append(verticeBytes, populated[p[i]+1])
		}

		vertices := make([]int, 0, n)
		for _, x := range verticeBytes {
			vertices = append(vertices, int(x)+1)
		}
		ng := newNgon(vertices)

		ngons = append(ngons, &ng)
	}

	return ngons
}
