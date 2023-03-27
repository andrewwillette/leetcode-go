package isbinarycomplete

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getMaxDepth(root.Left)
	rightDepth := getMaxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func isCompleteTree(root *TreeNode) bool {
	// Initialization
	traversalQ := list.New()
	traversalQ.PushBack(root)

	prevNode := root

	// Launch level-order traversal
	for traversalQ.Len() != 0 {

		queueHead := traversalQ.Front()
		traversalQ.Remove(queueHead)

		curNode := queueHead.Value.(*TreeNode)

		if curNode != nil {

			if prevNode == nil {
				// Empty in the middle means this is not a complete binary tree (not left-compact)
				return false
			}

			traversalQ.PushBack(curNode.Left)  // if there is no left, pushes nil
			traversalQ.PushBack(curNode.Right) // if there is no right, pushes nil
		}

		// update previous node
		prevNode = curNode
	}

	return true
}

func TestIsBinaryComplete(t *testing.T) {
	var cases = []struct {
		input    *TreeNode
		expected bool
	}{
		{
			input: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3,
					Left: &TreeNode{Val: 6},
				},
			},
			expected: true,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := isCompleteTree(c.input)
			require.Equal(t, c.expected, result)
		})
	}
}

func TestGetMaxDepth(t *testing.T) {
	tree := &TreeNode{Val: 5,
		Left:  &TreeNode{},
		Right: &TreeNode{},
	}
	result := getMaxDepth(tree)
	fmt.Printf("result: %v\n", result)
}

func TestList(t *testing.T) {
	list := list.New()
	list.PushBack(nil)
	list.PushBack(nil)
	list.PushBack(nil)
	fmt.Printf("%v\n", list.Len())
}
