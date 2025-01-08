package countprefixsuffix

// #3042
// https://leetcode.com/problems/count-prefix-and-suffix-pairs-i
func countPrefixSuffixPairs(words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

// check if candidate is a prefix and suffix of word
func isPrefixAndSuffix(candidate, word string) bool {
	if len(candidate) > len(word) {
		return false
	}
	rc := []rune(candidate)
	wc := []rune(word)
	prefixSlice := wc[:len(rc)]
	suffixSlice := wc[len(wc)-len(rc):]
	if string(rc) == string(prefixSlice) && string(rc) == string(suffixSlice) {
		return true
	}
	return false
}
