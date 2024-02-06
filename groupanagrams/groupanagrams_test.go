package groupanagrams

import (
	"sort"
	"strings"
)

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	memo := make(map[string][]string)

	for _, v := range strs {
		memoKey := []string{}
		for _, r := range v {
			memoKey = append(memoKey, string(r))
		}
		sort.Strings(memoKey)
		mks := strings.Join(memoKey[:], "")
		memo[mks] = append(memo[mks], v)
	}
	result := [][]string{}
	for _, anagramGroup := range memo {
		result = append(result, anagramGroup)
	}
	return result
}
