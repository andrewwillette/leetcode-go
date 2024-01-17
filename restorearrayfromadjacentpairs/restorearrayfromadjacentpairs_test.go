package restorearrayfromadjacentpairs

// https://leetcode.com/problems/restore-the-array-from-adjacent-pairs
func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	m := make(map[int][]int, n)
	for _, v := range adjacentPairs {
		a, b := v[0], v[1]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	start := -1
	for k := range m {
		if len(m[k]) == 1 {
			start = k
			break
		}
	}
	ans := make([]int, n)
	ans[0] = start
	prev := start
	start = m[start][0]
	for i := 1; i < n-1; i++ {
		ans[i] = start
		if prev != m[start][0] {
			prev = start
			start = m[start][0]
		} else {
			prev = start
			start = m[start][1]
		}
	}
	ans[n-1] = start
	return ans
}
