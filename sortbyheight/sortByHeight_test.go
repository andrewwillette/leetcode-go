package sortbyheight

import (
	"fmt"
	"testing"

	"github.com/andrewwillette/leetcode-go/sorting"
	"github.com/stretchr/testify/require"
)

// Some people are standing in a row in a park. There are trees between them which cannot be moved.
// Your task is to rearrange the people by their heights in a non-descending order without moving the trees.
// -1 is a tree.
func sortByHeight(a []int) []int {
	// iterate through array and collect non -1 and their position
	nonzeroIndexes := make([]int, 0)
	nonzeroValues := make([]int, 0)
	for i, v := range a {
		if v != -1 {
			nonzeroIndexes = append(nonzeroIndexes, i)
			nonzeroValues = append(nonzeroValues, v)
		}
	}
	nonzeroValues = sorting.Bubblesort(nonzeroValues)
	for i, v := range nonzeroIndexes {
		a[v] = nonzeroValues[i]
	}

	return a
}

func TestSortByHeight(t *testing.T) {
	var cases = []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{-1, 150, 190, 170, -1, -1, 160, 180},
			expected: []int{-1, 150, 160, 170, -1, -1, 180, 190},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := sortByHeight(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
