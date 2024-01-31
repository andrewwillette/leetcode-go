package reversepolishnotation

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/evaluate-reverse-polish-notation
func evalRPN(tokens []string) int {
	for {
		if len(tokens) == 1 {
			break
		}
		tokens = evalNextOperator(tokens)
	}
	i, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}

	return i
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func evalNextOperator(tokens []string) []string {
	output := make([]string, 0)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" {
			result := stringToInt(tokens[i-2]) + stringToInt(tokens[i-1])
			tokens[i] = fmt.Sprint(result)

			output = append(tokens[:i-2], tokens[i:]...)
			break
		} else if tokens[i] == "-" {
			result := stringToInt(tokens[i-2]) - stringToInt(tokens[i-1])
			tokens[i] = fmt.Sprintf("%v", result)
			output = append(tokens[:i-2], tokens[i:]...)
			break
		} else if tokens[i] == "*" {
			result := stringToInt(tokens[i-2]) * stringToInt(tokens[i-1])
			tokens[i] = fmt.Sprintf("%v", result)
			output = append(tokens[:i-2], tokens[i:]...)
			break
		} else if tokens[i] == "/" {
			result := stringToInt(tokens[i-2]) / stringToInt(tokens[i-1])
			tokens[i] = fmt.Sprintf("%v", result)
			output = append(tokens[:i-2], tokens[i:]...)
			break
		}
	}
	return output
}

func TestEvalRPN(t *testing.T) {
	var cases = []struct {
		tokens   []string
		expected int
	}{
		{
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := evalRPN(c.tokens)
			require.Equal(t, c.expected, result)
		})
	}
}
