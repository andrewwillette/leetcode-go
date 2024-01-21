package climbstairs

// https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	default:
		// fibonacci
		allWays := 0
		nMinusTwo := 1
		nMinusOne := 2
		for i := 2; i < n; i++ {
			allWays = nMinusOne + nMinusTwo
			nMinusTwo = nMinusOne
			nMinusOne = allWays
		}
		return allWays
	}
}
