package addtwonumbers

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) insert(next int) {
	node := n
	for {
		if node.Next != nil {
			node = node.Next
		} else {
			node.Next = &ListNode{Val: next}
			break
		}
	}
	return
}

func TestListNodeInsert(t *testing.T) {
	node := ListNode{Val: 5, Next: nil}
	node.insert(5)
	fmt.Printf("%v\n", node)
}

func getListNodeTotalReverseVal(l *ListNode) int {
	multiplier := 1
	total := l.Val
	l = l.Next
	for {
		if l == nil {
			break
		}
		toAdd := int(math.Pow10(multiplier))
		// fmt.Printf("toAdd: %v\n", toAdd)
		total += (l.Val * toAdd)
		// fmt.Printf("total: %v\n", total)
		l = l.Next
		multiplier += 1
	}
	return total
}

func intToReverseList(v int) *ListNode {
	first := v % 10
	v = v - first
	v = v / 10
	ln := ListNode{Val: first}
	fmt.Printf("firstVal: %v\n", first)
	for {
		nextVal := v % 10
		fmt.Printf("nextVal: %v\n", nextVal)
		ln.insert(nextVal)
		v = v - nextVal
		v = v / 10
		if v == 0 {
			break
		}
	}
	return &ln
}

// addTwoNumbers
// `https://leetcode.com/problems/add-two-numbers/`
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	l1Total := getListNodeTotalReverseVal(l1)
// 	l2Total := getListNodeTotalReverseVal(l2)
// 	result := l1Total + l2Total
// 	return intToReverseList(result)
// }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, dummy := 0, new(ListNode)
	for node := dummy; l1 != nil || l2 != nil || carry > 0; node = node.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		node.Next = &ListNode{carry % 10, nil}
		carry /= 10
	}
	return dummy.Next
}

func TestAddTwoNumbers(t *testing.T) {
	var cases = []struct {
		listOne  ListNode
		listTwo  ListNode
		expected ListNode
	}{
		{
			listOne:  ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			listTwo:  ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			expected: ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
		},
	}
	for _, c := range cases {
		result := addTwoNumbers(&c.listOne, &c.listTwo)
		require.Equal(t, c.expected, result)
	}
}
