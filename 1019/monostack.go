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

func nextLargerNodes(head *ListNode) []int {

	t := head
	n := 0

	for t != nil {
		n++
		t = t.Next
	}
	ans := make([]int, n)

	type record struct {
		index int
		val   int
	}

	stack := make([]record, 1)
	stack[0] = record{
		0,
		head.Val,
	}
	top := 0
	index := 1

	head = head.Next
	for head != nil {
		if head.Val <= stack[top].val {
			stack = append(stack, record{index, head.Val})
		} else {
			for top >= 0 && head.Val > stack[top].val {
				ans[stack[top].index] = head.Val
				top--
			}
			stack = stack[:top+1]
			stack = append(stack, record{index, head.Val})
		}
		top++
		head = head.Next
		index++
	}

	return ans
}
