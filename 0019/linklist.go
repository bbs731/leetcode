package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
you are beautiful
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	//step forward n-1 steps
	for n != 1 {
		fast = fast.Next
		n--
	}

	slow := head
	var prev *ListNode

	for fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}

	if prev == nil {
		// we need to remove the head
		return head.Next
	}
	// otherwise, remove slow
	prev.Next = slow.Next
	return head
}
