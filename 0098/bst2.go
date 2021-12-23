package leetcode

import "math"

func helper(root *TreeNode, min, max int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

func isValidBST2(root *TreeNode) bool {

	return helper(root, math.MinInt64, math.MaxInt64)

}
