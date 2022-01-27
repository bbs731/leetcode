package leetcode

/*
	因为是树， 所以明显，不需要是用 visited hashtable

但是， 还是需要，技巧！

 */

func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	ans := []int{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		root = root.Right
	}
	return ans
}
