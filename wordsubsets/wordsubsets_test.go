package wordsubsets

// # 916
// https://leetcode.com/problems/word-subsets/
// You are given two string arrays words1 and words2.
// A string b is a subset of string a if every letter in b occurs in a including multiplicity.
// For example, "wrr" is a subset of "warrior" but is not a subset of "world".
// A string a from words1 is universal if for every string b in words2, b is a subset of a.
// Return an array of all the universal strings in words1. You may return the answer in any order.
func wordSubsets(words1 []string, words2 []string) []string {
	maxFreq := make([]int, 26)
	for _, w := range words2 {
		freq := charFrequency(w)
		for i := 0; i < 26; i++ {
			maxFreq[i] = max(maxFreq[i], freq[i])
		}
	}
	result := make([]string, 0)
	for _, w := range words1 {
		wordFreq := charFrequency(w)
		if isUniversal(wordFreq, maxFreq) {
			result = append(result, w)
		}
	}
	return result
}

func charFrequency(s string) []int {
	freq := make([]int, 26)
	for _, c := range s {
		// Characters in Go (like 'a', 'b', etc.) are represented by their ASCII values.
		// 'a' has an ASCII value of 97, 'b' is 98, 'c' is 99, and so on.
		// when you do c - 'a', you calculate the zero-based index of c in the alphabet:
		freq[c-'a']++
	}
	return freq
}

func isUniversal(wordFreq, maxFreq []int) bool {
	for i := 0; i < 26; i++ {
		if wordFreq[i] < maxFreq[i] {
			return false
		}
	}
	return true
}
