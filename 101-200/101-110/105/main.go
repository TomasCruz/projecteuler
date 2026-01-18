package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 105; Special Subset Sums: Testing
Let S(A) represent the sum of elements in set A of size n. We shall call it a special sum set if for any two non-empty disjoint subsets,
B and C, the following properties are true:
	S(B) != S(C); that is, sums of subsets cannot be equal.
	If B contains more elements than C then S(B) > S(C).

For example, {81, 88, 75, 42, 87, 84, 86, 65} is not a special sum set because 65 + 87 + 88 = 75 + 81 + 84,
whereas {157, 150, 164, 119, 79, 159, 161, 139, 158} satisfies both rules for all possible subset pair combinations and S(A) = 1286.

Using sets.txt (right click and "Save Link/Target As..."), a 4K text file with one-hundred sets containing seven to twelve elements
(the two examples given above are the first two sets in the file), identify all the special sum sets, A_1, A_2, ..., A_k,
and find the value of S(A_1) + S(A_2) + ... + S(A_k).

NOTE: This problem is related to Problem 103 and Problem 106.
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
		limit = 1
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var fileName string
	if limit == 2 {
		fileName = "example.txt"
	} else {
		fileName = "0105_sets.txt"
	}

	var rowStrings []string
	if rowStrings, err = projecteuler.FileToStrings(fileName); err != nil {
		return
	}

	sum := 0
	for _, r := range rowStrings {
		candidate := row2Slice(r)
		if sss(candidate) {
			sum += sumSet(candidate)
		}
	}

	result = strconv.Itoa(sum)
	return
}

func row2Slice(rowString string) []int {
	var ret []int

	elemStrings := strings.Split(rowString, ",")
	for _, elStr := range elemStrings {
		el, err := strconv.Atoi(elStr)
		if err != nil {
			panic("NaN!")
		}
		ret = append(ret, el)
	}

	sort.Ints(ret)
	return ret
}

func sss(a []int) bool {
	odd := false
	if len(a)%2 == 1 {
		odd = true
	}

	half := len(a) / 2
	if odd {
		half++
	}

	// second rule
	for i := 2; i <= half; i++ {
		if sumSet(a[0:i]) <= sumSet(a[len(a)+1-i:]) {
			return false
		}
	}

	// first rule
	if odd {
		half--
	}
	for i := 2; i <= half; i++ {
		if !rule1(a, i) {
			return false
		}
	}

	return true
}

func rule1(a []int, ssl int) bool {
	subSetMap := make(map[int]struct{})
	subs := [][]int{}
	currSub := []int{}

	subs = extractSubs(a, ssl, currSub, subs)
	for _, cs := range subs {
		currSum := sumSet(cs)
		if _, present := subSetMap[currSum]; present {
			return false
		}
		subSetMap[currSum] = struct{}{}
	}

	return true
}

func extractSubs(a []int, ssl int, currSub []int, subs [][]int) [][]int {
	if ssl == 0 {
		dst := make([]int, len(currSub))
		copy(dst, currSub)
		subs = append(subs, dst)
		return subs
	}

	for i := 0; i < len(a)+1-ssl; i++ {
		currSub = append(currSub, a[i])
		subs = extractSubs(a[i+1:], ssl-1, currSub, subs)
		currSub = currSub[:len(currSub)-1]
	}

	return subs
}

func sumSet(a []int) int {
	sum := 0

	for _, x := range a {
		sum += x
	}

	return sum
}
