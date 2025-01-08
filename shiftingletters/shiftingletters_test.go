package shiftingletters

// #2381. Shifting Letters II
// https://leetcode.com/problems/shifting-letters-ii
// difference array pattern!
func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	differenceArray := make([]int, n+1)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			differenceArray[start] += 1
			differenceArray[end+1] -= 1
		} else {
			differenceArray[start] -= 1
			differenceArray[end+1] += 1
		}
	}

	currentShift := 0
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		currentShift += differenceArray[i]
		runes[i] = shiftRune(runes[i], currentShift)
	}
	return string(runes)
}

func shiftRune(runeToShift rune, shiftCount int) rune {
	alphabetLength := 26
	base := int('a')

	newPos := (int(int(runeToShift)-base) + shiftCount) % alphabetLength
	if newPos < 0 {
		newPos += alphabetLength
	}
	return rune(base + newPos)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
