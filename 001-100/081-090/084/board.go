package main

import "fmt"

type fld struct {
	ord     int
	name    string
	prob    frac
	doubles int
}

func (f fld) toString() string {
	return fmt.Sprintf("%d %s, %d %s %6f", f.ord, f.name, f.doubles, f.prob.toString(), float64(f.prob.num)/float64(f.prob.denom))
}

var boardFields = []fld{
	// GO,A1,CC1,A2,T1,R1,B1,CH1,B2,B3,
	{ord: 0, name: "GO", prob: frac{num: 0, denom: 1}},
	{ord: 1, name: "A1", prob: frac{num: 0, denom: 1}},
	{ord: 2, name: "CC1", prob: frac{num: 0, denom: 1}},
	{ord: 3, name: "A2", prob: frac{num: 0, denom: 1}},
	{ord: 4, name: "T1", prob: frac{num: 0, denom: 1}},
	{ord: 5, name: "R1", prob: frac{num: 0, denom: 1}},
	{ord: 6, name: "B1", prob: frac{num: 0, denom: 1}},
	{ord: 7, name: "CH1", prob: frac{num: 0, denom: 1}},
	{ord: 8, name: "B2", prob: frac{num: 0, denom: 1}},
	{ord: 9, name: "B3", prob: frac{num: 0, denom: 1}},

	// JAIL,C1,U1,C2,C3,R2,D1,CC2,D2,D3,
	{ord: 10, name: "JAIL", prob: frac{num: 0, denom: 1}},
	{ord: 11, name: "C1", prob: frac{num: 0, denom: 1}},
	{ord: 12, name: "U1", prob: frac{num: 0, denom: 1}},
	{ord: 13, name: "C2", prob: frac{num: 0, denom: 1}},
	{ord: 14, name: "C3", prob: frac{num: 0, denom: 1}},
	{ord: 15, name: "R2", prob: frac{num: 0, denom: 1}},
	{ord: 16, name: "D1", prob: frac{num: 0, denom: 1}},
	{ord: 17, name: "CC2", prob: frac{num: 0, denom: 1}},
	{ord: 18, name: "D2", prob: frac{num: 0, denom: 1}},
	{ord: 19, name: "D3", prob: frac{num: 0, denom: 1}},

	// FP,E1,CH2,E2,E3,R3,F1,F2,U2,F3,
	{ord: 20, name: "FP", prob: frac{num: 0, denom: 1}},
	{ord: 21, name: "E1", prob: frac{num: 0, denom: 1}},
	{ord: 22, name: "CH2", prob: frac{num: 0, denom: 1}},
	{ord: 23, name: "E2", prob: frac{num: 0, denom: 1}},
	{ord: 24, name: "E3", prob: frac{num: 0, denom: 1}},
	{ord: 25, name: "R3", prob: frac{num: 0, denom: 1}},
	{ord: 26, name: "F1", prob: frac{num: 0, denom: 1}},
	{ord: 27, name: "F2", prob: frac{num: 0, denom: 1}},
	{ord: 28, name: "U2", prob: frac{num: 0, denom: 1}},
	{ord: 29, name: "F3", prob: frac{num: 0, denom: 1}},

	// G2J,G1,G2,CC3,G3,R4,CH3,H1,T2,H2
	{ord: 30, name: "G2J", prob: frac{num: 0, denom: 1}},
	{ord: 31, name: "G1", prob: frac{num: 0, denom: 1}},
	{ord: 32, name: "G2", prob: frac{num: 0, denom: 1}},
	{ord: 33, name: "CC3", prob: frac{num: 0, denom: 1}},
	{ord: 34, name: "G3", prob: frac{num: 0, denom: 1}},
	{ord: 35, name: "R4", prob: frac{num: 0, denom: 1}},
	{ord: 36, name: "CH3", prob: frac{num: 0, denom: 1}},
	{ord: 37, name: "H1", prob: frac{num: 0, denom: 1}},
	{ord: 38, name: "T2", prob: frac{num: 0, denom: 1}},
	{ord: 39, name: "H2", prob: frac{num: 0, denom: 1}},
}

func isCommunityChest(index int) bool {
	if index == 2 || index == 17 || index == 33 {
		return true
	}

	return false
}

func isChance(index int) bool {
	if index == 7 || index == 22 || index == 36 {
		return true
	}

	return false
}

func isSpecial(index int) bool {
	if index == 30 || isCommunityChest(index) || isChance(index) {
		return true
	}

	return false
}

