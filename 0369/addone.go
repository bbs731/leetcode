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

func plusOneWithCarry(head *ListNode) int {
	if head == nil {
		return 1
	}
	carry := plusOneWithCarry(head.Next)
	if head.Val+carry == 10 {
		head.Val = 0
		return 1
	}
	head.Val = head.Val + carry
	return 0
}

func plusOne(head *ListNode) *ListNode {

	carry := plusOneWithCarry(head.Next)
	if head.Val+carry == 10 {
		head.Val = 0
		// create new Node
		newHead := &ListNode{
			1,
			head,
		}
		return newHead
	}
	head.Val = head.Val + carry
	return head
}
