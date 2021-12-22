package leetcode

import "github.com/bbs731/leetcode/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = structures.TreeNode

func mirror(root1, root2 *TreeNode) bool {
	if root1 == nil {
		return root2 == nil
	}
	if root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return mirror(root1.Left, root2.Right) && mirror(root1.Right, root2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirror(root.Left, root.Right)
}
