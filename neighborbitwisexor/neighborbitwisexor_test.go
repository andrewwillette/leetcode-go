package neighborbitwisexor

// # 2683
// https://leetcode.com/problems/neighboring-bitwise-xor
func doesValidArrayExist(derived []int) bool {
	xorSum := 0
	for _, value := range derived {
		xorSum ^= value
	}
	return xorSum == 0
}
