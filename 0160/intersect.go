package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
困扰了，你多年的题， 甚至让你到了怀疑人生的地步， 结果，它就是一个简单的题！
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}

	m, n := 1, 1

	posA := headA
	for posA.Next != nil {
		m++
		posA = posA.Next
	}
	posB := headB
	for posB.Next != nil {
		n++
		posB = posB.Next
	}

	// 2 list not intersect
	if posA != posB {
		return nil
	}

	if m < n {
		steps := n - m
		for steps > 0 {
			headB = headB.Next
			steps--
		}
	} else {
		steps := m - n
		for steps > 0 {
			headA = headA.Next
			steps--
		}
	}

	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}
