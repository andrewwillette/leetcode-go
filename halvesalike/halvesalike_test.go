package halvesalike

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/determine-if-string-halves-are-alike/
func halvesAreAlike(s string) bool {
	// split string in half
	firstHalf := s[:len(s)/2]
	secondHalf := s[len(s)/2:]
	return countVowels(firstHalf) == countVowels(secondHalf)
}

func countVowels(s string) int {
	var vowelsCount int
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
			char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			vowelsCount++
		}
	}
	return vowelsCount
}

func TestHalvesAreAlike(t *testing.T) {
	var cases = []struct {
		s        string
		expected bool
	}{
		{
			s:        "book",
			expected: true,
		},
		{
			s:        "textbook",
			expected: false,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := halvesAreAlike(c.s)
			require.Equal(t, c.expected, result)
		})
	}
}
