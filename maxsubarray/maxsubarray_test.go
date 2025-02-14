package maxsubarray

import "testing"

// Q: Maximum Subarray Sum Problem (https://leetcode.com/problems/maximum-subarray/):
//    "Here's an algorithmic problem: find the contiguous subarray with the largest sum."
//    "Given an array like [-2, 1, -3, 4, -1, 2, 1, -5, 4], what's the largest sum?"
//    Prompt: "Can you solve this using a dynamic programming approach in Go?"

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}

func TestMaxSubArray(t *testing.T) {
	t.Run("maxSubArray", func(t *testing.T) {
		tests := []struct {
			nums []int
			want int
		}{
			{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
			{[]int{1}, 1},
			{[]int{5, 4, -1, 7, 8}, 23},
		}

		for _, tt := range tests {
			got := maxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("maxSubArray(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		}
	})
}
