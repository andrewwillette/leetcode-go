package buildarrayfromperm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Given a zero-based permutation nums (0-indexed), build an array ans of the same length
// where ans[i] = nums[nums[i]] for each 0 <= i < nums.length and return it.
// A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).
func buildArrayFromPerm(nums []int) []int {
	toReturn := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		toReturn[i] = nums[nums[i]]
	}

	return toReturn
}

func TestAddTwoNumbers(t *testing.T) {
	var cases = []struct {
		input    []int
		expected []int
	}{
		{},
	}
	for _, c := range cases {
		result := buildArrayFromPerm(c.input)
		require.Equal(t, c.expected, result)
	}
}
