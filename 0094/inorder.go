package leetcode

import "github.com/bbs731/leetcode/structures"

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
	res := inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}
