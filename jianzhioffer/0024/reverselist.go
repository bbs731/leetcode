package offer

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
/*
	果然，你写第一遍的时候，还是错的! 链表啊，你永远的痛！
 */
func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	newHead := reverseList(head.Next)
	next.Next = head
	head.Next = nil
	return newHead

}
