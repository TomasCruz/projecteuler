package main

type arrangement struct {
	value    int
	hasDigit [10]byte
}

func makeArr(sl []byte) arrangement {
	arr := arrangement{}
	digCount := 0
	digValue := 0
	for i := 0; i < len(sl); i++ {
		if i == 6 || i == 9 || sl[i] == 0 {
			continue
		}

		digCount++
		digValue += 1 << i
		arr.hasDigit[i] = 1
	}

	if sl[6] == 1 || sl[9] == 1 {
		digCount += 2
		digValue += 1 << 6
		digValue += 1 << 9
		arr.hasDigit[6] = 1
		arr.hasDigit[9] = 1
	}

	arr.value = digCount<<12 + digValue

	return arr
}

func (arr arrangement) isValid() bool {
	if !(arr.hasDigit[0] == 1 || arr.hasDigit[1] == 1) {
		return false
	}
	if !(arr.hasDigit[0] == 1 || arr.hasDigit[4] == 1) {
		return false
	}
	if !(arr.hasDigit[0] == 1 || arr.hasDigit[9] == 1) {
		return false
	}
	if !(arr.hasDigit[1] == 1 || arr.hasDigit[6] == 1) {
		return false
	}
	if !(arr.hasDigit[2] == 1 || arr.hasDigit[5] == 1) {
		return false
	}
	if !(arr.hasDigit[3] == 1 || arr.hasDigit[6] == 1) {
		return false
	}
	if !(arr.hasDigit[4] == 1 || arr.hasDigit[9] == 1) {
		return false
	}
	if !(arr.hasDigit[4] == 1 || arr.hasDigit[6] == 1) {
		return false
	}
	if !(arr.hasDigit[1] == 1 || arr.hasDigit[8] == 1) {
		return false
	}
	return true
}

// this function presupposes validity of arr
func (arr arrangement) isComplementValid(c arrangement) bool {
	// 01
	if arr.hasDigit[1] == 0 && c.hasDigit[1] == 0 {
		return false
	} else if arr.hasDigit[0] == 0 && c.hasDigit[0] == 0 {
		return false
	} else {
		if c.hasDigit[0] == 0 && c.hasDigit[1] == 0 {
			return false
		}
	}

	// 04
	if arr.hasDigit[4] == 0 && c.hasDigit[4] == 0 {
		return false
	} else if arr.hasDigit[0] == 0 && c.hasDigit[0] == 0 {
		return false
	} else {
		if c.hasDigit[0] == 0 && c.hasDigit[4] == 0 {
			return false
		}
	}

	// 09
	if arr.hasDigit[9] == 0 && c.hasDigit[9] == 0 {
		return false
	} else if arr.hasDigit[0] == 0 && c.hasDigit[0] == 0 {
		return false
	} else {
		if c.hasDigit[0] == 0 && c.hasDigit[9] == 0 {
			return false
		}
	}

	// 16
	if arr.hasDigit[6] == 0 && c.hasDigit[6] == 0 {
		return false
	} else if arr.hasDigit[1] == 0 && c.hasDigit[1] == 0 {
		return false
	} else {
		if c.hasDigit[6] == 0 && c.hasDigit[1] == 0 {
			return false
		}
	}

	// 25
	if arr.hasDigit[5] == 0 && c.hasDigit[5] == 0 {
		return false
	} else if arr.hasDigit[2] == 0 && c.hasDigit[2] == 0 {
		return false
	} else {
		if c.hasDigit[2] == 0 && c.hasDigit[5] == 0 {
			return false
		}
	}

	// 36
	if arr.hasDigit[6] == 0 && c.hasDigit[6] == 0 {
		return false
	} else if arr.hasDigit[3] == 0 && c.hasDigit[3] == 0 {
		return false
	} else {
		if c.hasDigit[3] == 0 && c.hasDigit[6] == 0 {
			return false
		}
	}

	// 49
	if arr.hasDigit[9] == 0 && c.hasDigit[9] == 0 {
		return false
	} else if arr.hasDigit[4] == 0 && c.hasDigit[4] == 0 {
		return false
	} else {
		if c.hasDigit[4] == 0 && c.hasDigit[9] == 0 {
			return false
		}
	}

	// 64
	if arr.hasDigit[4] == 0 && c.hasDigit[4] == 0 {
		return false
	} else if arr.hasDigit[6] == 0 && c.hasDigit[6] == 0 {
		return false
	} else {
		if c.hasDigit[4] == 0 && c.hasDigit[6] == 0 {
			return false
		}
	}

	// 81
	if arr.hasDigit[1] == 0 && c.hasDigit[1] == 0 {
		return false
	} else if arr.hasDigit[8] == 0 && c.hasDigit[8] == 0 {
		return false
	} else {
		if c.hasDigit[8] == 0 && c.hasDigit[1] == 0 {
			return false
		}
	}

	return true
}
