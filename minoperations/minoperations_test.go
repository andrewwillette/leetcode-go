package minoperations

import (
	"fmt"
	"testing"
)

// # 1769. Minimum Number of Operations to Move All Balls to Each Box
// https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/
func minOperations(boxes string) []int {
	toReturn := []int{}
	for i := 0; i < len(boxes); i++ {
		operations := 0
		for j := 0; j < len(boxes); j++ {
			if boxes[j] == '1' {
				operations += abs(j - i)
			}
		}
		toReturn = append(toReturn, operations)
	}
	return toReturn
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func TestIntBehavior(t *testing.T) {
	toTest := "100"
	for i := 0; i < len(toTest); i++ {
		result := toTest[i] - '0'
		resultInt := int(toTest[i] - '0')
		fmt.Println(result)
		fmt.Printf("%T\n", result)
		fmt.Printf("%T\n", resultInt)
	}
}
