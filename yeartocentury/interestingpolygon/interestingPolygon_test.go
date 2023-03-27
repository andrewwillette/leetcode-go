package interestingpolygon

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// 1 = 1 (+1)
// 2 = 5 (+4)
// 3 = 13 (+8)
// 4 = 25 (+12)
// 5 = 41 (+16)
func interestingPolygon(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		if i == 1 {
			result += 1
		} else {
			result += ((i - 1) * 4)
		}
	}
	return result
}

func TestInterestingPolygon(t *testing.T) {
	var cases = []struct {
		input    int
		expected int
	}{
		{
			input:    2,
			expected: 5,
		},
		{
			input:    3,
			expected: 13,
		},
	}
	for _, c := range cases {
		result := interestingPolygon(c.input)
		require.Equal(t, c.expected, result)
	}
}
