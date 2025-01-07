package shiftingletters

// #2381. Shifting Letters II
// https://leetcode.com/problems/shifting-letters-ii
func shiftingLetters(s string, shifts [][]int) string {
	differenceArray := make([]int, len(s))
	for i := 0; i < len(shifts); i++ {
		for j := shifts[i][0]; j <= shifts[i][1]; j++ {
			if shifts[i][2] > 0 { // shift forwards
				differenceArray[j] += differenceArray[j] + 1
			} else { // shift backwards
				differenceArray[j] += differenceArray[j] - 1
			}
		}
	}
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		runes[i] = shiftRune(runes[i], differenceArray[i])
	}
	return string(runes)
}

func shiftRune(runeToShift rune, shiftCount int) rune {
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alphabetLength := len(alphabet)
	var toReturn rune
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] == runeToShift {
			index := i + shiftCount
			if index < 0 {
				index = abs(index)
				indexModed := index % alphabetLength
				// -1 is the last element in the array
				toReturn = alphabet[alphabetLength-1-indexModed]
			} else {
				indexModed := index % alphabetLength
				toReturn = alphabet[indexModed]
			}
		}
	}
	return toReturn
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
