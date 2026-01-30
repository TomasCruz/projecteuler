package main

import "strings"

func generateMasks(limit int) [][]map[string]struct{} {
	ret := make([][]map[string]struct{}, 10)

	ret[0] = genZeroMasks(limit)
	oddNonFiveMasks := genOddNonFiveMasks(limit)
	evenPlusFiveMasks := genEvenPlusFiveMasks(limit)

	for i := 1; i < 10; i += 2 {
		if i == 5 {
			continue
		}
		ret[i] = oddNonFiveMasks
	}
	for i := 2; i < 10; i += 2 {
		ret[i] = evenPlusFiveMasks
	}
	ret[5] = evenPlusFiveMasks

	return ret
}

func genZeroMasks(limit int) []map[string]struct{} {
	innerMasks := genInnerMasks(1, limit-3)

	finishedMasks := map[string]struct{}{}
	for s := range innerMasks {
		sb := strings.Builder{}
		sb.WriteRune('s')
		sb.WriteString(s)
		sb.WriteRune('e')
		finishedMasks[sb.String()] = struct{}{}
	}

	return orderMasks(limit, finishedMasks)
}

func genOddNonFiveMasks(limit int) []map[string]struct{} {
	innerMasks := genInnerMasks(0, limit-2)

	finishedMasks := map[string]struct{}{}
	for s := range innerMasks {
		x := s

		if s[0] == 'm' {
			x = "s" + s[1:]
		}

		if s[len(s)-1] == 'm' {
			x = x[:len(s)-1] + "e"
		}

		finishedMasks[x] = struct{}{}
	}

	return orderMasks(limit, finishedMasks)
}

func genEvenPlusFiveMasks(limit int) []map[string]struct{} {
	innerMasks := genInnerMasks(0, limit-3)

	finishedMasks := map[string]struct{}{}
	for s := range innerMasks {
		x := s

		if s[0] == 'm' {
			x = "s" + s[1:]
		}

		x = x + "e"
		finishedMasks[x] = struct{}{}
	}

	return orderMasks(limit, finishedMasks)
}

func orderMasks(limit int, masks map[string]struct{}) []map[string]struct{} {
	ret := make([]map[string]struct{}, limit)
	for i := 2; i < limit; i++ {
		ret[i] = map[string]struct{}{}
	}

	for fm := range masks {
		x := digitCount(fm)
		if x == limit {
			continue
		}
		ret[x][fm] = struct{}{}
	}

	return ret
}

func digitCount(s string) int {
	x := 0
	for _, r := range s {
		if r == 'd' {
			x++
		}
	}

	return x
}

func genInnerMasks(start, end int) map[string]struct{} {
	innerMasks := map[string]struct{}{}

	for i := start; i <= end; i++ {
		startMasks := map[string]struct{}{}
		genInnerMasksRec(start, i-1, "", startMasks)
		endMasks := map[string]struct{}{}
		genInnerMasksRec(i+2, end+1, "", endMasks)
		for im := range startMasks {
			for em := range endMasks {
				var sb strings.Builder
				sb.WriteString(im)
				sb.WriteString("dd")
				sb.WriteString(em)
				innerMasks[sb.String()] = struct{}{}
			}
		}
	}

	return innerMasks
}

func genInnerMasksRec(start, end int, soFar string, innerMasks map[string]struct{}) {
	if end < start {
		innerMasks[""] = struct{}{}
		return
	}

	if end == start {
		innerMasks[soFar+"d"] = struct{}{}
		innerMasks[soFar+"m"] = struct{}{}
		return
	}

	genInnerMasksRec(start+1, end, soFar+"d", innerMasks)
	genInnerMasksRec(start+1, end, soFar+"m", innerMasks)
}
