package freqsort

import "sort"

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
