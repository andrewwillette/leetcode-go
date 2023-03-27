package reverseinparentheses

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Write a function that reverses characters in (possibly nested)
// parentheses in the input string.
// Input strings will always be well-formed with matching ()s.
func reverseInParentheses(inputString string) string {
	stack, tmp := make([]rune, 0, len(inputString)), make([]rune, 0, len(inputString))
	for _, c := range inputString {
		if c == ')' {
			i := len(stack) - 1
			tmp = tmp[:0]
			for ; stack[i] != '('; i-- {
				tmp = append(tmp, stack[i])
			}
			stack = append(stack[:i], tmp...)
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

func TestReverseInParentheses(t *testing.T) {
	var cases = []struct {
		input    string
		expected string
	}{
		{
			input:    "(bar)",
			expected: "rab",
		},
		{
			input:    "foo(bar(baz))blim",
			expected: "foobazrabblim",
		},
	}
	for _, c := range cases {
		result := reverseInParentheses(c.input)
		require.Equal(t, c.expected, result)
	}
}
