package leftrightsumdif

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestSliceSplicing(t *testing.T) {
// 	arr := []int{1, 2, 3, 4, 5}
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Printf("leftSide : %v\n", arr[:i])
// 		fmt.Printf("rightSide : %v\n", arr[i+1:])
// 	}
// }

func leftRigthDifference(nums []int) []int {
	var toReturn = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		toReturn[i] = int(math.Abs(float64(getLeftSum(nums, i) - getRightSum(nums, i))))
	}
	return toReturn
}

func getLeftSum(nums []int, i int) int {
	leftSum := 0
	toSum := nums[:i]
	fmt.Printf("leftSide : %v\n", toSum)
	fmt.Printf("%v\n", nums)
	for j := 0; j < len(toSum); j++ {
		leftSum += toSum[j]
	}
	fmt.Printf("leftSum: %v\n", leftSum)
	return leftSum
}

func getRightSum(nums []int, i int) int {
	rightSum := 0
	toSum := nums[i+1:]
	fmt.Printf("rightSide : %v\n", toSum)
	fmt.Printf("%v\n", nums)
	for j := 0; j < len(toSum); j++ {
		rightSum += toSum[j]
	}
	fmt.Printf("rightSum: %v\n", rightSum)
	return rightSum
}

func TestLeftRightDiff(t *testing.T) {
	var cases = []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{10, 4, 8, 3},
			expected: []int{15, 1, 11, 22},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := leftRigthDifference(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}
