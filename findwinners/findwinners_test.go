package findwinners

import "sort"

// https://leetcode.com/problems/find-players-with-zero-or-one-losses
func findWinners(matches [][]int) [][]int {
	lossFreq := make(map[int]int)
	for _, match := range matches {
		lossFreq[match[0]] += 0
		lossFreq[match[1]] += 1
	}
	noLoss, oneLoss := []int{}, []int{}
	for teamName, lossCount := range lossFreq {
		if lossCount == 0 {
			noLoss = append(noLoss, teamName)
		} else if lossCount == 1 {
			oneLoss = append(oneLoss, teamName)
		}
	}
	sort.Ints(noLoss)
	sort.Ints(oneLoss)
	return [][]int{noLoss, oneLoss}
}
