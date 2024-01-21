package decodeways

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/decode-ways/
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	memo := []int{}
	for i := 0; i <= len(s); i++ {
		memo = append(memo, -1)
	}
	return numDecodingsSub(0, s, memo)
}

func numDecodingsSub(p int, s string, memo []int) int {
	if memo[p] > -1 {
		return memo[p]
	}
	n := len(s)
	if p == n {
		return 1
	}
	if s[p] == '0' {
		return 0
	}
	res := numDecodingsSub(p+1, s, memo)
	if p < n-1 &&
		// checks if a valid number
		(s[p] == '1' || (s[p] == '2' && s[p+1] < '7')) {
		res += numDecodingsSub(p+2, s, memo)
	}
	memo[p] = res
	return res
}

func TestNumDecodings(t *testing.T) {
	var cases = []struct {
		s        string
		expected int
	}{
		{
			s:        "12",
			expected: 2,
		},
		{
			s:        "226",
			expected: 3,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := numDecodings(c.s)
			require.Equal(t, c.expected, result)
		})
	}
}
