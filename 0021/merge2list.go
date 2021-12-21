package leetcode

import "github.com/bbs731/leetcode/structures"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	m := &ListNode{0, nil}
	head := m

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			m.Next = list1
			list1 = list1.Next
		} else {
			m.Next = list2
			list2 = list2.Next
		}
		m = m.Next
	}
	// 第一次提交， 这里的条件判断错误，少了一种情况， list1， list2 都可能先变成nil
	if list1 != nil {
		m.Next = list1
	} else {
		m.Next = list2
	}
	return head.Next
}

// 递归的写法
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}
