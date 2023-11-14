package insertgreatestcommondivisor

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	result := head
	for head.Next != nil {
		toInsert := ListNode{Val: GCD(head.Val, head.Next.Val)}
		next := head.Next
		head.Next = &toInsert
		toInsert.Next = next
		head = next
	}
	return result
}

func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	return GCD(y, x%y)
}
