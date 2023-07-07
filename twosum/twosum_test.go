package twosum

func twoSum(nums []int, target int) []int {
	// exists - the number and it's index
	exists := map[int]int{}
	for i, num := range nums {
		if val, ok := exists[target-num]; ok {
			return []int{i, val}
		}
		exists[num] = i
	}
	return []int{}
}
