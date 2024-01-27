package longestcommonsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/longest-common-subsequence

// longestCommonSubsequenceDP solved with dynamic programming
func longestCommonSubsequenceDP(text1 string, text2 string) int {
	// dp[i][j] represents the longest common subsequence of text1[0:i] and text2[0:j]
	n := len(text1)
	m := len(text2)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// base case
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	result := dp[n][m]
	return result
}

func TestLongestCommonSubsequenceDP(t *testing.T) {
	var cases = []struct {
		text1    string
		text2    string
		expected int
	}{
		{
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			text1:    "bl",
			text2:    "yby",
			expected: 1,
		},
		{
			text1:    "ezupkr",
			text2:    "ubmrapg",
			expected: 2,
		},
		{
			text1:    "oxcpqrsvwf",
			text2:    "shmtulqrypy",
			expected: 2,
		},
		{
			text1:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			text2:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			expected: 210,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := longestCommonSubsequenceDP(c.text1, c.text2)
			require.Equal(t, c.expected, result)
		})
	}
}

func longestCommonSubsequenceRecursion(text1, text2 string) int {
	type memoKey struct {
		I int
		J int
	}
	memo := make(map[memoKey]int)
	var helper func(i, j int) int
	helper = func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		if val, ok := memo[memoKey{I: i, J: j}]; ok {
			return val
		}

		res := 0
		if text1[i] == text2[j] {
			res = 1 + helper(i-1, j-1)
		}
		if text1[i] != text2[j] {
			res = max(helper(i-1, j), helper(i, j-1))
		}
		memo[memoKey{I: i, J: j}] = res
		return res
	}
	return helper(len(text1)-1, len(text2)-1)
}

func TestLongestCommonSubsequenceRecursion(t *testing.T) {
	var cases = []struct {
		text1    string
		text2    string
		expected int
	}{
		{
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			text1:    "bl",
			text2:    "yby",
			expected: 1,
		},
		{
			text1:    "ezupkr",
			text2:    "ubmrapg",
			expected: 2,
		},
		{
			text1:    "oxcpqrsvwf",
			text2:    "shmtulqrypy",
			expected: 2,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := longestCommonSubsequenceRecursion(c.text1, c.text2)
			require.Equal(t, c.expected, result)
		})
	}
}
