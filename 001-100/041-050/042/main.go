package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TomasCruz/projecteuler"
)

/*
Problem 42; Coded triangle numbers

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values
we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a
triangle number then we shall call the word a triangle word.

Using words.txt, a text file containing nearly two-thousand common English words, how many are triangle words?
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	var wordsLines []string
	if wordsLines, err = projecteuler.FileToStrings("p042_words.txt"); err != nil {
		fmt.Println(err)
		return
	}

	words := strings.Split(wordsLines[0], `","`)
	words[0] = words[0][1:]
	wordCount := len(words)
	words[wordCount-1] = words[wordCount-1][:len(words[wordCount-1])-1]

	// maximum number of letters in a common English word being 50 sounds like a reasonable assumption.
	// So, max value == 26*50 == 1300 => n(n+1)==2600 => n <= 50 (not quite, but should be enough :)

	triangles := buildTriangles(50)
	triangleWordCount := 0

	for _, w := range words {
		wordVal := wordValue(w)
		if _, ok := triangles[wordVal]; ok {
			triangleWordCount++
		}
	}

	result = strconv.Itoa(triangleWordCount)
	return
}

func buildTriangles(limit int) (triangles map[int]struct{}) {
	triangles = make(map[int]struct{})

	for i := 1; i <= limit; i++ {
		triangles[i*(i+1)/2] = struct{}{}
	}

	return
}

func wordValue(word string) int {
	var result int
	for _, c := range word {
		result += int(c - 'A' + 1)
	}

	return result
}
