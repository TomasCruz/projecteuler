package main

import (
	"fmt"
	"math/rand"
)

type fld struct {
	ord  int
	hits int
	name string
}

func (f fld) toString(hitSum int) string {
	return fmt.Sprintf("%02d %.6f %s", f.ord, 100.0*float64(f.hits)/float64(hitSum), f.name)
}

type board struct {
	fields []fld
	decks  decks
}

func newBoard() board {
	fields := []fld{
		// GO,A1,CC1,A2,T1,R1,B1,CH1,B2,B3,
		{ord: 0, hits: 0, name: "GO"},
		{ord: 1, hits: 0, name: "A1"},
		{ord: 2, hits: 0, name: "CC1"},
		{ord: 3, hits: 0, name: "A2"},
		{ord: 4, hits: 0, name: "T1"},
		{ord: 5, hits: 0, name: "R1"},
		{ord: 6, hits: 0, name: "B1"},
		{ord: 7, hits: 0, name: "CH1"},
		{ord: 8, hits: 0, name: "B2"},
		{ord: 9, hits: 0, name: "B3"},

		// JAIL,C1,U1,C2,C3,R2,D1,CC2,D2,D3,
		{ord: 10, hits: 0, name: "JAIL"},
		{ord: 11, hits: 0, name: "C1"},
		{ord: 12, hits: 0, name: "U1"},
		{ord: 13, hits: 0, name: "C2"},
		{ord: 14, hits: 0, name: "C3"},
		{ord: 15, hits: 0, name: "R2"},
		{ord: 16, hits: 0, name: "D1"},
		{ord: 17, hits: 0, name: "CC2"},
		{ord: 18, hits: 0, name: "D2"},
		{ord: 19, hits: 0, name: "D3"},

		// FP,E1,CH2,E2,E3,R3,F1,F2,U2,F3,
		{ord: 20, hits: 0, name: "FP"},
		{ord: 21, hits: 0, name: "E1"},
		{ord: 22, hits: 0, name: "CH2"},
		{ord: 23, hits: 0, name: "E2"},
		{ord: 24, hits: 0, name: "E3"},
		{ord: 25, hits: 0, name: "R3"},
		{ord: 26, hits: 0, name: "F1"},
		{ord: 27, hits: 0, name: "F2"},
		{ord: 28, hits: 0, name: "U2"},
		{ord: 29, hits: 0, name: "F3"},

		// G2J,G1,G2,CC3,G3,R4,CH3,H1,T2,H2
		{ord: 30, hits: 0, name: "G2J"},
		{ord: 31, hits: 0, name: "G1"},
		{ord: 32, hits: 0, name: "G2"},
		{ord: 33, hits: 0, name: "CC3"},
		{ord: 34, hits: 0, name: "G3"},
		{ord: 35, hits: 0, name: "R4"},
		{ord: 36, hits: 0, name: "CH3"},
		{ord: 37, hits: 0, name: "H1"},
		{ord: 38, hits: 0, name: "T2"},
		{ord: 39, hits: 0, name: "H2"},
	}

	decks := mixDecks()

	return board{
		fields: fields,
		decks:  decks,
	}
}

func (b board) makeMove(limit, fromBoardIndex, doubles int) int {
	dice1 := rand.Intn(limit) + 1
	dice2 := rand.Intn(limit) + 1

	if doubles == 2 && dice1 == dice2 {
		b.fields[10].hits++
		return 10
	}

	newBoardIndex := normalizeIndex(fromBoardIndex + dice1 + dice2)
	switch newBoardIndex {
	case 30:
		newBoardIndex = 10
	case 2, 17, 33:
		{
			// isCommunityChest
			cardIndex := b.decks.takeCard(b.decks.communityChest)
			newBoardIndex = b.decks.moveTo(false, cardIndex, newBoardIndex)
		}
	case 7, 22, 36:
		{
			// isChance
			cardIndex := b.decks.takeCard(b.decks.chance)
			newBoardIndex = b.decks.moveTo(true, cardIndex, newBoardIndex)
		}
	}

	b.fields[newBoardIndex].hits++

	if dice1 == dice2 {
		return b.makeMove(limit, newBoardIndex, doubles+1)
	} else {
		return newBoardIndex
	}
}

func (b board) print() {
	hitSum := 0
	for i := 0; i < 40; i++ {
		hitSum += b.fields[i].hits
	}

	for i := 0; i < 40; i++ {
		fmt.Println(b.fields[i].toString(hitSum))
	}
}

func normalizeIndex(nextIndex int) int {
	if nextIndex < 0 {
		nextIndex += 40
	} else if nextIndex > 39 {
		nextIndex -= 40
	}

	return nextIndex
}
