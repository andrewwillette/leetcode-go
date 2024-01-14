package closeststrings

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/determine-if-two-strings-are-close/
func closeStrings(word1 string, word2 string) bool {
	counter := func(word string) (keys, vals [26]int) {
		for i := range word {
			keys[word[i]-'a'] = 1
			vals[word[i]-'a'] += 1
		}
		sort.Ints(vals[:])
		return keys, vals
	}
	keys1, vals1 := counter(word1)
	keys2, vals2 := counter(word2)
	return keys1 == keys2 && vals1 == vals2
}

func TestCloseStrings(t *testing.T) {
	var cases = []struct {
		word1    string
		word2    string
		expected bool
	}{
		{
			word1:    "abc",
			word2:    "bca",
			expected: true,
		},
		{
			word1:    "a",
			word2:    "aa",
			expected: false,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := closeStrings(c.word1, c.word2)
			require.Equal(t, c.expected, result)
		})
	}
}
