package sortvowels

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/sort-vowels-in-a-string
func sortVowels(s string) string {
	bytes := []byte(s)
	vowels := make([]byte, 0)
	for i := 0; i < len(bytes); i++ {
		if isVowel(bytes[i]) {
			vowels = append(vowels, bytes[i])
		}
	}
	sort.Slice(vowels, func(i, j int) bool {
		return vowels[i] < vowels[j]
	})

	ptr := 0
	for i := 0; i < len(bytes); i++ {
		if isVowel(bytes[i]) {
			bytes[i] = vowels[ptr]
			ptr++
		}
	}
	return string(bytes)
}

const vowels = "aeiouAEIOU"

func isVowel(char byte) bool {
	return strings.ContainsRune(vowels, rune(char))
}
