package majorityelement

// https://leetcode.com/problems/majority-element/
func majorityElement(nums []int) int {
	memo := make(map[int]int)
	for _, v := range nums {
		memo[v] = memo[v] + 1
		if memo[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}
