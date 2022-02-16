package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	head.Next = sortList(head.Next)

	pos := head
	for pos.Next != nil && pos.Val > pos.Next.Val {
		// swap value
		pos.Val, pos.Next.Val = pos.Next.Val, pos.Val
		pos = pos.Next
	}
	return head
}
