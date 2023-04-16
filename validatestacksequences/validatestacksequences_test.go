package validatestacksequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		return *new(stack), 0
	}
	return s[:l-1], s[l-1]
}

// validateStackSequences
// https://leetcode.com/problems/validate-stack-sequences/
func validateStackSequences(pushed []int, popped []int) bool {
	s := *new(stack)
	var i, j int
	for i < len(pushed) {
		s = s.Push(pushed[i])
		for j < len(popped) && len(s) >= 1 && s[len(s)-1] == popped[j] {
			s, _ = s.Pop()
			j++
		}
		i++
	}
	return len(s) == 0
}

func TestValidateStackSequences(t *testing.T) {
	var cases = []struct {
		pushed   []int
		popped   []int
		expected bool
	}{
		{
			pushed:   []int{1, 2, 3, 4, 5},
			popped:   []int{4, 5, 3, 2, 1},
			expected: true,
		},
		{
			pushed:   []int{1, 0},
			popped:   []int{1, 0},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := validateStackSequences(c.pushed, c.popped)
			require.Equal(t, c.expected, result)
		})
	}
}
