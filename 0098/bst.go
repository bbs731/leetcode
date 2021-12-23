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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBST(root.Left) || !isValidBST(root.Right) {
		return false
	}

	l := root.Left
	if l != nil {
		for ; l.Right != nil; l = l.Right {
		}
		if root.Val <= l.Val {
			return false
		}
	}

	r := root.Right
	if r != nil {
		for ; r.Left != nil; r = r.Left {
		}
		if root.Val >= r.Val {
			return false
		}
	}
	return true
}
