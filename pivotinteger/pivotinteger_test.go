package pivotinteger

// https://leetcode.com/problems/find-the-pivot-integer
func pivotInteger(n int) int {
	sum := getSum(n)
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		firstPart := getSum(mid)
		secondPart := sum - firstPart + mid
		if firstPart == secondPart {
			return mid
		} else if firstPart > secondPart {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func getSum(x int) int {
	return x * (x + 1) / 2
}
