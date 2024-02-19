package poweroftwo

// https://leetcode.com/problems/power-of-two
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	// check if n is a power of 2
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}
