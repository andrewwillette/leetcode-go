package minlengthafteroperations

// # 3223
// https://leetcode.com/problems/minimum-length-of-string-after-operations
func minimumLength(s string) int {
	runes := []rune(s)
	freqArray := make([]int, 26)
	for _, r := range runes {
		freqArray[r-'a']++
	}
	count := 0
	for _, freq := range freqArray {
		if freq == 0 {
			continue
		}
		if freq%2 == 0 {
			count = count + 2
		} else {
			count++
		}
	}
	return count
}
