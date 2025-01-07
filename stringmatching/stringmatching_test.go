package stringmatching

import "strings"

// Given an array of string words, return all strings in words that is a substring of another word. You can return the answer in any order.
// A substring is a contiguous sequence of characters within a string
func stringMatching(words []string) []string {
	toReturn := []string{}
	for _, v := range words {
		exists := false
		for _, w := range words {
			if v != w && strings.Contains(w, v) {
				exists = true
			}
		}
		if exists {
			toReturn = append(toReturn, v)
		}
	}
	return toReturn
}
