package checkpalindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func checkPalindrome(inputString string) bool {
	length := len(inputString)
	chars := []rune(inputString)
	for i := 0; i < len(chars)/2; i++ {
		if chars[i] != chars[length-1-i] {
			return false
		}
	}
	return true
}

func TestCheckPalindrome(t *testing.T) {
	var cases = []struct {
		input    string
		expected bool
	}{
		{
			input:    "aabaa",
			expected: true,
		},
		{
			input:    "abac",
			expected: false,
		},
		{
			input:    "a",
			expected: true,
		},
	}
	for _, c := range cases {
		result := checkPalindrome(c.input)
		require.Equal(t, c.expected, result)
	}
}
