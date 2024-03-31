package countsubarrayswithfixedbound

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// You are given an integer array nums and two integers minK and maxK.
// A fixed-bound subarray of nums is a subarray that satisfies the following conditions:
// The minimum value in the subarray is equal to minK.
// The maximum value in the subarray is equal to maxK.
// Return the number of fixed-bound subarrays.
// A subarray is a contiguous part of an array.
// https://leetcode.com/problems/count-subarrays-with-fixed-bounds
func countSubarrays(nums []int, minK int, maxK int) int64 {
	var result int64
	left := 0
	pmin, pmax := -1, -1
	for right, num := range nums {
		if num < minK || num > maxK {
			left = right + 1
			pmin, pmax = -1, -1
		} else {
			if num == minK {
				pmin = right
			}
			if num == maxK {
				pmax = right
			}
			result += int64(max(0, min(pmin, pmax)-left+1))
		}
	}
	return int64(result)
}

func TestCountSubarrays(t *testing.T) {
	var cases = []struct {
		nums     []int
		minK     int
		maxK     int
		expected int64
	}{
		{
			nums:     []int{1, 3, 5, 2, 7, 5},
			minK:     1,
			maxK:     5,
			expected: 2,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := countSubarrays(c.nums, c.minK, c.maxK)
			require.Equal(t, c.expected, result)
		})
	}
}
