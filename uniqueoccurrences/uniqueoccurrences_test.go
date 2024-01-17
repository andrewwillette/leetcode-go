package uniqueoccurrences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}
	freqExists := make(map[int]bool)
	for _, v := range freq {
		if _, ok := freqExists[v]; ok {
			return false
		}
		freqExists[v] = true
	}
	return true
}

func TestUniqueOccurrences(t *testing.T) {
	var cases = []struct {
		arr      []int
		expected bool
	}{
		{
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			arr:      []int{1, 2},
			expected: false,
		},
		{
			arr:      []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := uniqueOccurrences(c.arr)
			require.Equal(t, c.expected, result)
		})
	}
}