func cloneBoard() []fld {
	ret := make([]fld, 40)

	for i := 0; i < 40; i++ {
		ret[i] = fld{
			ord:  i,
			name: boardFields[i].name,
			prob: frac{num: 0, denom: 1},
		}
	}

	return ret
}

func printBoard(board []fld) {
	for i := 0; i < 40; i++ {
		fmt.Println(board[i].toString())
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

func calcProbs(limit, startIndex, doubles int, probs []frac, probMatrix [][][]fld) []fld {
	if probMatrix[doubles][startIndex] != nil {
		return probMatrix[doubles][startIndex]
	}

	ret := cloneBoard()

	for i := 3; i < 2*limit; i++ {
		nextIndex := normalizeIndex(startIndex + i)
		nextProb := probs[i]
		if isSpecial(nextIndex) {
			fields := processSpecial(nextIndex)
			for j := range fields {
				if fields[j].prob.num == 0 {
					continue
				}

				ret[j].prob = ret[j].prob.add(nextProb.mul(fields[j].prob))
			}
		} else {
			ret[nextIndex].prob = ret[nextIndex].prob.add(nextProb)
		}
	}

	// process doubles
	if doubles == 2 {
		ret[10].prob = ret[10].prob.add(frac{num: 1, denom: int64(limit)})
		return ret
	}

	limitSquare := limit * limit
	for i := 1; i <= limit; i++ {
		nextIndex := normalizeIndex(startIndex + 2*i)
		nextProb := frac{num: 1, denom: int64(limitSquare)}
		if isSpecial(nextIndex) {
			fields := processSpecial(nextIndex)
			for j := range fields {
				if fields[j].prob.num == 0 {
					continue
				}

				fields2 := calcProbs(limit, j, doubles+1, probs, probMatrix)
				for k := 0; k < 40; k++ {
					if fields2[k].prob.num == 0 {
						continue
					}

					ret[k].prob = ret[k].prob.add(nextProb.mul(fields[j].prob).mul(fields2[k].prob))
				}
			}
		} else {
			fields2 := calcProbs(limit, nextIndex, doubles+1, probs, probMatrix)
			for k := 0; k < 40; k++ {
				if fields2[k].prob.num == 0 {
					continue
				}

				ret[k].prob = ret[k].prob.add(nextProb.mul(fields2[k].prob))
			}
		}
	}

	return ret
}

func processSpecial(index int) []fld {
	if index == 30 {
		ret := cloneBoard()
		ret[10].prob = frac{num: 1, denom: 1}
		return ret
	} else if isCommunityChest(index) {
		return processCommunityChest(index)
	}

	return processChance(index)
}

func processCommunityChest(index int) []fld {
	ret := cloneBoard()

	ret[index].prob = frac{num: 14, denom: 16}
	ret[0].prob = frac{num: 1, denom: 16}
	ret[10].prob = frac{num: 1, denom: 16}

	return ret
}

func processChance(index int) []fld {
	ret := cloneBoard()

	// Advance to GO
	// Go to JAIL
	// Go to C1
	// Go to E3
	// Go to H2
	// Go to R1
	// Go to next R (railway company)
	// Go to next R
	// Go to next U (utility company)
	// Go back 3 squares.

	ret[index].prob = frac{num: 6, denom: 16}
	ret[0].prob = frac{num: 1, denom: 16}
	ret[10].prob = frac{num: 1, denom: 16}
	ret[11].prob = frac{num: 1, denom: 16}
	ret[24].prob = frac{num: 1, denom: 16}
	ret[39].prob = frac{num: 1, denom: 16}
	ret[5].prob = frac{num: 1, denom: 16}

	switch index {
	case 7:
		ret[15].prob = frac{num: 2, denom: 16}
		ret[12].prob = frac{num: 1, denom: 16}
		ret[4].prob = frac{num: 1, denom: 16}
	case 22:
		ret[25].prob = frac{num: 2, denom: 16}
		ret[28].prob = frac{num: 1, denom: 16}
		ret[19].prob = frac{num: 1, denom: 16}
	case 36:
		ret[5].prob = frac{num: 3, denom: 16}
		ret[12].prob = frac{num: 1, denom: 16}

		fields := processCommunityChest(33)
		for i := range fields {
			ret[i].prob = ret[i].prob.add(frac{num: 1, denom: 16}.mul(fields[i].prob))
		}
	}

	return ret
}
