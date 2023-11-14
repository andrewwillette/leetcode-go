package checkifarraygood

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func isGood(nums []int) bool {
	memo := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		memo[nums[i]] = memo[nums[i]] + 1
	}
	if len(memo) != len(nums)-1 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		if i == len(nums)-1 {
			if memo[i] != 2 {
				return false
			}
			continue
		}
		if memo[i] != 1 {
			return false
		}

	}
	return true
}

func TestIsGood(t *testing.T) {
	var cases = []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 3, 3, 2},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := isGood(c.nums)
			require.Equal(t, c.expected, result)
		})
	}
}
