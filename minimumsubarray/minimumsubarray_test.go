package minimumsubarray

import "math"

// https://leetcode.com/problems/minimum-size-subarray-sum
func minSubArrayLen(target int, nums []int) int {
	j, sum, count, min := 0, 0, 0, math.MaxInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return 1
		}

		if nums[i]+sum >= target {
			if count < min {
				min = count
			}
			sum -= nums[j]
			j++
			i--
			count--
		} else {
			sum += nums[i]
			count++
		}
	}
	if min == math.MaxInt64 {
		return 0
	}
	return min + 1
}
