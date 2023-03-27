package aresimilar

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// searchForValue returns the matching indexes in the slice
func searchForValue(a []int, v int) []int {
	existingIndexes := make([]int, 0)
	for i := 0; i < len(a); i++ {
		if a[i] == v {
			existingIndexes = append(existingIndexes, i)
		}
	}
	return existingIndexes
}

func swapValues(a []int, i, j int) []int {
	// fmt.Printf("swapValues: %v, %v, %v\n", a, i, j)
	result := make([]int, len(a))
	copy(result, a)
	tmpI := result[i]
	tmpJ := result[j]
	result[i], result[j] = tmpJ, tmpI
	return result
}

func slicesEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Two arrays are called similar if one can be obtained from another by swapping at most one pair of elements in one of the arrays.
// Given two arrays a and b, check whether they are similar.
func areSimilar(a []int, b []int) bool {
Loop:
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			// fmt.Printf("a[i] != b[i]: %v != %v\n", a[i], b[i])
			existingIndexes := searchForValue(b, a[i])
			// fmt.Printf("existingIndexes: %v\n", existingIndexes)
			for _, index := range existingIndexes {
				bSwapped := swapValues(b, i, index)
				if slicesEqual(a, bSwapped) {
					break Loop
				}
			}
			return false
		}
	}
	return true
}

func TestAreSimilar(t *testing.T) {
	var cases = []struct {
		name     string
		inputa   []int
		inputb   []int
		expected bool
	}{
		{
			name:     "case1",
			inputa:   []int{1, 2, 3},
			inputb:   []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "case2",
			inputa:   []int{1, 2, 3},
			inputb:   []int{2, 1, 3},
			expected: true,
		},
		{
			name:     "case3",
			inputa:   []int{1, 2, 2},
			inputb:   []int{2, 1, 1},
			expected: false,
		},
		{
			name:     "case3",
			inputa:   []int{1, 2, 1, 2},
			inputb:   []int{2, 2, 1, 1},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := areSimilar(c.inputa, c.inputb)
			require.Equal(t, c.expected, result)
		})
	}
}
