package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 98; Anagramic Squares
By replacing each of the letters in the word CARE with 1, 2, 9, and 6 respectively, we form a square number: 1296 = 36^2.
What is remarkable is that, by using the same digital substitutions, the anagram, RACE, also forms a square number: 9216 = 96^2.
We shall call CARE (and RACE) a square anagram word pair and specify further that leading zeroes are not permitted, neither may a
different letter have the same digital value as another letter.

Using "words.txt", a 16K text file containing nearly two-thousand common English words, find all the square anagram word pairs
(a palindromic word is NOT considered to be an anagram of itself).

What is the largest square number formed by any member of such a pair?
NOTE: All anagrams formed must be contained in the given text file.
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
		limit = 10
	}

	projecteuler.Timed(calc, limit)
}

func calc(args ...interface{}) (result string, err error) {
	limit := args[0].(int)

	var fileName string
	if limit == 2 {
		fileName = "example.txt"
	} else {
		fileName = "0098_words.txt"
	}

	var rowStrings []string
	if rowStrings, err = projecteuler.FileToStrings(fileName); err != nil {
		return
	}

	words := strings.Split(rowStrings[0], "\"")
	anagrams := makeAnagrams(words)
	maxAL := len(anagrams) - 1
	squares := makeSquares(maxAL)

	squareAnagramWordSet := map[int]struct{}{}
	for ll := len(anagrams); ll > 0; ll-- {
		l := ll - 1
		if len(anagrams[l]) == 0 {
			continue
		}

		digitCombinations, _ := projecteuler.Combinations(10, byte(l), nil)
		for _, currCombination := range digitCombinations {
			digIndex := 0
			digits := make([]byte, l)
			for c := 0; c < 10; c++ {
				if currCombination[c] == 1 {
					digits[digIndex] = byte(c)
					digIndex++
					if digIndex == l {
						break
					}
				}
			}

			perms := projecteuler.Permutations(byte(l), nil)
			for _, cp := range perms {
				currPerm := make([]byte, l)
				for i := 0; i < l; i++ {
					currPerm[i] = digits[cp[i]]
				}

				for wClass := range anagrams[l] {
					mapping := map[byte]byte{} // letter to digit
					for i := 0; i < l; i++ {
						mapping[wClass[i]] = currPerm[i]
					}

					currSAWSet := map[int]struct{}{}
					for currWord := range anagrams[l][wClass] {
						if mapping[currWord[0]] == 0 ||
							mapping[currWord[l-1]] == 2 || mapping[currWord[l-1]] == 3 ||
							mapping[currWord[l-1]] == 7 || mapping[currWord[l-1]] == 8 {
							break
						}

						x := makeNumber(mapping, currWord)
						if isSquare(x, squares) {
							currSAWSet[x] = struct{}{}
						}
					}

					if len(currSAWSet) > 1 {
						for saw := range currSAWSet {
							squareAnagramWordSet[saw] = struct{}{}
						}
					}
				}
			}
		}

		if len(squareAnagramWordSet) != 0 {
			break
		}
	}

	// fmt.Println("squareAnagramWordSet: ", squareAnagramWordSet)
	resInt := maxElement(squareAnagramWordSet)
	result = strconv.Itoa(resInt)

	return
}

// return value: word length -> word cardinality class -> {word}
func makeAnagrams(words []string) []map[string]map[string]struct{} {
	anagramMap := map[int]map[string]map[string]struct{}{}
	for _, w := range words {
		l := len(w)
		if _, exist := anagramMap[l]; !exist {
			anagramMap[l] = map[string]map[string]struct{}{}
		}
		wClass := sortString(w)
		if _, exist := anagramMap[l][wClass]; !exist {
			anagramMap[l][wClass] = map[string]struct{}{}
		}
		anagramMap[l][wClass][w] = struct{}{}
	}

	deleteLen := []int{}
	deletes := []string{}
	for l := range anagramMap {
		for wClass := range anagramMap[l] {
			if len(anagramMap[l][wClass]) < 2 {
				deleteLen = append(deleteLen, l)
				deletes = append(deletes, wClass)
			}
		}

		for i := 0; i < len(deletes); i++ {
			delete(anagramMap[deleteLen[i]], deletes[i])
		}
	}

	maxAL := 0
	for l := range anagramMap {
		if l > maxAL && len(anagramMap[l]) > 0 {
			maxAL = l
		}
	}

	anagrams := make([]map[string]map[string]struct{}, maxAL+1)
	for l := 2; l <= maxAL; l++ {
		if _, exist := anagramMap[l]; exist {
			anagrams[l] = anagramMap[l]
		}
	}

	return anagrams
}

func isSquare(x int, squares map[int]struct{}) bool {
	_, exists := squares[x]
	return exists
}

func makeSquares(maxLen int) map[int]struct{} {
	power := 1
	for i := 0; i < maxLen; i++ {
		power *= 10
	}

	squares := map[int]struct{}{}
	for i := 1; ; i++ {
		sq := i * i
		if sq >= power {
			break
		}

		squares[sq] = struct{}{}
	}

	return squares
}

func sortString(word string) string {
	s := []rune(word)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func makeNumber(mapping map[byte]byte, word string) int {
	l := len(word)
	num := 0

	for i := 0; i < l; i++ {
		num *= 10
		d := mapping[word[i]]
		num += int(d)
	}

	return num
}

func maxElement(intSet map[int]struct{}) int {
	m := math.MinInt
	for i := range intSet {
		if i > m {
			m = i
		}
	}

	return m
}
