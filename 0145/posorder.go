package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
还可以不用 visited[] hashtable 吗？

可以， 要了命了， 做不出来！

 */
func postorderTraversal(root *TreeNode) []int {
	// 先左， 再右， 最后 node
	stack := make([]*TreeNode, 0)
	var prev *TreeNode
	ans := make([]int, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			// 重新入栈
			stack = append(stack, root)
			root = root.Right
		}
	}
	return ans
}
