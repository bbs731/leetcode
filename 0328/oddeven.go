package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	pos := head
	even := &ListNode{}
	etail := even
	for pos.Next != nil {
		// pos.Next is an even node, remove it from list
		n := pos.Next
		// add it to even list
		etail.Next = n
		etail = n

		pos.Next = n.Next
		if n.Next != nil {
			pos = pos.Next
		}
	}

	etail.Next = nil
	pos.Next = even.Next

	return head
}
