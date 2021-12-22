package leetcode

/*
你可以运用递归和迭代两种方法解决这个问题吗？
 */
type pair struct {
	r1 *TreeNode
	r2 *TreeNode
}

func isSymmetricLoop(root *TreeNode) bool {
	if root == nil {
		return true
	}
	stack := make([]pair, 0)
	stack = append(stack, pair{root.Left, root.Right})

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if p.r1 == nil && p.r2 == nil {
			continue
		}

		if p.r1 == nil && p.r2 != nil {
			return false
		}
		if p.r2 == nil && p.r1 != nil {
			return false
		}

		if p.r1.Val != p.r2.Val {
			return false
		}
		stack = append(stack, pair{p.r1.Left, p.r2.Right})
		stack = append(stack, pair{p.r1.Right, p.r2.Left})
	}

	return true
}
