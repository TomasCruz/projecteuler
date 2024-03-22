package main

import (
	"github.com/TomasCruz/projecteuler"
)

/*
Problem 79; Passcode Derivation


A common security method used for online banking is to ask the user for three random characters from a passcode. For example, if the passcode was 531278,
they may ask for the 2nd, 3rd, and 5th characters; the expected reply would be: 317.

The text file, keylog.txt, contains fifty successful login attempts. Given that the three characters are always asked for in order,
analyse the file so as to determine the shortest possible secret passcode of unknown length.
*/

func main() {
	projecteuler.Timed(calc)
}

func calc(args ...interface{}) (result string, err error) {
	var numStrings []string
	if numStrings, err = projecteuler.FileToStrings("0079_keylog.txt"); err != nil {
		return
	}

	numStringSet := map[string]struct{}{}
	for i := 0; i < len(numStrings); i++ {
		numStringSet[numStrings[i]] = struct{}{}
	}

	followers := map[byte]map[byte]struct{}{}
	for s := range numStringSet {
		d0 := byte(s[0])
		d1 := byte(s[1])
		d2 := byte(s[2])

		if _, present := followers[d0]; !present {
			followers[d0] = map[byte]struct{}{}
		}
		followers[d0][d1] = struct{}{}

		if _, present := followers[d1]; !present {
			followers[d1] = map[byte]struct{}{}
		}
		followers[d1][d2] = struct{}{}

		if _, present := followers[d2]; !present {
			followers[d2] = map[byte]struct{}{}
		}
	}

	l := len(followers)
	mask := make([]byte, 0, l)
	for k := range followers {
		mask = append(mask, k)
	}

	result = ""
	projecteuler.Permutations(byte(l), checkChars, mask, followers, &result)

	return
}

func checkChars(args ...interface{}) bool {
	mask := args[0].([]byte)
	followers := args[1].(map[byte]map[byte]struct{})
	result := args[2].(*string)
	currPermBytes := args[3].([]byte)

	l := len(followers)
	currPerm := make([]byte, l)
	for i := 0; i < l; i++ {
		currPerm[i] = mask[int(currPermBytes[i])]
	}

	for i := 0; i < l-1; i++ {
		if _, present := followers[currPerm[i]][currPerm[i+1]]; !present {
			return false
		}
	}

	currPermStr := string(currPerm)
	*result = currPermStr

	return true
}
