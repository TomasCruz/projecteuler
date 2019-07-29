package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 22; Names scores

Using names.txt, a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order.
Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to
obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is
the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

type nameType []string

func (a nameType) Len() int           { return len(a) }
func (a nameType) Less(i, j int) bool { return a[i] < a[j] }
func (a nameType) Swap(i, j int)      { t := a[i]; a[i] = a[j]; a[j] = t }

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	var textNumbers []string
	if textNumbers, err = projecteuler.FileToStrings("p022_names.txt"); err != nil {
		fmt.Println(err)
		return
	}

	names := strings.Split(textNumbers[0], ",")
	for i := 0; i < len(names); i++ {
		names[i] = names[i][1 : len(names[i])-1]
	}
	sort.Sort(nameType(names))

	totalSum := 0
	for i := 0; i < len(names); i++ {
		totalSum += (i + 1) * alphaSum(names[i])
	}

	result = strconv.Itoa(totalSum)
	return
}

func alphaSum(s string) (sum int) {
	for _, c := range s {
		sum += int(c - 'A' + 1)
	}
	return
}
