package main

import (
	"math/rand"
)

// chance
// 0-5 stuff
// 6.  Advance to GO 0
// 7.  Go to JAIL 10
// 8.  Go to C1 11
// 9.  Go to E3 24
// 10. Go to H2 39
// 11. Go to R1 5
// 12. Go to next R (railway company)
// 13. Go to next R
// 14. Go to next U (utility company)
// 15. Go back 3 squares.

// community chest
// 0-13 stuff
// 14. Advance to GO
// 15. Go to JAIL

type decks struct {
	chance         []int
	communityChest []int
}

func mixDecks() decks {
	chCardSet := make(map[int]struct{})
	ccCardSet := make(map[int]struct{})
	ret := decks{
		chance:         make([]int, 16),
		communityChest: make([]int, 16),
	}

	var card int
	for i := 0; i < 16; i++ {
		for {
			card = rand.Intn(16)
			if _, present := chCardSet[card]; !present {
				chCardSet[card] = struct{}{}
				ret.chance[i] = card
				break
			}
		}
	}

	for i := 0; i < 16; i++ {
		for {
			card = rand.Intn(16)
			if _, present := ccCardSet[card]; !present {
				ccCardSet[card] = struct{}{}
				ret.communityChest[i] = card
				break
			}
		}
	}

	return ret
}

func (d decks) takeCard(currDeck []int) int {
	card := currDeck[0]

	for i := 0; i < 15; i++ {
		currDeck[i] = currDeck[i+1]
	}
	currDeck[15] = card

	return card
}

func (d decks) moveTo(isChance bool, cardIndex, boardIndex int) int {
	if isChance {
		switch cardIndex {
		case 6:
			return 0
		case 7:
			return 10
		case 8:
			return 11
		case 9:
			return 24
		case 10:
			return 39
		case 11:
			return 5
		case 12, 13: // next R
			switch boardIndex {
			case 7:
				return 15
			case 22:
				return 25
			case 36:
				return 5
			}
		case 14: // next U
			switch boardIndex {
			case 7:
				return 12
			case 22:
				return 28
			case 36:
				return 12
			}
		case 15:
			switch boardIndex {
			case 7:
				return 4
			case 22:
				return 19
			case 36:
				ccIndex := d.takeCard(d.communityChest)
				return d.moveTo(false, ccIndex, 33)
			}
		}
	} else {
		switch cardIndex {
		case 14:
			return 0
		case 15:
			return 10
		}
	}

	return boardIndex
}
