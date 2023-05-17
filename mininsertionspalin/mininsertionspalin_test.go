package mininsertionspalin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// minInsertions returns the minimum number of insertions needed to make a string a palindrome.
// https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for d := 1; d < n; d++ { // d is distance
		for i := 0; i+d < n; i++ {
			if s[i] == s[i+d] {
				dp[i][i+d] = dp[i+1][i+d-1]
			} else {
				dp[i][i+d] = 1 + min2(dp[i][i+d-1], dp[i+1][i+d])
			}
		}
	}
	return dp[0][n-1]
}

func min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func TestMinInsertions(t *testing.T) {
	var cases = []struct {
		s        string
		expected int
	}{
		{
			s:        "zzazz",
			expected: 0,
		},
		{
			s:        "mbadm",
			expected: 2,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := minInsertions(c.s)
			require.Equal(t, c.expected, result)
		})
	}
}
