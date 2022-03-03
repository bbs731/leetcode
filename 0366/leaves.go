package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) (bool, []int) {
	if root == nil {
		return false, []int{}
	}

	if root.Left == nil && root.Right == nil {
		return true, []int{root.Val}
	}

	lr, ls := dfs(root.Left)
	rr, rs := dfs(root.Right)

	leaves := append([]int{}, ls...)
	leaves = append(leaves, rs...)
	if lr {
		root.Left = nil
	}
	if rr {
		root.Right = nil
	}
	return false, leaves

}
func findLeaves(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	/* 这样看起来是不是很酷！ */
	for r, l := dfs(root); r == false; r, l = dfs(root) {
		ans = append(ans, l)
	}
	ans = append(ans, []int{root.Val})
	return ans
}
