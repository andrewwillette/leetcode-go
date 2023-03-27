package alllongeststrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func allLongestStrings(inputArray []string) []string {
	m := make(map[int][]string)
	maxLength := 0
	for _, v := range inputArray {
		length := len(v)
		if length > maxLength {
			maxLength = length
		}
		m[length] = append(m[length], v)
	}
	return m[maxLength]
}

func TestAllLongestStrings(t *testing.T) {
	var cases = []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"aba", "aa", "ad", "vcd", "aba"},
			expected: []string{"aba", "vcd", "aba"},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := allLongestStrings(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
