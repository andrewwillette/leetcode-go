package minsteps

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/
func minSteps(s string, t string) int {
	freqs, freqt := [26]int{}, [26]int{}
	for i := 0; i < len(s); i++ {
		freqs[s[i]-'a']++
		freqt[t[i]-'a']++
	}
	result := 0.0
	for i := 0; i < 26; i++ {
		result += math.Abs(float64(freqs[i] - freqt[i]))
	}
	return int(result) / 2
}

func TestMinSteps(t *testing.T) {
	var cases = []struct {
		string
		t        string
		s        string
		expected int
	}{
		{
			t:        "bab",
			s:        "aba",
			expected: 1,
		},
		{
			t:        "leetcode",
			s:        "practice",
			expected: 5,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := minSteps(c.t, c.s)
			require.Equal(t, c.expected, result)
		})
	}
}
