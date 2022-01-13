package leetcode

import "github.com/bbs731/leetcode/structures"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode structures.ListNode

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	newHead := reverseList(head.Next)

	next.Next = head
	head.Next = nil // 没有这句的话， out of memory

	return newHead

}
