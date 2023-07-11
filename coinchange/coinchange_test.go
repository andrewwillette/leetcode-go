package coinchange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	return 0
}

func TestCoinChange(t *testing.T) {
	var cases = []struct {
		coins    []int
		amount   int
		expected int
	}{
		{
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := coinChange(c.coins, c.amount)
			require.Equal(t, c.expected, result)
		})
	}
}
