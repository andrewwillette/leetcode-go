package longestincreasingsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/longest-increasing-subsequence/
// Given an integer array nums, return the length of the longest strictly increasing subsequence.
//
// What is the "subproblem"?
// LIS[n] = 1 + max{LIS[k] | k < n, A[k] < A[n]}
func lengthOfLIS(nums []int) int {
	l := make([]int, len(nums))
	for i := 0; i < len(l); i++ {
		l[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && l[i] < (l[j]+1) {
				l[i] = l[j] + 1
			}
		}
	}
	max := 0
	for i := range l {
		if l[i] > max {
			max = l[i]
		}
	}
	return max
}

func TestLengthOfLIS(t *testing.T) {
	var cases = []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			nums:     []int{3, 2, 1},
			expected: 1,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := lengthOfLIS(c.nums)
			require.Equal(t, c.expected, result)
		})
	}
}
