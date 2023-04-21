package numways

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// numWays
// https://leetcode.com/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/
func numWays(words []string, target string) int {
	n := len(words[0])
	m := len(target)
	mod := 1000000007
	dp := make([]int, m+1)
	dp[0] = 1

	// a map representing each word's count of each alphabetical letter in each position
	count := make([][]int, n)
	const alphabetSize = 26
	for i := range count {
		count[i] = make([]int, alphabetSize)
	}

	for i := 0; i < n; i++ {
		for _, word := range words {
			count[i][int(word[i]-'a')]++
		}
	}
	for i := 0; i < n; i++ {
		for j := m - 1; j >= 0; j-- {
			dp[j+1] += dp[j] * count[i][int(target[j]-'a')]
			dp[j+1] %= mod
		}
	}

	return dp[m]
}

func TestNumWays(t *testing.T) {
	var cases = []struct {
		words    []string
		target   string
		expected int
	}{
		{
			words:    []string{"acca", "bbbb", "caca"},
			target:   "aba",
			expected: 6,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := numWays(c.words, c.target)
			require.Equal(t, c.expected, result)
		})
	}
}
