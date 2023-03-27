package almostincreasingsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func isArrayIncreasing(sequence []int) bool {
	for i := 0; i <= len(sequence)-2; i++ {
		if sequence[i] >= sequence[i+1] {
			return false
		}
	}
	return true
}

func remove(slice []int, s int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	fmt.Printf("remove : %v\n", slice)
	return append(newSlice[:s], newSlice[s+1:]...)
}

// Given a sequence of integers as an array, determine whether it is possible to obtain a strictly increasing
// sequence by removing no more than one element from the array.
func almostIncreasingSequence(sequence []int) bool {
	for i := 0; i < len(sequence)-1; i++ {
		if sequence[i] >= sequence[i+1] {
			if i == len(sequence)-2 {
				return true
			}
			fmt.Printf("sequence before removed: %v\n", sequence)
			greaterRemoved := remove(sequence, i)
			fmt.Printf("sequence after removed: %v\n", greaterRemoved)
			if !isArrayIncreasing(greaterRemoved) {
				if i+1 == len(sequence)-1 {
					return false
				}
				greaterRemoved = remove(sequence, i+1)
				if !isArrayIncreasing(greaterRemoved) {
					return false
				}
			}

		}
	}
	return true
}

func TestAlmostIncreasingSequence(t *testing.T) {
	var cases = []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 3, 2, 1},
			expected: false,
		},
		{
			input:    []int{1, 3, 2},
			expected: true,
		},
		{
			input:    []int{1, 2, 1, 2},
			expected: false,
		},
		{
			input:    []int{1, 1, 2, 3, 4, 4},
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 4, 3, 6},
			expected: true,
		},
		{
			input:    []int{3, 5, 67, 98, 3},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := almostIncreasingSequence(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
