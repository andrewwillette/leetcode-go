package makearrayconsecutive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func bubblesort(input []int) []int {
	for i := len(input); i > 0; i-- {
		for j := 1; j < i; j++ {
			if input[j-1] > input[j] {
				intermediate := input[j]
				input[j] = input[j-1]
				input[j-1] = intermediate
			}
		}
	}
	return input
}

// Ratiorg got statues of different sizes as a present from CodeMaster for his birthday, each statue
// having an non-negative integer size. Since he likes to make things perfect, he wants to arrange
// them from smallest to largest so that each statue will be bigger than the previous one exactly by 1.
// He may need some additional statues to be able to accomplish that. Help him figure out the minimum
// number of additional statues needed.
func makeArrayConsecutive(statues []int) int {
	sorted := bubblesort(statues)
	total := 0
	for i := 0; i < len(sorted)-1; i++ {
		difference := sorted[i+1] - sorted[i]
		if difference > 1 {
			total += difference - 1
		}
	}
	return total
}

func TestMakeArrayConsecutive(t *testing.T) {
	var cases = []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{6, 2, 3, 8},
			expected: 3,
		},
	}
	for _, c := range cases {
		result := makeArrayConsecutive(c.input)
		require.Equal(t, c.expected, result)
	}
}
