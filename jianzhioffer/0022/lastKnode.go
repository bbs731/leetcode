package offer

import "github.com/bbs731/leetcode/structures"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	pos := head
	for i := 1; i != k && pos != nil; i++ {
		pos = pos.Next
	}

	if pos == nil {
		return nil
	}

	for pos.Next != nil {
		head = head.Next
		pos = pos.Next
	}
	return head
}
