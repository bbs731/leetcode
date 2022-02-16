package leetcode

/*
这个复杂度是多少？ 
 */
func sortList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	// lets do the quick sort
	pivot := head
	for pos := head.Next; pos != nil; pos = pos.Next {
		if pos.Val >= head.Val {
			continue
		}
		//swap
		pivot.Next.Val, pos.Val = pos.Val, pivot.Next.Val
		pivot = pivot.Next
	}
	//swap pivot and heat
	pivot.Val, head.Val = head.Val, pivot.Val

	next := pivot.Next
	pivot.Next = nil

	// sort and join
	head = sortList(head)
	pivot.Next = sortList(next)

	return head
}
