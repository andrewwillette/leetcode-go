package phonelettercombinations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func letterCombinations(digits string) []string {
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var res []string
	if digits == "" {
		return res
	}
	backtrack("", &res, 0, digits, m)
	return res
}

func backtrack(curStr string, res *[]string, index int, digits string, digitsMap map[byte][]string) {
	fmt.Printf("%v\nres: %v\nindex: %v\n", curStr, *res, index)
	if index == len(digits) {
		*res = append(*res, curStr)
		return
	}
	for _, ch := range digitsMap[digits[index]] {
		backtrack(curStr+ch, res, index+1, digits, digitsMap)
	}
}

func TestLetterCombinations(t *testing.T) {
	var cases = []struct {
		input    string
		expected []string
	}{
		{
			input:    "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := letterCombinations(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
