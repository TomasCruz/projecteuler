package main

import "strings"

// values
const (
	jack = iota + 11
	queen
	king
	ace
)

func handToIntSlice(input string) (hand []int) {
	cardStrings := strings.Split(input, " ")
	hand = make([]int, 10)

	i := 0
	for _, cs := range cardStrings {
		switch cs[0] {
		case 'A':
			hand[i] = ace
		case 'K':
			hand[i] = king
		case 'Q':
			hand[i] = queen
		case 'J':
			hand[i] = jack
		case 'T':
			hand[i] = 10
		default:
			hand[i] = int(cs[0] - '0')
		}
		hand[i] *= 10

		switch cs[1] {
		case 'C':
			hand[i]++
		case 'S':
			hand[i] += 2
		case 'D':
			hand[i] += 3
		}

		i++
	}

	return
}
