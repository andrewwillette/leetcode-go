package subarrayproductlessk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/subarray-product-less-than-k
func numSubarrayProductLessThanK(nums []int, k int) int {
	var left, result int
	curr := 1
	for right := 0; right < len(nums); right++ {
		curr *= nums[right]
		for left <= right && curr >= k {
			curr /= nums[left]
			left++
		}
		result += right - left + 1
	}
	return result
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	var cases = []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{10, 5, 2, 6},
			k:        100,
			expected: 8,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := numSubarrayProductLessThanK(c.nums, c.k)
			require.Equal(t, c.expected, result)
		})
	}
}
