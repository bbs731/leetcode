package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	/* 因为是树， 所以无环， 可以不使用 visited[] array 或者 hashtable */

	stack := []*TreeNode{root}
	ans := make([]int, 0)

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		ans = append(ans, top.Val)
		stack = stack[:len(stack)-1]
		/*
			先放 Right , 再放 Left 才能保证，是 pre-order
		 */
		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}
	return ans
}
