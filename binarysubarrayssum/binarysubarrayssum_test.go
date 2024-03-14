package binarysubarrayssum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/binary-subarrays-with-sum
func numSubarraysWithSum(nums []int, goal int) int {
	total := 0
	for i, v := range nums {
		sum := v
		if sum == goal {
			total++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == goal {
				total++
			}
			if sum > goal {
				break
			}
		}
	}
	return total
}

func TestNumSubarraysWithSum(t *testing.T) {
	var cases = []struct {
		nums     []int
		goal     int
		expected int
	}{
		{
			nums:     []int{1, 0, 1, 0, 1},
			goal:     2,
			expected: 4,
		},
		{
			nums:     []int{0, 0, 0, 0, 0},
			goal:     0,
			expected: 15,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := numSubarraysWithSum(c.nums, c.goal)
			require.Equal(t, c.expected, result)
		})
	}
}
