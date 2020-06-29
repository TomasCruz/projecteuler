package main

import "sort"

// ranks
const (
	onePair = iota + 1
	twoPairs
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
	royalFlush
)

// weights
const (
	rank       = 3200000
	highValue1 = 160000
	highValue2 = 8000
	highValue3 = 400
	highValue4 = 20
	highValue5 = 1
)

func rankPlayerHand(hand []int, valueChannel chan int) {
	sort.Ints(hand)

	values := make([]int, 5)
	suits := make([]int, 5)
	for i := 0; i < 5; i++ {
		values[i] = hand[i] / 10
		suits[i] = hand[i] % 10
	}

	sameSuit := true
	suitZero := suits[0]
	for i := 1; i < 5; i++ {
		sameSuit = sameSuit && suits[i] == suitZero
		if !sameSuit {
			break
		}
	}

	consecutiveValues := true
	for i := 1; i < 5; i++ {
		consecutiveValues = consecutiveValues && values[i] == values[i-1]+1
		if !consecutiveValues {
			break
		}
	}

	handValue := 0
	rankFound := false

	if consecutiveValues {
		if sameSuit {
			if values[4] == ace {
				handValue = rank * royalFlush
			} else {
				handValue = rank*straightFlush + highValue1*values[4]
			}
		} else {
			handValue = rank*straight + highValue1*values[4]
		}
		rankFound = true
	} else {
		if sameSuit {
			handValue = rank * flush
			handValue += highValue1 * values[4]
			handValue += highValue2 * values[3]
			handValue += highValue3 * values[2]
			handValue += highValue4 * values[1]
			handValue += highValue5 * values[0]
			rankFound = true
		}
	}

	if !rankFound {
		// must be one of those strange cases :)
		valueChannel <- duplicateBasedValue(values)
		return
	}

	valueChannel <- handValue
}

func duplicateBasedValue(values []int) (handValue int) {
	occ := make(map[int]int)
	for _, x := range values {
		occ[x]++
	}

	looseValues := make([]int, 0)
	pairValues := make([]int, 0)

	switch len(occ) {
	case 2:
		fourRank := false
		for v, o := range occ {
			switch o {
			case 1:
				handValue += v * highValue2
			case 2:
				handValue += v * highValue2
			case 3:
				handValue += v * highValue1
			case 4:
				fourRank = true
				handValue += v * highValue1
			}
		}

		if fourRank {
			handValue += rank * fourOfAKind
		} else {
			handValue += rank * fullHouse
		}
	case 3:
		threeRank := false
		for v, o := range occ {
			switch o {
			case 1:
				looseValues = append(looseValues, v)
			case 2:
				pairValues = append(pairValues, v)
			case 3:
				threeRank = true
				handValue += v * highValue1
			}
		}

		sort.Ints(looseValues)
		sort.Ints(pairValues)

		if threeRank {
			handValue += rank * threeOfAKind
			handValue += looseValues[1]*highValue2 + looseValues[0]*highValue3
		} else {
			handValue += rank * twoPairs
			handValue += pairValues[1]*highValue1 + pairValues[0]*highValue2 + looseValues[0]*highValue3
		}
	case 4:
		handValue += rank * onePair
		for v, o := range occ {
			switch o {
			case 1:
				looseValues = append(looseValues, v)
			case 2:
				handValue += v * highValue1
			}
		}

		sort.Ints(looseValues)
		handValue += looseValues[2]*highValue2 + looseValues[1]*highValue3 + looseValues[0]*highValue4
	case 5:
		handValue += values[4]*highValue1 + values[3]*highValue2 + values[2]*highValue3 +
			values[1]*highValue4 + values[0]*highValue5
	}

	return
}
