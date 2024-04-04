package maxnestdepth

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func maxDepth(s string) int {
	maxV := 0
	count := 0
	for _, c := range s {
		if c == '(' {
			count++
			if count > maxV {
				maxV = count
			}
		} else if c == ')' {
			count--
		}
	}
	return maxV
}

func TestMaxDepth(t *testing.T) {
	var cases = []struct {
		s        string
		expected int
	}{
		{
			s:        "(1+(2*3)+((8)/4))+1",
			expected: 3,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := maxDepth(c.s)
			require.Equal(t, c.expected, result)
		})
	}
}
