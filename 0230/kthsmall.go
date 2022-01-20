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
/*
Tree travseral 的 loop 的版本
 */
func kthSmallest(root *TreeNode, k int) int {

	visited := make(map[*TreeNode]bool)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	ans := make([]int, 0)

	for len(stack) != 0 {
		top := stack[len(stack)-1]
		if top.Left != nil && visited[top.Left] == false {
			stack = append(stack, top.Left)
		} else {
			// pop the top
			visited[top] = true
			stack = stack[:len(stack)-1]
			ans = append(ans, top.Val)
			if len(ans) == k{
				return ans[k-1]
			}
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
		}
	}
	return ans[k-1]
}
