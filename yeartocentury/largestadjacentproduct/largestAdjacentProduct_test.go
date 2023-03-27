package largestadjacentproduct

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func largestAdjacentProduct(inputArray []int) int {
	var largestProduct = 0
	for i := 0; i < len(inputArray); i++ {
		if i+1 < len(inputArray) {
			product := inputArray[i] * inputArray[i+1]
			if i == 0 {
				largestProduct = product
			} else {
				if product > largestProduct {
					largestProduct = product
				}
			}
		}
	}
	return largestProduct
}

func TestLargestAdjacentProduct(t *testing.T) {
	var cases = []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{3, 6, -2, -5, 7, 3},
			expected: 21,
		},
		{
			input:    []int{-23, 4, -3, 8, -12},
			expected: -12,
		},
	}
	for _, c := range cases {
		result := largestAdjacentProduct(c.input)
		require.Equal(t, c.expected, result)
	}
}
