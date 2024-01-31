package dailytemp

// https://leetcode.com/problems/daily-temperatures/
func dailyTemperatures(temperatures []int) []int {
	result := []int{}

	memo := make(map[int][]int)
	// populate memo
	for i, vi := range temperatures {
		if _, ok := memo[vi]; !ok {
			memo[vi] = []int{i}
		} else {
			memo[vi] = append(memo[vi], i)
		}
	}

	for i := range temperatures {
		if i == len(temperatures)-1 {
			result = append(result, 0)
		}
		for j := 0; j < len(memo); j++ {
		}
	}
	return result
}
