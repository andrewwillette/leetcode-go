package commoncharactercount

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func stringAsCountedMap(s string) map[rune]int {
	var countMap = make(map[rune]int)
	for _, c := range s {
		countMap[c]++
	}
	return countMap
}

func commonCharacterCount(s1, s2 string) int {
	s1CountMap := stringAsCountedMap(s1)
	s2CountMap := stringAsCountedMap(s2)
	commonCharacters := 0
	for k := range s1CountMap {
		if s2CountMap[k] > 0 {
			for {
				commonCharacters += 1
				s1CountMap[k] -= 1
				s2CountMap[k] -= 1
				if s1CountMap[k] == 0 || s2CountMap[k] == 0 {
					break
				}
			}
		}
	}
	return commonCharacters
}

func TestCommonCharacterCount(t *testing.T) {
	var cases = []struct {
		s1       string
		s2       string
		expected int
	}{
		{
			s1:       "aabcc",
			s2:       "adcaa",
			expected: 3,
		},
		{
			s1:       "zzzz",
			s2:       "zzzzzzz",
			expected: 4,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v %v", c.s1, c.s2), func(t *testing.T) {
			result := commonCharacterCount(c.s1, c.s2)
			require.Equal(t, c.expected, result)
		})
	}
}
