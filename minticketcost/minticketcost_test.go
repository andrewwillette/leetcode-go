package minticketcost

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// mincostTickets
// https://leetcode.com/problems/minimum-cost-for-tickets/
func mincostTickets(days []int, costs []int) int {
	var dp [396]int
	for _, d := range days {
		dp[d] = 1
	}
	for d := 365; d > 0; d-- {
		if dp[d] == 0 {
			dp[d] = dp[d+1]
			continue
		}
		dp[d] = min(
			dp[d+1]+costs[0],
			dp[d+7]+costs[1],
			dp[d+30]+costs[2])
	}
	return dp[1]
}

func min(nums ...int) int {
	for _, n := range nums {
		if n < nums[0] {
			nums[0] = n
		}
	}
	return nums[0]
}

func TestMincostTickets(t *testing.T) {
	var cases = []struct {
		days     []int
		costs    []int
		expected int
	}{
		{
			days:     []int{1, 4, 6, 7, 8, 20},
			costs:    []int{2, 7, 15},
			expected: 11,
		},
		{
			days:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			costs:    []int{2, 7, 15},
			expected: 17,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := mincostTickets(c.days, c.costs)
			require.Equal(t, c.expected, result)
		})
	}
}
