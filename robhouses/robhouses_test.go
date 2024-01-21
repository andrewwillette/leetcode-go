package robhouses

// M(k) = money at the kth house
// P(0) = 0
// P(1) = M(1)
// P(k) = max(P(k−2) + M(k), P(k−1))
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	l1 := nums[0]
	l2 := max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		l2, l1 = max(l1+nums[i], l2), l2
	}
	return l2
}
