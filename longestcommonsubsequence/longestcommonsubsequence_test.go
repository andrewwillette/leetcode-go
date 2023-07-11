package longestcommonsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/longest-increasing-subsequence/
// Given an integer array nums, return the length of the longest strictly increasing subsequence.
//
// What is the "subproblem"?
// Given nums[x], nums[x+n] > nums[x]
func lengthOfLIS(nums []int) int {
	lis := 0
	tmpLis := 0
	for i := range nums {
		if i == len(nums)-1 {
			if tmpLis > lis {
				lis = tmpLis
			}
			break
		}
		if nums[i] < nums[i+1] {
			tmpLis++
		} else {
			if tmpLis > lis {
				lis = tmpLis
			}
			tmpLis = 0
		}
	}
	return lis
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
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := lengthOfLIS(c.nums)
			require.Equal(t, c.expected, result)
		})
	}
}
