package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
面试的时候，你会死在这道题上的！

我见过最难的题！

 */

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	minD := math.MaxInt64
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}

	return minD + 1
}
