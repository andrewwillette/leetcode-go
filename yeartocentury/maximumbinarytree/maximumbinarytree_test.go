package maxiumumbinarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	max, left, right := getMaxLeftAndRight(nums)
	root := &TreeNode{Val: max}
	if len(left) > 0 {
		root.Left = constructMaximumBinaryTree(left)
	}
	if len(right) > 0 {
		root.Right = constructMaximumBinaryTree(right)
	}
	return root
}

func getMaxLeftAndRight(nums []int) (max int, left []int, right []int) {
	max = 0
	maxIndex := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= max {
			max = nums[i]
			maxIndex = i
		}
	}
	return max, nums[:maxIndex], nums[maxIndex+1:]
}

func TestConstructMaximumBinaryTree(t *testing.T) {
	var cases = []struct {
		input    []int
		expected *TreeNode
	}{
		{
			input: []int{3, 2, 1, 6, 0, 5},
			expected: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
				Right: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 0,
					},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := constructMaximumBinaryTree(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}

func TestGetMaxLeftAndRight(t *testing.T) {
	arr := []int{1, 2, 8, 4, 5}
	_, left, right := getMaxLeftAndRight(arr)
	fmt.Printf("left %v\n", left)
	fmt.Printf("right %v\n", right)
}
