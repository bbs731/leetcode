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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0

	head := &ListNode{}
	tail := head

	for l1 != nil || l2 != nil {

		val1, val2 := 0, 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		val := carry + val1 + val2
		if val >= 10 {
			val -= 10
			carry = 1
		} else {
			carry = 0
		}

		tail.Next = &ListNode{val, nil}
		tail = tail.Next

	}

	// special case with carry = 1
	if carry == 1 {
		tail.Next = &ListNode{1, nil}
	}
	return head.Next
}
