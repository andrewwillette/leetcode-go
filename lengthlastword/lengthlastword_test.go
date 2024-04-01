package lengthlastword

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	words := strings.Split(s, " ")
	for i := len(words) - 1; i >= 0; i-- {
		word := strings.TrimSpace(words[i])
		if len(word) > 0 {
			return len(words[i])
		}
	}
	return len(words[len(words)-1])
}

func TestLengthOfLastWord(t *testing.T) {
	var cases = []struct {
		s        string
		expected int
	}{
		{
			s:        "Hello World",
			expected: 5,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := lengthOfLastWord(c.s)
			require.Equal(t, c.expected, result)
		})
	}
}
