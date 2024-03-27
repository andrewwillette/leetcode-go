package taskscheduler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func leastInterval(tasks []byte, n int) int {
	valuesFreq := make(map[string]int)
	for _, v := range tasks {
		valuesFreq[string(v)]++
	}
	maxFreq, sameMaxCount := getMaxFreqMaxCount(valuesFreq)
	res := (maxFreq-1)*(n+1) + sameMaxCount
	if res < len(tasks) { // if res is less than tasks, return tasks
		return len(tasks)
	} else {
		return res
	}
}

func getMaxFreqMaxCount(valuesFreq map[string]int) (int, int) {
	maxFreq := 0
	sameMaxCount := 0
	for _, v := range valuesFreq {
		if v > maxFreq {
			maxFreq = v
			sameMaxCount = 1
		} else if v == maxFreq { // if v == max, increase maxCount
			sameMaxCount++
		}
	}
	return maxFreq, sameMaxCount
}

func TestLeastInterval(t *testing.T) {
	var cases = []struct {
		tasks    []byte
		n        int
		expected int
	}{
		{
			tasks:    []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:        2,
			expected: 8,
		},
		{
			tasks:    []byte{'A', 'B', 'C', 'A'},
			n:        2,
			expected: 4,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := leastInterval(c.tasks, c.n)
			require.Equal(t, c.expected, result)
		})
	}
}
