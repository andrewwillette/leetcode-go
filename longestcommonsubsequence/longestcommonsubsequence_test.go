package longestcommonsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/longest-common-subsequence

// longestCommonSubsequenceDP solved with dynamic programming
// not working currently
func longestCommonSubsequenceDP(text1 string, text2 string) int {
	dp := [][]int{}
	// check if text2[0] exists in text1
	text20Exists := false
	for _, v := range text1 {
		if v == rune(text2[0]) {
			text20Exists = true
		}
	}
	for i, v1 := range text1 {
		baseCase := make([]int, len(text2))
		if text20Exists {
			baseCase[0] = 1
		}
		if i == 0 && v1 == rune(text2[0]) {
			// populate baseCase array with 1
			for j := 0; j < len(baseCase); j++ {
				baseCase[j] = 1
			}
		}
		for j := 0; j < len(baseCase); j++ {
			if text2[j] == text1[0] {
				baseCase[j] = 1
			}
		}
		if text1[0] == text2[0] {
			baseCase[0] = 1
		}
		dp = append(dp, baseCase)
	}

	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Printf("%v\n", dp)
	result := dp[len(text1)-1][len(text2)-1]
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
