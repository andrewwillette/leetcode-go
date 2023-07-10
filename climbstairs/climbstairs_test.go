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
		oneStepBefore := 2
		twoStepsBefore := 1
		allWays := 0
		for i := 2; i < n; i++ {
			allWays = oneStepBefore + twoStepsBefore
			twoStepsBefore = oneStepBefore
			oneStepBefore = allWays
		}
		return allWays
	}
}
