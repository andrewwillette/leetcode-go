package longestsubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func lengthOfLongestSubstring(s string) int {
	runesInSequence := make(map[rune]bool, len(s))
	longestSubString := 0
	currentSubString := 0
	for _, r := range s {
		if _, ok := runesInSequence[r]; ok {
			if currentSubString > longestSubString {
				longestSubString = currentSubString
			}
			runesInSequence = make(map[rune]bool, len(s))
			currentSubString = 1
			runesInSequence[r] = true
		} else {
			runesInSequence[r] = true
			currentSubString++
		}
	}
	if currentSubString > longestSubString {
		longestSubString = currentSubString
	}
	return longestSubString
}

func TestLengthOfLongestSubstring(t *testing.T) {
	var cases = []struct {
		s        string
		expected int
	}{
		// {
		// 	s:        "abcabcbb",
		// 	expected: 3,
		// },
		// {
		// 	s:        "bbbbb",
		// 	expected: 1,
		// },
		// {
		// 	s:        "pwwkew",
		// 	expected: 3,
		// },
		{
			s:        "dvdf",
			expected: 3,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := lengthOfLongestSubstring(c.s)
			require.Equal(t, c.expected, result)
		})
	}
}
