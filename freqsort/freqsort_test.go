package freqsort

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/sort-characters-by-frequency/
func frequencySort(s string) string {
	memo := make(map[string]int)
	for _, v := range s {
		memo[string(v)] = memo[string(v)] + 1
	}
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range memo {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	result := ""
	for _, v := range ss {
		for i := 0; i < v.Value; i++ {
			result = result + v.Key
		}
	}
	return result
}

func frequencySortOptimized(s string) string {
	if len(s) < 2 {
		return s
	}
	dict := [256]int{}
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}

	b := []byte(s) //type conversion

	sort.Slice(b, func(i, j int) bool {
		if dict[b[i]] == dict[b[j]] {
			return b[i] > b[j]
		}
		return dict[b[i]] > dict[b[j]]
	})

	return string(b)
}

func TestFrequencySort(t *testing.T) {
	// cases commented out because the result is not deterministic
	var cases = []struct {
		s        string
		expected string
	}{
		// {
		// 	s:        "tree",
		// 	expected: "eert",
		// },
		// {
		// 	s:        "cccaaa",
		// 	expected: "aaaccc",
		// },
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			result := frequencySort(c.s)
			require.Equal(t, c.expected, result)
		})
	}
}
