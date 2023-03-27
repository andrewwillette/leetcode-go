package islucky

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// Ticket numbers usually consist of an even number of digits. A ticket number is considered lucky if the
// sum of the first half of the digits is equal to the sum of the second half.
// Given a ticket number n, determine if it's lucky or not.
func isLucky(n int) bool {
	nLetters := strconv.Itoa(n)
	nLettersArray := []rune(nLetters)
	nNumbersArray := make([]int, len(nLettersArray))
	for i := 0; i < len(nLettersArray); i++ {
		nNumbersArray[i], _ = strconv.Atoi(string(nLettersArray[i]))
	}
	leftTotal := 0
	for i := 0; i < len(nNumbersArray)/2; i++ {
		leftTotal += nNumbersArray[i]
	}
	rightTotal := 0
	for i := len(nNumbersArray) / 2; i < len(nNumbersArray); i++ {
		rightTotal += nNumbersArray[i]
	}
	return leftTotal == rightTotal
}

func TestIsLucky(t *testing.T) {
	var cases = []struct {
		input    int
		expected bool
	}{
		{
			input:    1230,
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := isLucky(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
