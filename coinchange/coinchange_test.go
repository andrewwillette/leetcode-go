package coinchange

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	var dfs func(remain int) int
	var memo = make([]*int, amount+1)
	dfs = func(remain int) int {
		if memo[remain] != nil {
			return *memo[remain]
		}
		if remain == 0 {
			return 0
		}
		if remain < 0 {
			return math.MaxInt32
		}
		numberOfCoins := math.MaxInt32
		for _, c := range coins {
			if c > remain {
				continue
			}
			numberOfCoins = min(numberOfCoins, 1+dfs(remain-c))
		}
		memo[remain] = &numberOfCoins
		return numberOfCoins
	}
	numberOfCoins := dfs(amount)
	if numberOfCoins == math.MaxInt32 {
		return -1
	}
	return numberOfCoins
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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
		{
			coins:    []int{1, 2, 5},
			amount:   100,
			expected: 20,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := coinChange(c.coins, c.amount)
			require.Equal(t, c.expected, result)
		})
	}
}
