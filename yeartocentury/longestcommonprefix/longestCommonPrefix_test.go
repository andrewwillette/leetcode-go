package longestcommonprefix

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func longestCommonPrefixBetter(strs []string) string {
	var p = strs[0]
	for _, s := range strs {
		i := 0
		for ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {
		}
		p = p[:i]
	}
	return p
}

func TestLongestCommonPrefixBetter(t *testing.T) {
	var cases = []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"flower", "flower", "flower"},
			expected: "flower",
		},
		{
			input:    []string{"ab", "a"},
			expected: "a",
		},
	}
	for _, c := range cases {
		result := longestCommonPrefixBetter(c.input)
		require.Equal(t, c.expected, result)
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	longestPrefix := ""
	firstItem := strs[0]
	for i := 1; i < len(firstItem)+1; i++ {
		possibleLongestPrefix := firstItem[0:i]
		for j := 1; j < len(strs); j++ {
			checkPrefixString := strs[j]
			if len(checkPrefixString) < i {
				return longestPrefix
			}
			if checkPrefixString[0:i] != possibleLongestPrefix {
				return longestPrefix
			}
			if j == len(strs)-1 {
				longestPrefix = possibleLongestPrefix
			}
		}
	}
	return longestPrefix
}
func TestLongestCommonPrefix(t *testing.T) {
	var cases = []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"flower", "flower", "flower"},
			expected: "flower",
		},
		{
			input:    []string{"ab", "a"},
			expected: "a",
		},
	}
	for _, c := range cases {
		result := longestCommonPrefix(c.input)
		require.Equal(t, c.expected, result)
	}
}
